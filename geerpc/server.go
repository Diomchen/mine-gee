package geerpc

import (
	"encoding/json"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"sync"
)

const MagicNumber = 0x3b1f1a00

type Options struct {
	MagicNumber uint32
	CodecType   codec.Type
}

var DefaultOption = &Options{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (s *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc accept error:", err)
			return
		}

		go s.ServeConn(conn)
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()

	var opt Options
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc decode options error:", err)
		return
	}

	if opt.MagicNumber != MagicNumber {
		log.Println("rpc magic number error:", opt.MagicNumber)
		return
	}

	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Println("rpc codec type error:", opt.CodecType)
		return
	}
	s.ServeCodec(f(conn))

}

var invalidRequest = struct{}{}

func (s *Server) ServeCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	for {
		req, err := s.readRequest()
		if err != nil {
			if req == nil {
				break
			}

			// 说命请求还可恢复
			s.sendResponse()
			continue
		}
		wg.Add(1)
		go s.handleRequest()
	}

	wg.Wait()
	_ = cc.Close()
}
