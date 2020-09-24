package discover

import (
	"crypto/ecdsa"
	"errors"
	"net"
	"time"

	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/p2p/enode"
	"gitlab.com/q-dev/q-client/p2p/enr"
	"gitlab.com/q-dev/q-client/p2p/netutil"
	"gitlab.com/q-dev/q-client/rlp"
)

// RPC packet types
const (
	p_pingV4 = iota + 1 // zero is 'reserved'
	p_pongV4
	p_findnodeV4
	p_neighborsV4
	p_enrRequestV4
	p_enrResponseV4
	p_requestAllV4
	p_respondWithAllV4
)

// RPC request structures
type (
	pingV4 struct {
		senderKey *ecdsa.PublicKey // filled in by preverify

		Version    uint
		From, To   rpcEndpoint
		Expiration uint64
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// pongV4 is the reply to pingV4.
	pongV4 struct {
		// This field should mirror the UDP envelope address
		// of the ping packet, which provides a way to discover the
		// the external address (after NAT).
		To rpcEndpoint

		ReplyTok   []byte // This contains the hash of the ping packet.
		Expiration uint64 // Absolute timestamp at which the packet becomes invalid.
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// findnodeV4 is a query for nodes close to the given target.
	findnodeV4 struct {
		Target     encPubkey
		Expiration uint64
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// neighborsV4 is the reply to findnodeV4.
	neighborsV4 struct {
		Nodes      []rpcNode
		Expiration uint64

		// TotalCount of nodes that a requesting peer should receive.
		// Note: without this number, the requesting peer was waiting
		// for more than bucketSize nodes even if responding peer
		// sent everything it had and then dropped responder as timed out.
		// Because of this, peers with small number of items in table
		// didn't participate in discovery. ¯\_(ツ)_/¯
		TotalCount uint

		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// enrRequestV4 queries for the remote node's record.
	enrRequestV4 struct {
		Expiration uint64
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// enrResponseV4 is the reply to enrRequestV4.
	enrResponseV4 struct {
		ReplyTok []byte // Hash of the enrRequest packet.
		Record   enr.Record
		// Ignore additional fields (for forward compatibility).
		Rest []rlp.RawValue `rlp:"tail"`
	}

	// requestAll asks node for its known peers.
	requestAll struct {
		Expiration uint64
	}

	// respondWithAll is the reply to requestAll.
	respondWithAll struct {
		Nodes      []rpcNode
		TotalCount uint
		Expiration uint64
	}

	rpcNode struct {
		IP  net.IP // len 4 for IPv4 or 16 for IPv6
		UDP uint16 // for discovery protocol
		TCP uint16 // for RLPx protocol
		ID  encPubkey
	}

	rpcEndpoint struct {
		IP  net.IP // len 4 for IPv4 or 16 for IPv6
		UDP uint16 // for discovery protocol
		TCP uint16 // for RLPx protocol
	}
)

// packetV4 is implemented by all v4 protocol messages.
type packetV4 interface {
	// preverify checks whether the packet is valid and should be handled at all.
	preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error
	// handle handles the packet.
	handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte)
	// packet name and type for logging purposes.
	name() string
	kind() byte
}

// PING/v4

func (req *pingV4) name() string { return "PING/v4" }
func (req *pingV4) kind() byte   { return p_pingV4 }

func (req *pingV4) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	if expired(req.Expiration) {
		return errExpired
	}
	key, err := decodePubkey(crypto.S256(), fromKey)
	if err != nil {
		return errors.New("invalid public key")
	}
	req.senderKey = key
	return nil
}

func (req *pingV4) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {
	// Reply.
	seq, _ := rlp.EncodeToBytes(t.localNode.Node().Seq())
	t.send(from, fromID, &pongV4{
		To:         makeEndpoint(from, req.From.TCP),
		ReplyTok:   mac,
		Expiration: uint64(time.Now().Add(expiration).Unix()),
		Rest:       []rlp.RawValue{seq},
	})

	// Ping back if our last pong on file is too far in the past.
	n := wrapNode(enode.NewV4(req.senderKey, from.IP, int(req.From.TCP), from.Port))
	if time.Since(t.db.LastPongReceived(n.ID(), from.IP)) > bondExpiration {
		t.sendPing(fromID, from, func() {
			t.tab.addVerifiedNode(n)
		})
	} else {
		t.tab.addVerifiedNode(n)
	}

	// Update node database and endpoint predictor.
	t.db.UpdateLastPingReceived(n.ID(), from.IP, time.Now())
	t.localNode.UDPEndpointStatement(from, &net.UDPAddr{IP: req.To.IP, Port: int(req.To.UDP)})
}

// PONG/v4

func (req *pongV4) name() string { return "PONG/v4" }
func (req *pongV4) kind() byte   { return p_pongV4 }

func (req *pongV4) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	if expired(req.Expiration) {
		return errExpired
	}
	if !t.handleReply(fromID, from.IP, req) {
		return errUnsolicitedReply
	}
	return nil
}

func (req *pongV4) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {
	t.localNode.UDPEndpointStatement(from, &net.UDPAddr{IP: req.To.IP, Port: int(req.To.UDP)})
	t.db.UpdateLastPongReceived(fromID, from.IP, time.Now())
}

// FINDNODE/v4

func (req *findnodeV4) name() string { return "FINDNODE/v4" }
func (req *findnodeV4) kind() byte   { return p_findnodeV4 }

func (req *findnodeV4) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	return preverifyNodesQuery(req.Expiration, t, from, fromID)
}

func preverifyNodesQuery(expiration uint64, t *UDPv4, from *net.UDPAddr, fromID enode.ID) error {
	if expired(expiration) {
		return errExpired
	}

	if !t.checkBond(fromID, from.IP) {
		// No endpoint proof pong exists, we don't process the packet. This prevents an
		// attack vector where the discovery protocol could be used to amplify traffic in a
		// DDOS attack. A malicious actor would send a findnode request with the IP address
		// and UDP port of the target as the source address. The recipient of the findnode
		// packet would then send a neighbors packet (which is a much bigger packet than
		// findnode) to the victim.
		return errUnknownNode
	}

	return nil
}

func (req *findnodeV4) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {
	// Determine closest nodes.
	target := enode.ID(crypto.Keccak256Hash(req.Target[:]))
	t.tab.mutex.Lock()
	closest := t.tab.closest(target, bucketSize, true).entries
	t.tab.mutex.Unlock()

	// Send neighbors in chunks with at most maxNeighbors per packet
	// to stay below the packet size limit.
	p := neighborsV4{
		Expiration: uint64(time.Now().Add(expiration).Unix()),
		TotalCount: uint(len(closest)),
	}

	var sent bool
	for _, n := range closest {
		if netutil.CheckRelayIP(from.IP, n.IP()) == nil {
			p.Nodes = append(p.Nodes, nodeToRPC(n))
		}
		if len(p.Nodes) == maxNeighbors {
			t.send(from, fromID, &p)
			p.Nodes = p.Nodes[:0]
			sent = true
		}
	}
	if len(p.Nodes) > 0 || !sent {
		t.send(from, fromID, &p)
	}
}

// NEIGHBORS/v4

func (req *neighborsV4) name() string { return "NEIGHBORS/v4" }
func (req *neighborsV4) kind() byte   { return p_neighborsV4 }

func (req *neighborsV4) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	return preverifyNodesResponse(req.Expiration, req, t, from, fromID)
}

