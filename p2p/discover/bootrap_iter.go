package discover

import (
	"context"

	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/p2p/enode"
)

type bootstrapIter struct {
	ctx      context.Context
	cancelFn context.CancelFunc
	log      log.Logger

	knownNodes map[enode.ID]struct{}

	queue    []*node
	buffer   []*node
	isInited bool

	queryFn queryFunc
}

func newBootstrapIter(ctx context.Context, tab *Table, qFn queryFunc) *bootstrapIter {
	allNodes := tab.getAllNodes()
	set := make(map[enode.ID]struct{})

	for _, node := range allNodes {
		set[node.ID()] = struct{}{}
	}

	ctx, cancel := context.WithCancel(ctx)
	return &bootstrapIter{
		ctx:        ctx,
		cancelFn:   cancel,
		log:        tab.log,
		knownNodes: set,
		queue:      allNodes,
		buffer:     allNodes,
		queryFn:    qFn,
	}
}

// Node returns the current node.
func (it *bootstrapIter) Node() *enode.Node {
	if len(it.buffer) == 0 {
		return nil
	}

	return unwrapNode(it.buffer[0])
}

// Next node in iter (if available).
func (it *bootstrapIter) Next() bool {
	if len(it.buffer) > 0 {
		// prevent lost of 0 element
		// on the first call of this method
		if !it.isInited {
			it.isInited = true
			return true
		}

		it.buffer = it.buffer[1:]
	}

	for len(it.buffer) == 0 {
		if !it.advance() {
			return false
		}
	}

	return true
}

// advance takes next node to query from the queue.
// returns false when there's no more peers to query.
func (it *bootstrapIter) advance() bool {
	if it.ctx.Err() != nil {
		it.log.Trace("bootstrap iter is canceled", "err", it.ctx.Err())
		return false
	}

	if len(it.queue) == 0 {
		return false
	}

	dest := it.queue[0]
	it.queue = it.queue[1:]

	// TODO: give it a few more chances
	all, err := it.queryFn(dest)
	if err != nil {
		it.log.Debug("failed to query nodes from peer", "id", dest.ID(), "err", err)
		return true
	}

	var found bool
	for _, node := range all {
		if _, ok := it.knownNodes[node.ID()]; ok {
			continue
		}

		found = true
		it.knownNodes[node.ID()] = struct{}{}
		it.buffer = append(it.buffer, node)
	}

	if !found {
		it.log.Trace("no new nodes discovered from the peer", "id", dest.ID())
	}

	return true
}

// Close the iterator.
func (it *bootstrapIter) Close() {
	it.cancelFn()
}
