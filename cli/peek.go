package cli

import (
	"fmt"

	"github.com/kr/beanstalk"
)

type PeekCommand struct {
	Host  string `short:"h" long:"host" description:"beanstalkd host addr." required:"true" default:"localhost:11300"`
	Tube  string `short:"t" long:"tube" description:"tube to be tailed." required:"true"`
	State string `short:"s" long:"state" description:"peek from 'buried', 'ready' or 'delayed' queues." required:"true" default:"buried"`

	conn *beanstalk.Conn
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

func (c *PeekCommand) Init() error {
	var err error
	c.conn, err = beanstalk.Dial("tcp", c.Host)
	if err != nil {
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
