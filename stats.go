package main

import (
	"github.com/jakub-galecki/cs_stats/demo"
	"github.com/jakub-galecki/cs_stats/faceit"
	"github.com/rs/zerolog"
)

type stats struct {
	fc *faceit.Client
	l  zerolog.Logger
}

func (s *stats) getPlayers(info *demo.DemoInfo) []*faceit.Player {
	res := make([]*faceit.Player, 0, len(info.Players))

	for _, p := range info.Players {
		player, err := s.fc.FindPlayer(p)
		if err != nil {
			s.l.Error().Err(err).Str("name", p).Msg("error while searching for a player")
			continue
		}
		res = append(res, player)
	}

	return res
}

func (s *stats) processDemoInfo(info *demo.DemoInfo) error {
	// players := s.getPlayers(info)
	// s.l.Trace().Any("players", players).Msg("got following players for match")

	s.l.Trace().Any("kills", info.Game.Kills[info.Player]).Msg("player kills")
	s.l.Trace().Any("flashes", info.Game.Flashes[info.Player]).Msg("player flashes")
	return nil
}
