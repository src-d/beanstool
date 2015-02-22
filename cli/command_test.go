package cli

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/kr/beanstalk"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CommandSuite struct{}

var _ = Suite(&CommandSuite{})

func (s *CommandSuite) TestCommand_GetStatsForTube(c *C) {
	cmd := &Command{Host: "localhost:11300"}
	cmd.Init()

	tube := getRandomTube(cmd.conn)
	tube.Put([]byte(""), 1024, 0, 0)

	stats, err := cmd.GetStatsForTube(tube.Name)
	c.Assert(err, IsNil)
	c.Assert(stats.JobsReady, Equals, 1)
	c.Assert(stats.JobsBuried, Equals, 0)
}

func getRandomTube(conn *beanstalk.Conn) *beanstalk.Tube {
	rb := make([]byte, 32)
	if _, err := rand.Read(rb); err != nil {
		panic(err)
	}

	name := strings.Replace(base64.URLEncoding.EncodeToString(rb), "=", "0", -1)

	return &beanstalk.Tube{conn, name}
}
