package gate

import (
	"context"
	"net"
	"time"

	"github.com/Code-Hex/grpc-gate/internal/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Dial(network, address string) (net.Conn, error) {
	var d Dialer
	return d.Dial(network, address)
}

type Dialer struct {
	streamClient proto.StreamClient
}

func NewDialer(gateAddr string, opts ...grpc.DialOption) (*Dialer, error) {
	conn, err := grpc.Dial(gateAddr, opts...)
	if err != nil {
		return nil, err
	}
	return &Dialer{streamClient: proto.NewStreamClient(conn)}, nil
}

func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		upstreamNetworkKey, network,
		upstreamHostKey, host,
		upstreamPortKey, port,
	)
	c, err := d.streamClient.ServerStream(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Conn{stream: c}, nil
}

type Conn struct {
	net.Conn
	stream proto.Stream_ServerStreamClient
}

func (c *Conn) Read(b []byte) (int, error) {
	chunk, err := c.stream.Recv()
	if err != nil {
		return 0, err
	}
	return copy(b, chunk.GetData()), nil
}

func (c *Conn) Write(b []byte) (int, error) {
	err := c.stream.Send(&proto.Chunk{
		Data: b,
	})
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (c *Conn) Close() error {
	return c.stream.CloseSend()
}

func (c *Conn) SetDeadline(t time.Time) error {
	return nil
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return nil
}
