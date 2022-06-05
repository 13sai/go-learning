package rpc

import (
	"encoding/json"
	"io"
	"net"
	"sync"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (s *Server) Accept(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		go s.ServeConn(conn)
	}
}

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() {
		conn.Close()
	}()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		return
	}

	if opt.MagicNumber != MagicNumber {
		return
	}

	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		return
	}

	s.serverCodec(f(conn))
}

var invalidRequest = struct{}{}

func (s *server) serverCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for {
		req, err := s.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			s.sendResp(cc, req.h, invalidRequest, sending)
			continue
		}

		wg.Add(1)
		go s.handleRequest(cc, req, sending, wg)
	}

	wg.Wait()
	cc.Close()
}

func Accept(l net.Listener) {
	DefaultServer.Accept(l)
}
