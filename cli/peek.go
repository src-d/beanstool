package cli

import (
	"fmt"

	"github.com/kr/beanstalk"
)

type PeekCommand struct {
	Tube  string `short:"t" long:"tube" description:"tube to be tailed." required:"true"`
	State string `short:"s" long:"state" description:"peek from 'buried', 'ready' or 'delayed' queues." required:"true" default:"buried"`
	Command
}

func (c *PeekCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	if err := c.Peek(); err != nil {
		return err
	}

	return nil
}

func (c *PeekCommand) Peek() error {
	t := &beanstalk.Tube{c.conn, c.Tube}
	var id uint64
	var body []byte
	var err error
	id, body, err = t.PeekBuried()
	if err != nil {
		return err
	}

	fmt.Printf(
		"id: %d\nlen: %d\nbody: %q\n\n",
		id, len(body), body,
	)

	return nil
}
