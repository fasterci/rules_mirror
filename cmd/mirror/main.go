package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/fasterci/rules_mirror/pkg/mirror"
	"github.com/google/go-containerregistry/pkg/logs"
)

var (
	Timeout time.Duration
)

func init() {
	logs.Warn.SetOutput(os.Stderr)
	logs.Progress.SetOutput(os.Stderr)
	flag.DurationVar(&Timeout, "timeout", time.Second*30, "Timeout for the mirror operation")
}

func main() {
	flag.Parse()
	ctx, cancelT := context.WithTimeout(context.Background(), Timeout)
	defer cancelT()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	if err := mirror.ExecuteContext(ctx); err != nil {
		logs.Progress.Printf("Error: %v", err)
		cancel()
		os.Exit(1)
	}
}
