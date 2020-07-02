package discover

import (
	"context"
	"net"
	"testing"

	"gitlab.com/q-dev/go-ethereum/p2p/enode"
)

func TestBootstrapIter_Node(t *testing.T) {
	transp := newPingRecorder()

	tab, db := newTestTable(transp)
	<-tab.initDone
	defer db.Close()
	defer tab.close()

	self := transp.Self()
	boot := nodeAtDistance(self.ID(), 1, net.IP{54, 14, 21, 4})
	nodeA := nodeAtDistance(boot.ID(), 2, net.IP{130, 18, 0, 1})
	nodeB := nodeAtDistance(boot.ID(), 3, net.IP{93, 115, 5, 2})

	tab.addSeenNode(boot)
	tab.addSeenNode(nodeB)

	queryFn := mockQueryFunc(t, wrapNode(self), boot, nodeA, nodeB)
	iter := newBootstrapIter(context.Background(), tab, queryFn)

	set := map[enode.ID]struct{}{}
	for iter.Next() {
		node := iter.Node()
		if node == nil {
			t.Fatal("Node can't return nil after Next returned true")
		}
		set[node.ID()] = struct{}{}
	}

	for _, expected := range []*node{wrapNode(self), boot, nodeA, nodeB} {
		if _, ok := set[expected.ID()]; !ok {
			t.Errorf("expected %s to be discovered", expected.IP().String())
		}
	}
}

// emulate the next configuration:
// self <-> {boot, nodeB}
// boot <-> {nodeA, nodeB}
func mockQueryFunc(t *testing.T, self, boot, nodeA, nodeB *node) queryFunc {
	return func(n *node) ([]*node, error) {
		switch n.ID() {
		case boot.ID():
			return []*node{self, nodeA, nodeB}, nil
		case nodeB.ID():
			return []*node{self, boot}, nil
		case nodeA.ID():
			t.Error("nodeA is returned from boot and shouldn't be called")
			return nil, nil
		default:
			panic("unknown node")
		}
	}
}
