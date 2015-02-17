package cli

import (
	"fmt"
	"time"

	"github.com/kr/beanstalk"
)

type Tail struct {
	Host string `short:"h" long:"host" description:"beanstalkd host addr." required:"true" default:"localhost:11300"`
	Tube string `short:"t" long:"tube" description:"tube to be tailed." required:"true"`

	conn *beanstalk.Conn
}

func (t *Tail) Execute(args []string) error {
	if err := t.Init(); err != nil {
		return err
	}

	if err := t.Tail(); err != nil {
		return err
	}

	return nil
}

func (t *Tail) Init() error {
	var err error
	t.conn, err = beanstalk.Dial("tcp", t.Host)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tail) Tail() error {
	ts := beanstalk.NewTubeSet(t.conn, t.Tube)

	for {
		id, body, err := ts.Reserve(time.Hour * 24)
		if err != nil {
			return err
		}

		s, _ := t.conn.StatsJob(id)

		fmt.Printf(
			"id: %d\nlen: %d\npriority: %s\nbody: %q\n\n",
			id, len(body), s["pri"], body,
		)
	}

}
