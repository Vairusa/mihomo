package utils

import (
	"fmt"
	"net"
)

type WithUpstream interface {
	Upstream() any
}

type stdWithUpstreamNetConn interface {
	NetConn() net.Conn
}

func Cast[T any](obj any) (_ T, _ bool) {
	if c, ok := obj.(T); ok {
		fmt.Printf("Got 1: %T\n", obj) // TODO
		return c, true
	}
	if u, ok := obj.(WithUpstream); ok {
		fmt.Printf("Upstream 2: %T\n", obj) // TODO
		return Cast[T](u.Upstream())
	}
	if u, ok := obj.(stdWithUpstreamNetConn); ok {
		fmt.Printf("Std 3: %T\n", obj) // TODO
		return Cast[T](u.NetConn())
	}
	fmt.Printf("Failed: %T\n", obj) // TODO
	return
}
