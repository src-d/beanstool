package cli

import (
	"github.com/kr/beanstalk"
)

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
