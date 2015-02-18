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
