package demo

type Tick struct {
	Player string
	Pos    struct {
		X float64
		Y float64
	}
}

type TickHandler func(data any) error

type Player struct {
	*parser
	cfg

	Map string

	handler TickHandler
}

func NewPlayer(path string, eventHandler TickHandler, opts ...Opt) (*Player, error) {
	p := &Player{}

	internal, err := internalParser(path)
	if err != nil {
		return nil, err
	}

	p.parser = internal
	p.handler = eventHandler

	for _, fn := range opts {
		fn(&p.cfg)
	}

	return p, nil
}

func (pl *Player) Parse() error {
	defer func() {
		err := pl.p.Close()
		if err != nil {
			pl.l.Error().Err(err).Msg("error closing parser")
		}
	}()

	h, err := pl.p.ParseHeader()
	if err != nil {
		return err
	}

	pl.l.Trace().Any("header", h).Msg("parsed demo header")

	for {
		more, err := pl.p.ParseNextFrame()
		if !more {
			return nil
		}

		if err != nil {
			return err
		}
		gs := pl.p.GameState()
		// if !gs.IsMatchStarted() || !gs.IsFreezetimePeriod() || !gs.IsWarmupPeriod() {
		// 	continue
		// }
		t := Tick{}
		for _, p := range gs.Participants().Playing() {
			if p.Name != pl.cfg.player {
				continue
			}

			// todo: add noise cancel
			pos := p.Position()
			t.Player = p.Name
			t.Pos.X = pos.X
			t.Pos.Y = pos.Y
		}

		tick := pl.p.CurrentFrame()

		pl.l.Trace().Float64("x", t.Pos.X).Float64("y", t.Pos.Y).Int("tick", tick).Msg("generating tick")

		if err := pl.handler(t); err != nil {
			pl.l.Error().Err(err).Msg("handling frame failed")
		}
	}
}
