package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kr/beanstalk"
)

type BuryCommand struct {
	Tube string `short:"t" long:"tube" description:"tube to bury jobs in." required:"true"`
	Num  int    `short:"n" long:"num" description:"number of jobs to bury."`
	Command
}

func (c *BuryCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	if err := c.Bury(); err != nil {
		return err
	}

	return nil
}

func (c *BuryCommand) Bury() error {
	if err := c.calcNumIfNeeded(); err != nil {
		return err
	}

	if c.Num == 0 {
		fmt.Printf("Empty ready queue at tube %q.\n", c.Tube)
		return nil
	}

	fmt.Printf("Trying to bury %d jobs from %q...\n", c.Num, c.Tube)

	count := 0
	ts := beanstalk.NewTubeSet(c.conn, c.Tube)
	for count < c.Num {
		id, _, err := ts.Reserve(time.Hour * 24)
		if err != nil {
			return err
		}

		s, err := c.conn.StatsJob(id)
		if err != nil {
			return err
		}

		pri, err := strconv.Atoi(s["pri"])
		if err != nil {
			return err
		}

		if err := c.conn.Bury(id, uint32(pri)); err != nil {
			return err
		}
		fmt.Println(pri)
		count++
	}

	fmt.Printf("Actually buried %d.\n", count)
	return nil
}

func (c *BuryCommand) calcNumIfNeeded() error {
	if c.Num == 0 {
		s, err := c.GetStatsForTube(c.Tube)
		if err != nil {
			return err
		}

		c.Num = s.JobsReady
	}

	return nil
}
