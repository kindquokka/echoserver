package server

import (
	"log"
	"net"
	"io"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListenAndServe(addr string) error {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConnection(conn)
	}

	return nil
}

func (s *Server) handleConnection(client net.Conn) {
	io.Copy(client, client)
}





