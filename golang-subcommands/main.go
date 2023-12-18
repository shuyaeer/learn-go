package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/shuyaeer/golang-subcommands/cmd"
)

func main() {
  subcommands.Register(subcommands.HelpCommand(), "")
  subcommands.Register(subcommands.FlagsCommand(), "")
  subcommands.Register(subcommands.CommandsCommand(), "")
  subcommands.Register(&cmd.ShowCSVContentCmd{}, "")

  flag.Parse()
  ctx := context.Background()
  os.Exit(int(subcommands.Execute(ctx)))
}