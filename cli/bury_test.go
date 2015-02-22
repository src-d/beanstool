package cli

import (
	"github.com/kr/beanstalk"
	. "gopkg.in/check.v1"
)

type BuryCommandSuite struct {
	c *BuryCommand
	t *beanstalk.Tube
}

var _ = Suite(&BuryCommandSuite{})

func (s *BuryCommandSuite) SetUpTest(c *C) {
	s.c = &BuryCommand{}
	s.c.Host = "localhost:11300"
	s.c.Init()

	s.t = getRandomTube(s.c.conn)
	s.c.Tube = s.t.Name
}

func (s *BuryCommandSuite) TestBuryCommand_Bury(c *C) {
	s.t.Put([]byte(""), 1024, 0, 0)

	err := s.c.Bury()
	c.Assert(err, IsNil)

	stats, _ := s.c.GetStatsForTube(s.c.Tube)
	c.Assert(stats.JobsBuried, Equals, 1)
}

func (s *BuryCommandSuite) TestBuryCommand_BuryWithLimitUnder(c *C) {
	s.t.Put([]byte(""), 1024, 0, 0)
	s.t.Put([]byte(""), 1024, 0, 0)

	s.c.Num = 1
	err := s.c.Bury()
	c.Assert(err, IsNil)

	stats, _ := s.c.GetStatsForTube(s.c.Tube)
	c.Assert(stats.JobsBuried, Equals, 1)
	c.Assert(stats.JobsReady, Equals, 1)
}

func (s *BuryCommandSuite) TestBuryCommand_BuryWithLimitOver(c *C) {
	s.t.Put([]byte(""), 1024, 0, 0)

	s.c.Num = 10
	err := s.c.Bury()
	c.Assert(err, IsNil)

	stats, _ := s.c.GetStatsForTube(s.c.Tube)
	c.Assert(stats.JobsBuried, Equals, 1)
	c.Assert(stats.JobsReady, Equals, 0)
}
