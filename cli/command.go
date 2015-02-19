package cli

import (
	"fmt"

	"github.com/agtorre/gocolorize"
	"github.com/kr/beanstalk"
)

var TitleStyle = gocolorize.NewColor("green")
var InfoStyle = gocolorize.NewColor("yellow")

type Command struct {
	Host string `short:"h" long:"host" description:"beanstalkd host addr." required:"true" default:"localhost:11300"`

	conn *beanstalk.Conn
}

func (c *Command) Init() error {
	var err error
	c.conn, err = beanstalk.Dial("tcp", c.Host)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) PrintJob(id uint64, body []byte) error {
	s, err := c.conn.StatsJob(id)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s: %d, %s: %d, %s: %s, %s: %s, %s: %s, %s: %s\n"+
			"%s: %s, %s: %s, %s: %s, %s: %s, %s: %s\n"+
			"%s:\n%q\n",
		TitleStyle.Paint("id"), id,
		TitleStyle.Paint("length"), len(body),
		TitleStyle.Paint("priority"), s["pri"],
		TitleStyle.Paint("delay"), s["delay"],
		TitleStyle.Paint("age"), s["age"],
		TitleStyle.Paint("ttr"), s["ttr"],

		InfoStyle.Paint("reserves"), s["reserves"],
		InfoStyle.Paint("releases"), s["releases"],
		InfoStyle.Paint("buries"), s["buries"],
		InfoStyle.Paint("kicks"), s["kicks"],
		InfoStyle.Paint("timeouts"), s["timeouts"],

		InfoStyle.Paint("body"), body,
	)

	return nil
}

func (c *Command) GetStatsForTube(tube string) (*TubeStats, error) {
	t := &beanstalk.Tube{c.conn, tube}
	s, err := t.Stats()
	if err != nil {
		return nil, err
	}

	if name, ok := s["name"]; !ok || name != tube {
		return nil, TubeStatsRetrievalError
	}

	return &TubeStats{
		JobsBuried:   mustConvertToInt(s["current-jobs-buried"]),
		JobsReady:    mustConvertToInt(s["current-jobs-ready"]),
		JobsDelayed:  mustConvertToInt(s["current-jobs-delayed"]),
		JobsReserved: mustConvertToInt(s["current-jobs-reserved"]),
		JobsUrgent:   mustConvertToInt(s["current-jobs-urgent"]),
		Waiting:      mustConvertToInt(s["current-waiting"]),
		TotalJobs:    mustConvertToInt(s["total-jobs"]),
	}, nil
}

type TubeStats struct {
	JobsBuried   int
	JobsDelayed  int
	JobsReady    int
	JobsReserved int
	JobsUrgent   int
	Waiting      int
	TotalJobs    int
}
