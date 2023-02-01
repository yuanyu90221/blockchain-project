package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.Mutex
	peers     map[NetAddr]*LocalTransport
}
