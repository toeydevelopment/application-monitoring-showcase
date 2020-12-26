package main

import (
	"context"
	"log"
	"net"
	"time"

	trace "fulltrace/protobuf"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		panic(err)
	}

	defer listen.Close()

	s := grpc.NewServer(grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery())))

	serve := new(server)

	trace.RegisterFullTraceServiceServer(s, serve)

	defer s.Stop()

	log.Println("[INFO] Server Starting on 0.0.0.0:50051")

	if err := s.Serve(listen); err != nil {
		log.Println("[ERROR] Create server error : " + err.Error())
	}

}

func (s *server) HelloWorld(ctx context.Context, req *trace.HelloWorldRequest) (*trace.HelloWorldResponse, error) {

	span, ctx := apm.StartSpan(ctx, "SELECT HELLO WORLD", "db")

	defer span.End()

	time.Sleep(time.Second)

	return &trace.HelloWorldResponse{
		Name: "OK",
	}, nil
}
