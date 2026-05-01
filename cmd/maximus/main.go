package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var version = "dev" // overridden by ldflags at build time

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCmd(os.Args[2:])
	case "version":
		fmt.Println(version)
	default:
		usage()
		os.Exit(1)
	}
}

func runCmd(args []string) {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	configPath := fs.String("config", "config.toml", "path to TOML config file")
	_ = fs.Parse(args)

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, *configPath); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

// run is the entry point for the pipeline. It is separate from runCmd so it
// can be tested without os.Exit.
func run(ctx context.Context, configPath string) error {
	panic("not implemented")
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: maximus <command> [flags]

Commands:
  run      Start the replication pipeline
  version  Print the version and exit

Run 'maximus run --help' for flags.
`)
}
