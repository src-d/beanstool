package cli

import (
	"github.com/kr/beanstalk"
)

type DeleteCommand struct {
	Tube  string `short:"t" long:"tube" description:"tube to be delete." required:"true"`
	State string `short:"" long:"state" description:"peek from 'buried', 'ready' or 'delayed' queues." default:"buried"`
	Print bool   `short:"" long:"print" description:"prints the jobs after delete it." default:"true"`
	Empty bool   `short:"" long:"empty" description:"delete all jobs with the given status in the given tube." default:"false"`
	Command
}

func (c *DeleteCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	return c.Delete()
}

func (c *DeleteCommand) Delete() error {
	t := &beanstalk.Tube{Conn: c.conn, Name: c.Tube}
	for {
		if err := c.deleteJob(t); err != nil {
			if err.Error() == "peek-ready: not found" {
				return nil
			}

			return err
		}

		if !c.Empty {
			return nil
		}
	}
}

func (c *DeleteCommand) deleteJob(t *beanstalk.Tube) error {
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

	if c.Print {
		c.PrintJob(id, body)
	}

	c.conn.Delete(id)

	return nil
}
