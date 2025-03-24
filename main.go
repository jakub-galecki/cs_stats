package main

import (
	"context"
	"os"

	"github.com/jakub-galecki/cs_stats/faceit"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var (
		auth = os.Getenv("AUTH")
		base = os.Getenv("BASE")
	)
	ctx := context.WithValue(context.Background(), "key", auth)
	fc := faceit.NewClient(base, ctx)
	logger := zerolog.New(os.Stdout).Level(zerolog.TraceLevel).With().Timestamp().Logger()
	stats := &stats{
		fc: fc,
		l:  logger,
	}
	s := NewServer(stats, logger)
	if err := s.Run(context.Background()); err != nil {
		panic(err)
	}
}
