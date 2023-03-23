package main

import (
	"flag"

	"github.com/smallnest/rpcx/server"
	"golang/rpc/proto"
)


var (
    addr = flag.String("addr", "localhost:8972", "server address")
)


func main() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(proto.Arith), "")
	s.Serve("tcp", *addr)
}
