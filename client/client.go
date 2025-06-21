package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jakub-galecki/cs_stats/pb"
	"storj.io/drpc/drpcconn"
)

func main() {
	rawconn, err := net.Dial("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	conn := drpcconn.New(rawconn)
	defer conn.Close()

	client := pb.NewDRPCStatsClient(conn)
	res, err := client.ProcessDemo(context.Background(), &pb.ProcessDemoReq{
		Path:   "/mnt/c/Users/galec/Documents/1-39d11ddf-1b84-4e7a-b5eb-3cd76a918986-1-1.dem",
		Player: "jac30b",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
