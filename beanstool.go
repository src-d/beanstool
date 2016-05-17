package main

import (
	"os"

	"github.com/agtorre/gocolorize"
	"github.com/jessevdk/go-flags"
	"github.com/src-d/beanstool/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		gocolorize.SetPlain(true)
	}

	parser := flags.NewNamedParser("beanstool", flags.Default)
	parser.AddCommand("stats", "print stats on all tubes", "", &cli.StatsCommand{})
	parser.AddCommand("tail", "tails a tube and prints his content", "", &cli.TailCommand{})
	parser.AddCommand("peek", "peeks a job from a queue", "", &cli.PeekCommand{})
	parser.AddCommand("delete", "delete a job from a queue", "", &cli.DeleteCommand{})
	parser.AddCommand("kick", "kicks jobs from buried back into ready", "", &cli.KickCommand{})
	parser.AddCommand("put", "put a job into a tube", "", &cli.PutCommand{})
	parser.AddCommand("bury", "bury existing jobs from ready state", "", &cli.BuryCommand{})

	_, err := parser.Parse()
	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}
