package kshaka

import (
	"sync"
)

// InmemTransport Implements the Transport interface, to allow Raft to be
// tested in-memory without going over a network.
type InmemTransport struct {
	sync.RWMutex
	nodeID      uint64
	nodeAddress string
	nodeport    string
}

// NewInmemTransport is used to initialize a new transport
func NewInmemTransport(nodeID uint64, nodeAddress string) *InmemTransport {
	trans := &InmemTransport{nodeID: nodeID, nodeAddress: nodeAddress, nodeport: "15000"}
	return trans
}

// SendRPC sends the appropriate RPC to the target node.
// SendRPC(nodeID uint64, nodeAddress string, rpcMethod string, req RPCrequest, resp *RPCresponse) error

// // SendRPC implements the Transport interface.
// func (i *InmemTransport) SendRPC(nodeID uint64, nodeAddress string, rpcMethod string, req RPCrequest, resp *RPCresponse) error{
// 	rpcResp, err := i.makeRPC(target, args, nil, i.timeout)
// 	if err != nil {
// 		return err
// 	}

// 	// Copy the result back
// 	out := rpcResp.Response.(*AppendEntriesResponse)
// 	*resp = *out
// 	return nil
// }
