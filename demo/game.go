package demo

import (
	"github.com/jakub-galecki/cs_stats/pb"
	"github.com/markus-wa/demoinfocs-golang/msg"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type KillCounter struct {
	Count         int
	HeadshotCount int
}

func (kc *KillCounter) Add(e events.Kill) {
	kc.Count++
	if e.IsHeadshot {
		kc.HeadshotCount++
	}
}

type Kill struct {
	*KillCounter
	pm *PositionsManager

	Terrorist         *KillCounter `json:",omitempty"`
	CounterTerrorists *KillCounter `json:",omitempty"`

	Gun map[string]*KillCounter `json:",omitempty"`
}

type Flash struct {
	EnemiesFlashed int
}

type GameStats struct {
	// nickname -> kill
	Kills   map[string]*Kill
	Flashes map[string]*Flash
}

func (k *Kill) Add(e events.Kill) {
	if k == nil {
		return
	}

	if k.KillCounter == nil {
		k.KillCounter = &KillCounter{}
	}

	k.KillCounter.Add(e)

	switch e.Killer.Team {
	case common.TeamCounterTerrorists:
		if k.CounterTerrorists == nil {
			k.CounterTerrorists = &KillCounter{}
		}
		k.CounterTerrorists.Add(e)
	case common.TeamTerrorists:
		if k.Terrorist == nil {
			k.Terrorist = &KillCounter{}
		}
		k.Terrorist.Add(e)
	}

	if e.Weapon != nil {
		tp := e.Weapon.Type.String()
		if k.Gun == nil {
			k.Gun = make(map[string]*KillCounter)
		}
		if k.Gun[tp] == nil {
			k.Gun[tp] = &KillCounter{}
		}
		k.Gun[tp].Add(e)
	}

	if k.pm == nil {
		k.pm = NewPositionsManager()
	}

	k.pm.Add(e)
}

func (gs *GameStats) addKill(e events.Kill) {
	nick := e.Killer.Name

	if gs.Kills == nil {
		gs.Kills = make(map[string]*Kill)
	}

	_, ok := gs.Kills[nick]
	if !ok {
		gs.Kills[nick] = &Kill{}
	}

	gs.Kills[nick].Add(e)
}

func (gs *GameStats) addFlashed(e events.PlayerFlashed) {
	if e.Attacker == nil || e.Player == nil {
		return
	}

	if gs.Flashes == nil {
		gs.Flashes = make(map[string]*Flash)
	}

	nick := e.Attacker.Name
	if gs.Flashes[nick] == nil {
		gs.Flashes[nick] = &Flash{}
	}
	gs.Flashes[nick].EnemiesFlashed++
}

func (gs *GameStats) Add(e any) {
	if e == nil {
		return
	}

	switch e := e.(type) {
	case events.Kill:
		gs.addKill(e)
	case events.PlayerFlashed:
		gs.addFlashed(e)
	case *msg.CSVCMsg_ServerInfo:
	}
}

func (gs *GameStats) ToProto() *pb.GameStats {
	var (
		kills   = make(map[string]*pb.Kill)
		flashes = make(map[string]*pb.Flash)
	)

	for player, kill := range gs.Kills {
		kills[player] = kill.ToProto()
	}

	for player, flash := range gs.Flashes {
		flashes[player] = &pb.Flash{
			EnemiesFlashed: int32(flash.EnemiesFlashed),
		}
	}

	return &pb.GameStats{
		Kills:   kills,
		Flashes: flashes,
	}
}

func (kc *KillCounter) ToProto() *pb.KillCounter {
	if kc == nil {
		return nil
	}
	return &pb.KillCounter{
		Count:         int32(kc.Count),
		HeadshotCount: int32(kc.HeadshotCount),
	}
}

func (k *Kill) ToProto() *pb.Kill {
	weapons := make(map[string]*pb.KillCounter)
	for gun, counter := range k.Gun {
		if counter != nil {
			weapons[gun] = counter.ToProto()
		}
	}

	return &pb.Kill{
		Overall:           k.KillCounter.ToProto(),
		Terrorists:        k.Terrorist.ToProto(),
		CounterTerrorists: k.CounterTerrorists.ToProto(),
		Weapons:           weapons,
		KillPositions:     k.pm.ToProto(),
	}
}
