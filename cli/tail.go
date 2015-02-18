package cli

import (
	"fmt"
	"time"

	"github.com/kr/beanstalk"
)

type TailCommand struct {
	Tube string `short:"t" long:"tube" description:"tube to be tailed." required:"true"`
	Command
}

func (c *TailCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	if err := c.Tail(); err != nil {
		return err
	}

	return nil
}

func (c *TailCommand) Tail() error {
	ts := beanstalk.NewTubeSet(c.conn, c.Tube)

	for {
		id, body, err := ts.Reserve(time.Hour * 24)
		if err != nil {
			return err
		}

		s, _ := c.conn.StatsJob(id)

		fmt.Printf(
			"id: %d\nlen: %d\npriority: %s\nbody: %q\n\n",
			id, len(body), s["pri"], body,
		)
	}

}
