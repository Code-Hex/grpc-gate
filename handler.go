package gate

import (
	"errors"
	"io"
	"log"
	"net"

	"github.com/Code-Hex/grpc-gate/internal/proto"
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
	log.Println("called stream")
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
	log.Println(upstreamNetwork, net.JoinHostPort(upstreamHost, upstreamPort))
	conn, err := net.Dial(upstreamNetwork, net.JoinHostPort(upstreamHost, upstreamPort))
	if err != nil {
		return err
	}
	defer conn.Close()

	// grpc -> upstream
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("grpc -> upstream ctx error", ctx.Err())
				return
			default:
			}

			chunk, err := ss.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("closed gRPC (grpc -> upstream)", err)
					return
				}
				log.Println("grpc -> upstream error", err)
				return
			}
			if _, err := conn.Write(chunk.GetData()); err != nil {
				log.Println("grpc -> upstream conn write error", err)
				return
			}
		}
	}()

	// upstream -> grpc
	go func() {
		b := make([]byte, bufferSize)
		for {
			select {
			case <-ctx.Done():
				log.Println("upstream -> grpc ctx error", ctx.Err())
				return
			default:
			}

			n, err := conn.Read(b)
			if err != nil {
				if err == io.EOF {
					log.Println("closed tcp down stream", err)
					return
				}
				log.Println("upstream -> grpc error", err)
				return
			}

			err = ss.Send(&proto.Chunk{
				Data: b[:n],
			})
			if err == io.EOF {
				log.Println("closed gRPC (upstream -> grpc)", err)
				return
			}
			if err != nil {
				log.Println("upstream -> grpc error", err)
				return
			}
		}
	}()

	return nil
}
