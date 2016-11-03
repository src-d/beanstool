package cli

import (
	"fmt"

	"github.com/kr/beanstalk"
)

type KickCommand struct {
	Tube string `short:"t" long:"tube" description:"tube to kick jobs in." required:"true"`
	Num  int    `short:"" long:"num" description:"number of jobs to kick."`
	Command
}

func (c *KickCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	return c.Kick()
}

func (c *KickCommand) Kick() error {
	if err := c.calcNumIfNeeded(); err != nil {
		return err
	}

	if c.Num == 0 {
		fmt.Printf("Empty buried queue at tube %q.\n", c.Tube)
		return nil
	}

	fmt.Printf("Trying to kick %d jobs from %q ...\n", c.Num, c.Tube)

	t := &beanstalk.Tube{Conn: c.conn, Name: c.Tube}
	kicked, err := t.Kick(c.Num)
	if err != nil {
		return err
	}

	fmt.Printf("Actually kicked %d.\n", kicked)
	return nil
}

func (c *KickCommand) calcNumIfNeeded() error {
	if c.Num == 0 {
		s, err := c.GetStatsForTube(c.Tube)
		if err != nil {
			return err
		}

		c.Num = s.JobsBuried
	}

	return nil
}
