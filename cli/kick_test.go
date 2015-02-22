package cli

import (
	"github.com/kr/beanstalk"
	. "gopkg.in/check.v1"
)

type KickCommandSuite struct {
	c *KickCommand
	t *beanstalk.Tube
}

var _ = Suite(&KickCommandSuite{})

func (s *KickCommandSuite) SetUpTest(c *C) {
	s.c = &KickCommand{}
	s.c.Host = "localhost:11300"
	s.c.Init()

	s.t = getRandomTube(s.c.conn)
	s.c.Tube = s.t.Name
}

func (s *KickCommandSuite) TestKickCommand_Kick(c *C) {
	s.t.Put([]byte(""), 1024, 0, 0)

	bury := &BuryCommand{}
	bury.conn = s.c.conn
	bury.Tube = s.c.Tube
	bury.Bury()

	stats, _ := s.c.GetStatsForTube(s.c.Tube)
	c.Assert(stats.JobsBuried, Equals, 1)

	err := s.c.Kick()
	c.Assert(err, IsNil)

	stats, _ = s.c.GetStatsForTube(s.c.Tube)
	c.Assert(stats.JobsBuried, Equals, 0)
	c.Assert(stats.JobsReady, Equals, 1)
}
