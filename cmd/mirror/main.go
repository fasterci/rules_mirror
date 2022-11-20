package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/fasterci/rules_mirror/pkg/mirror"
	"github.com/google/go-containerregistry/pkg/logs"
)

func init() {
	logs.Warn.SetOutput(os.Stderr)
	logs.Progress.SetOutput(os.Stderr)
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := mirror.ExecuteContext(ctx); err != nil {
		logs.Progress.Printf("Error: %v", err)
		cancel()
		os.Exit(1)
	}
}
