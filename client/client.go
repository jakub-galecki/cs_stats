package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jakub-galecki/cs_stats/pb"
)

func main() {
	client := pb.NewStatsJSONClient("http://localhost:8080", &http.Client{})
	res, err := client.ProcessDemo(context.Background(), &pb.ProcessDemoReq{
		Path:   "/Users/batman/Downloads/1-0a4e7ada-05a4-4004-91f5-74462dc9e9a0-1-1.dem",
		Player: "jac30b",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
