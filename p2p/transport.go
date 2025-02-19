package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}

// Transport handles the communication between the nodes in the network
// This can be of the form (TCP, UDP, websockets, etc....)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