func preverifyNodesResponse(expiration uint64, packet packetV4, t *UDPv4, from *net.UDPAddr, fromID enode.ID) error {
	if expired(expiration) {
		return errExpired
	}

	// TODO: it's actual 'handle' in verify, not cool
	if !t.handleReply(fromID, from.IP, packet) {
		return errUnsolicitedReply
	}

	return nil
}

func (req *neighborsV4) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {}

// REQUESTALL

func (req *requestAll) name() string { return "REQUESTALL/v4" }

func (req *requestAll) kind() byte { return p_requestAllV4 }

func (req *requestAll) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	return preverifyNodesQuery(req.Expiration, t, from, fromID)
}

func (req *requestAll) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {
	// same as with neighbours - send in chunks
	allNodes := t.tab.getAllNodes()
	p := &respondWithAll{
		Expiration: uint64(time.Now().Add(expiration).Unix()),
		TotalCount: uint(len(allNodes)),
	}

	var sent bool
	for _, node := range allNodes {
		if netutil.CheckRelayIP(from.IP, node.IP()) == nil {
			p.Nodes = append(p.Nodes, nodeToRPC(node))
		}

		if len(p.Nodes) == maxNeighbors {
			_, _ = t.send(from, fromID, p)
			p.Nodes = nil
			sent = true
		}
	}

	if len(p.Nodes) > 0 || !sent {
		_, _ = t.send(from, fromID, p)
	}
}

// RESPONDWITHALL/V4

func (req *respondWithAll) name() string { return "RESPONDWITHALL/V4" }

func (req *respondWithAll) kind() byte { return p_respondWithAllV4 }

func (req *respondWithAll) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	return preverifyNodesResponse(req.Expiration, req, t, from, fromID)
}

func (req *respondWithAll) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {}

// ENRREQUEST/v4

func (req *enrRequestV4) name() string { return "ENRREQUEST/v4" }
func (req *enrRequestV4) kind() byte   { return p_enrRequestV4 }

func (req *enrRequestV4) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	if expired(req.Expiration) {
		return errExpired
	}
	if !t.checkBond(fromID, from.IP) {
		return errUnknownNode
	}
	return nil
}

func (req *enrRequestV4) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {
	t.send(from, fromID, &enrResponseV4{
		ReplyTok: mac,
		Record:   *t.localNode.Node().Record(),
	})
}

// ENRRESPONSE/v4

func (req *enrResponseV4) name() string { return "ENRRESPONSE/v4" }
func (req *enrResponseV4) kind() byte   { return p_enrResponseV4 }

func (req *enrResponseV4) preverify(t *UDPv4, from *net.UDPAddr, fromID enode.ID, fromKey encPubkey) error {
	if !t.handleReply(fromID, from.IP, req) {
		return errUnsolicitedReply
	}
	return nil
}

func (req *enrResponseV4) handle(t *UDPv4, from *net.UDPAddr, fromID enode.ID, mac []byte) {
}
