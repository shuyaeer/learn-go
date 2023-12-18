package cmd

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type ShowCSVContentCmd struct {
	show string
}

func (*ShowCSVContentCmd) Name() string     { return "show" } // サブコマンド名指定
func (*ShowCSVContentCmd) Synopsis() string { return "show args csv content to stdout." }
func (*ShowCSVContentCmd) Usage() string {
	return `show [-capitalize] <some text>:
	Show args csv content to stdout.
`
}

func (p *ShowCSVContentCmd) SetFlags(f *flag.FlagSet) {
}

func (p *ShowCSVContentCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Printf("Failed to open file: %v\n", err)
			return subcommands.ExitFailure
		}
		defer file.Close()

		reader := csv.NewReader(file)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Failed to read CSV record: %v\n", err)
				return subcommands.ExitFailure
			}
			fmt.Println(strings.Join(record, ","))
		}
	}
	return subcommands.ExitSuccess
}

