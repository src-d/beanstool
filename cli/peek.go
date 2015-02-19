package cli

import (
	"github.com/kr/beanstalk"
)

type PeekCommand struct {
	Tube  string `short:"t" long:"tube" description:"tube to be tailed." required:"true"`
	State string `short:"" long:"state" description:"peek from 'buried', 'ready' or 'delayed' queues." default:"buried"`
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

	switch c.State {
	case "buried":
		id, body, err = t.PeekBuried()
	case "ready":
		id, body, err = t.PeekReady()
	case "delayed":
		id, body, err = t.PeekDelayed()
	}

	if err != nil {
		return err
	}

	c.PrintJob(id, body)

	return nil
}
