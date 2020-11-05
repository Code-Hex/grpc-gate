package gate

import (
	"errors"
	"io"
	"net"

	"github.com/Code-Hex/grpc-gate/internal/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewHandler(gRPCOpts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(gRPCOpts...)
	proto.RegisterStreamServer(srv, &streamServer{})
	return srv
}

type streamServer struct{}

// bufferSize for coybuffer and grpc
const bufferSize = 256 * 1024

func (s *streamServer) ServerStream(ss proto.Stream_ServerStreamServer) error {
	ctx := ss.Context()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is unavailable")
	}
	upstreamNetworks := md.Get(upstreamNetworkKey)
	if len(upstreamNetworks) == 0 {
		return errors.New("metadata upstream-port is unavailable")
	}

	upstreamHosts := md.Get(upstreamHostKey)
	if len(upstreamHosts) == 0 {
		return errors.New("metadata upstream-host is unavailable")
	}

	upstreamPorts := md.Get(upstreamPortKey)
	if len(upstreamPorts) == 0 {
		return errors.New("metadata upstream-port is unavailable")
	}

	upstreamNetwork := upstreamNetworks[0]
	upstreamHost, upstreamPort := upstreamHosts[0], upstreamPorts[0]
	conn, err := net.Dial(upstreamNetwork, net.JoinHostPort(upstreamHost, upstreamPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	eg, ctx := errgroup.WithContext(ctx)

	// grpc -> upstream
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			chunk, err := ss.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			if _, err := conn.Write(chunk.GetData()); err != nil {
				return err
			}
		}
	})

	// upstream -> grpc
	eg.Go(func() error {
		b := make([]byte, bufferSize)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			n, err := conn.Read(b)
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}

			err = ss.Send(&proto.Chunk{
				Data: b[:n],
			})
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
		}
	})
	return eg.Wait()
}
