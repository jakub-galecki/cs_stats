package demo

import (
	"os"

	"github.com/markus-wa/demoinfocs-golang/msg"
	d "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/msgs2"
	"github.com/rs/zerolog"
)

var ZeroDemoInfo = DemoInfo{}

type DemoInfo struct {
	Player string
	Map    string

	Players []string
	Game    GameStats
}

type Info struct {

	player string

	info DemoInfo

	l zerolog.Logger
}

type Opt func(*Info)

func WithPlayer(player string) Opt {
	return func(p *Info) {
		p.player = player
	}
}

func WithLogger(l zerolog.Logger) Opt {
	return func(p *Info) {
		p.l = l
	}
}

func NewInfo(path string, opts ...Opt) (*Info, error) {
	pp := &Info{}

	internal, err := internalParser(path)
	if err != nil {
		return nil, err
	}

	pp/.

	for _, fn := range opts {
		fn(pp)
	}

	pp.register()

	pp.info.Player = pp.player

	return pp, nil
}

func (p *Info) register() {
	p.p.RegisterNetMessageHandler(func(m *msgs2.CSVCMsg_ServerInfo) {
		name := m.GetMapName()
		p.l.Trace().Str("map_name", name).Msg("decoding map name")
		p.info.Map = name
	})

	p.p.RegisterEventHandler(func(e events.PlayerConnect) {
		p.l.Trace().Str("name", e.Player.Name).Msg("appending player")
		p.info.Players = append(p.info.Players, e.Player.Name)
	})

	p.p.RegisterEventHandler(func(e events.Kill) {
		if e.Killer == nil {
			p.l.Debug().Msg("player was killed but killer was not provided")
			return
		}
		p.info.Game.Add(e)
	})

	p.p.RegisterEventHandler(func(e events.PlayerFlashed) {
		if e.Attacker == nil || e.Player == nil {
			p.l.Debug().Msg("player was flashed but info was not provided")
			return
		}
		p.info.Game.Add(e)
	})

	p.p.RegisterNetMessageHandler(func(msg *msg.CSVCMsg_ServerInfo) {
		if msg == nil {
			return
		}

		p.info.Game.Add(msg)
	})

	// p.p.RegisterEventHandler(func(e events.HeExplode) {
	// 	events.
	// })
}

//func getBestAction()

func (p *Info) Parse() (DemoInfo, error) {
	defer func() {
		err := p.p.Close()
		if err != nil {
			p.l.Error().Err(err).Msg("error closing parser")
		}
	}()

	h, err := p.p.ParseHeader()
	if err != nil {
		return ZeroDemoInfo, err
	}

	p.l.Trace().Any("header", h).Msg("parsed demo header")

	err = p.p.ParseToEnd()
	if err != nil {
		return ZeroDemoInfo, err
	}

	return p.info, nil
}
