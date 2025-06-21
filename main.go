package main

import (
	"context"
	"flag"
	"os"

	"github.com/jakub-galecki/cs_stats/faceit"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var (
	enableCache = flag.Bool("no-cache", false, "enable caching responses")
)

func main() {
	flag.Parse()

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
	if enableCache != nil && !(*enableCache) {
		s.useCache = false
	}
	if err := s.Run(context.Background()); err != nil {
		panic(err)
	}
}
