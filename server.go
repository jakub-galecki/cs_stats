package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jakub-galecki/cs_stats/demo"
	"github.com/jakub-galecki/cs_stats/pb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedStatsServer

	l     zerolog.Logger
	stats *stats

	useCache bool
	cache    map[string]*pb.ProcessDemoResponse
}

func NewServer(stats *stats, l zerolog.Logger) *Server {
	return &Server{
		cache: make(map[string]*pb.ProcessDemoResponse),
		l:     l,
		stats: stats,
	}
}

func (s *Server) ProcessDemo(ctx context.Context, req *pb.ProcessDemoReq) (*pb.ProcessDemoResponse, error) {
	p, err := demo.NewInfo(req.Path, demo.WithPlayer(req.Player), demo.WithLogger(s.l))
	if err != nil {
		return nil, err
	}

	if res, ok := s.cache[req.Path]; ok {
		return res, nil
	}

	info, err := p.Parse()
	if err != nil {
		return nil, err
	}

	err = s.stats.processDemoInfo(&info)
	if err != nil {
		return nil, err
	}

	res := &pb.ProcessDemoResponse{
		Player:  info.Player,
		Map:     info.Map,
		Players: info.Players,
		Stats:   info.Game.ToProto(),
	}
	if s.useCache {
		s.cache[req.Path] = res
	}
	return res, nil
}

func (s *Server) tickHandler(dst grpc.ServerStreamingServer[pb.Pos]) demo.TickHandler {
	return func(data any) error {
		t, ok := data.(demo.Tick)
		if !ok {
			return fmt.Errorf("invalid data type")
		}

		return dst.Send(&pb.Pos{
			X: t.Pos.X,
			Y: t.Pos.Y,
		})
	}
}

func (s *Server) ReplayDemo(req *pb.ProcessDemoReq, res grpc.ServerStreamingServer[pb.Pos]) error {
	dem, err := demo.NewPlayer(req.Path, s.tickHandler(res), demo.WithPlayer(req.Player), demo.WithLogger(s.l))
	if err != nil {
		return err
	}

	if err = dem.Parse(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		return err
	}
	defer listener.Close()

	gs := grpc.NewServer()
	pb.RegisterStatsServer(gs, s)
	reflection.Register(gs)
	return gs.Serve(listener)
}
