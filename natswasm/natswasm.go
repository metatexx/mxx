package natswasm

import (
	"context"
	"net"
	"time"

	"github.com/nats-io/nats.go"
	"nhooyr.io/websocket"
)

type ConnectionWrapper struct {
	TimeOut time.Duration
	NoTLS   bool
	ws      *websocket.Conn
}

var _ nats.CustomDialer = (*ConnectionWrapper)(nil) // Verify the implementation

func (cw ConnectionWrapper) Dial(_ string, address string) (net.Conn, error) {
	var ctx context.Context
	if cw.TimeOut > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), cw.TimeOut)
		defer cancel()
	} else {
		ctx = context.Background()
	}

	if cw.NoTLS {
		address = "ws://" + address
	} else {
		address = "wss://" + address
	}
	c, _, err := websocket.Dial(ctx, address, nil)
	if err != nil {
		return nil, err
	}
	cw.ws = c
	return websocket.NetConn(context.Background(), c, websocket.MessageBinary), nil
}

func (cw ConnectionWrapper) SkipTLSHandshake() bool {
	return true
}
