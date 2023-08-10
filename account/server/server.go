package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	Address string
}

func (s *Server) Launch(bind func(*grpc.Server), opts ...grpc.ServerOption) {
	server := grpc.NewServer(opts...)
	bind(server)

	reflection.Register(server)

	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on: %s", s.Address)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
