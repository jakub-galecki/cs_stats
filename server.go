package main

import (
	"context"
	"net/http"

	"github.com/jakub-galecki/cs_stats/demo"
	"github.com/jakub-galecki/cs_stats/pb"
	"github.com/rs/zerolog"
)

type Server struct {
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
	p, err := demo.NewParser(req.Path, demo.WithPlayer(req.Player), demo.WithLogger(s.l))
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

func (s *Server) Run(ctx context.Context) error {
	ss := pb.NewStatsServer(s)
	return http.ListenAndServe(":8080", ss)
}
