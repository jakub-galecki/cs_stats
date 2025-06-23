package demo

import (
	"os"

	d "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/rs/zerolog"
)

type cfg struct {
	player string

	l zerolog.Logger
}

type Opt func(*cfg)

func WithPlayer(player string) Opt {
	return func(p *cfg) {
		p.player = player
	}
}

func WithLogger(l zerolog.Logger) Opt {
	return func(p *cfg) {
		p.l = l
	}
}

type parser struct {
	p d.Parser
	f *os.File
}

func internalParser(path string) (*parser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	p := d.NewParser(f)

	return &parser{
		f: f,
		p: p,
	}, nil
}
