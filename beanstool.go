package main

import (
	"os"

	"github.com/tyba/beanstool/cli"

	"github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewNamedParser("event-relay", flags.Default)
	parser.AddCommand("stats", "print stats on all tubes", "", &cli.Monitor{})
	parser.AddCommand("tail", "tails a tube and prints his content", "", &cli.Tail{})

	_, err := parser.Parse()
	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			parser.WriteHelp(os.Stdout)
		}

		os.Exit(1)
	}
}
