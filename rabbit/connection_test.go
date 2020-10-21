package rabbit

import (
    "library/network/transport"
    "testing"
)

var rabbitAddrs = []string{"10.31.100.123:5672"}
func TestNewTransport(t *testing.T) {
   NewTransport(func(op *transport.Options) {
        op.Addrs = rabbitAddrs
    })


}