package p2p

import "errors"

// ErrInvalidHandshake is returned if the handshake between
// the local and remote node could not be established
var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
