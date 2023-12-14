package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/shuyaeer/learn-go/app/day13/cmd"
)

func main() {
  subcommands.Register(subcommands.HelpCommand(), "")
  subcommands.Register(subcommands.FlagsCommand(), "")
  subcommands.Register(subcommands.CommandsCommand(), "")
  subcommands.Register(&cmd.PrintCmd{}, "")

  flag.Parse()
  ctx := context.Background()
  os.Exit(int(subcommands.Execute(ctx)))
}