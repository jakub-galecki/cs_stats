package demo

import (
	"github.com/golang/geo/r2"
	"github.com/jakub-galecki/cs_stats/pb"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type PositionsManager struct {
	All []r2.Point
}

func NewPositionsManager() *PositionsManager {
	return &PositionsManager{}
}

func (p *PositionsManager) Add(e events.Kill) {
	if e.Killer == nil || e.Killer.Entity == nil {
		return
	}
	pos := e.Killer.Position()
	p.All = append(p.All, r2.Point{X: pos.X, Y: pos.Y})
}

func (p *PositionsManager) ToProto() []*pb.Position {
	var res []*pb.Position
	for _, pos := range p.All {
		res = append(res, &pb.Position{X: pos.X, Y: pos.Y})
	}
	return res
}
