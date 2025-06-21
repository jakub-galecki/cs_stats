package demo

import (
	"os"

	d "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
)

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
