package demo

import (
	"os"

	"github.com/markus-wa/demoinfocs-golang/msg"
	d "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"github.com/rs/zerolog"
)

var ZeroDemoInfo = DemoInfo{}

type DemoInfo struct {
	Player string

	Players []string
	Game    GameStats
}

type Parser struct {
	p d.Parser
	f *os.File

	player string

	info DemoInfo

	l zerolog.Logger
}

type Opt func(*Parser)

func WithPlayer(player string) Opt {
	return func(p *Parser) {
		p.player = player
	}
}

func WithLogger(l zerolog.Logger) Opt {
	return func(p *Parser) {
		p.l = l
	}
}

func NewParser(path string, opts ...Opt) (*Parser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	p := d.NewParser(f)

	pp := &Parser{
		p: p,
		f: f,
	}

	for _, fn := range opts {
		fn(pp)
	}

	pp.register()

	pp.info.Player = pp.player

	return pp, nil
}

func (p *Parser) register() {
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

func (p *Parser) Parse() (DemoInfo, error) {
	defer func() {
		err := p.p.Close()
		if err != nil {
			p.l.Error().Err(err).Msg("error closing parser")
		}
	}()

	_, err := p.p.ParseHeader()
	if err != nil {
		return ZeroDemoInfo, err
	}

	err = p.p.ParseToEnd()
	if err != nil {
		return ZeroDemoInfo, err
	}

	return p.info, nil
}
