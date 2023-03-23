package main

import (
    "context"
    "flag"
    "log"
	"golang/rpc/proto"

    "github.com/smallnest/rpcx/client"
)

var (
    addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
    flag.Parse()

    d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
    xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
    defer xclient.Close()

    args := &proto.Args{
        A: 10,
        B: 0,
    }

    reply := &proto.Reply{}
    err := xclient.Call(context.Background(), "Div", args, reply)
    if err != nil {
        log.Fatalf("failed to call: %v", err)
    }

    log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
