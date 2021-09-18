package cli

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kr/beanstalk"
)

var TooManyErrorsError = errors.New("Too many errors")

type TailCommand struct {
	Tube   string `short:"t" long:"tube" description:"tube to be tailed." required:"true"`
	Action string `short:"a" long:"action" description:"action to perform after reserving the job. (release, bury, delete)" default:"release"`

	Command
}

func (c *TailCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	return c.Tail()
}

func (c *TailCommand) Tail() error {
	ts := beanstalk.NewTubeSet(c.conn, c.Tube)

	errors := 0
	for {
		if errors > 100 {
			return TooManyErrorsError
		}

		id, body, err := ts.Reserve(time.Hour * 24)
		if err != nil {
			if err.Error() != "reserve-with-timeout: deadline soon" {
				errors++
				fmt.Println("Error", err)
			}

			continue
		}

		if err := c.PrintJob(id, body); err != nil {
			errors++
			fmt.Println("Error", err)
			continue
		}

		if err := c.postPrintAction(id); err != nil {
			return err
		}

		fmt.Println(strings.Repeat("-", 80))
	}

	return nil
}

func (c *TailCommand) postPrintAction(id uint64) error {
	var err error

	switch c.Action {
	case "release":
		err = c.conn.Release(id, 1024, 0)
	case "bury":
		err = c.conn.Bury(id, 1024)
	case "delete":
		err = c.conn.Delete(id)
	}

	return err
}
