package cli

import (
	"fmt"
	"time"

	"github.com/kr/beanstalk"
)

type PutCommand struct {
	Tube     string        `short:"t" long:"tube" description:"tube to be tailed." required:"true"`
	Body     string        `short:"b" long:"body" description:"plain text data for the job." required:"true"`
	Priority uint32        `short:"p" long:"priority" description:"priority for the job." default:"1024"`
	Delay    time.Duration `short:"d" long:"delay" description:"delay for the job." default:"0"`
	TTR      time.Duration `short:"" long:"ttr" description:"TTR for the job." default:"60s"`

	Command
}

func (c *PutCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	return c.Put()
}

func (c *PutCommand) Put() error {
	t := beanstalk.Tube{Conn: c.conn, Name: c.Tube}

	id, err := t.Put([]byte(c.Body), c.Priority, c.Delay, c.TTR)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Added job with id %d to %s with priority %d, delay %s, TTR %d\n",
		id, c.Tube, c.Priority, c.Delay.String(), c.TTR,
	)

	return nil
}
