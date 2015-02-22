package cli

import (
	"time"

	"github.com/kr/beanstalk"
	. "gopkg.in/check.v1"
)

type PutCommandSuite struct {
	c *PutCommand
	t *beanstalk.Tube
}

var _ = Suite(&PutCommandSuite{})

func (s *PutCommandSuite) SetUpTest(c *C) {
	s.c = &PutCommand{}
	s.c.Host = "localhost:11300"
	s.c.Init()

	s.t = getRandomTube(s.c.conn)
	s.c.Tube = s.t.Name
}

func (s *PutCommandSuite) TestPutCommand_Put(c *C) {
	s.c.Body = "foo"
	s.c.Priority = 1024
	s.c.Delay = time.Second

	err := s.c.Put()
	c.Assert(err, IsNil)

	stats, _ := s.c.GetStatsForTube(s.c.Tube)
	c.Assert(stats.JobsDelayed, Equals, 1)
}
