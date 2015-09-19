package cli

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/agtorre/gocolorize"
	"github.com/mcuadros/termtable"
)

const (
	HighSeverity = iota
	NormalSeverity
	LowSeverity
	DefaultTube = "default"
)

var TubeStatsRetrievalError = errors.New("Unable to retrieve tube stats")
var HighSeverityStyle = gocolorize.NewColor("white:red")
var NormalSeverityStyle = gocolorize.NewColor("green")

type StatsCommand struct {
	Tubes string `short:"t" long:"tubes" description:"tubes to be listed (separated by ,). By default all are listed"`

	Command
}

func (c *StatsCommand) Execute(args []string) error {
	if err := c.Init(); err != nil {
		return err
	}

	if err := c.PrintStats(); err != nil {
		return err
	}

	return nil
}

func (c *StatsCommand) PrintStats() error {
	stats, err := c.GetStats()
	if err != nil {
		return err
	}

	table := termtable.NewTable(nil, &termtable.TableOptions{
		Padding:      1,
		UseSeparator: true,
	})

	table.SetHeader([]string{
		"Name", "Buried", "Delayed", "Ready", "Reserved", "Urgent", "Waiting", "Total",
	})

	table.AddRow(c.buildLineFromTubeStats(DefaultTube, stats[DefaultTube]))

	for _, t := range sortedKeys(stats) {
		if t == DefaultTube {
			continue
		}

		table.AddRow(c.buildLineFromTubeStats(t, stats[t]))
	}

	fmt.Println(table.Render())
	return nil
}

func (c *StatsCommand) buildLineFromTubeStats(name string, s *TubeStats) []string {
	var l []string

	l = append(l, name)
	l = append(l, addStyle(s.JobsBuried, 8, HighSeverity))
	l = append(l, addStyle(s.JobsDelayed, 8, NormalSeverity))
	l = append(l, addStyle(s.JobsReady, 8, NormalSeverity))
	l = append(l, addStyle(s.JobsReserved, 8, NormalSeverity))
	l = append(l, addStyle(s.JobsUrgent, 8, NormalSeverity))
	l = append(l, addStyle(s.Waiting, 8, LowSeverity))
	l = append(l, addStyle(s.TotalJobs, 8, LowSeverity))

	return l
}

func addStyle(i int, l int, severity int) string {
	value := strconv.Itoa(i)
	needs := l - len(value)
	if needs <= 0 {
		return value
	}

	padded := value + strings.Repeat(" ", needs)
	if i > 0 {
		switch severity {
		case HighSeverity:
			return HighSeverityStyle.Paint(padded)
		case NormalSeverity:
			return NormalSeverityStyle.Paint(padded)
		}
	}

	return padded
}

func (c *StatsCommand) GetStats() (map[string]*TubeStats, error) {
	tubes, err := c.getTubesName()
	if err != nil {
		return nil, err
	}

	stats := make(map[string]*TubeStats, 0)
	for _, tube := range tubes {
		s, err := c.GetStatsForTube(tube)
		if err != nil {
			return nil, err
		}

		stats[tube] = s
	}

	return stats, nil
}

func (c *StatsCommand) getTubesName() ([]string, error) {
	if c.Tubes != "" {
		return strings.Split(strings.Replace(c.Tubes, " ", "", -1), ","), nil
	}

	return c.conn.ListTubes()
}

func mustConvertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func sortedKeys(m map[string]*TubeStats) []string {
	keys := make([]string, len(m))

	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	return keys
}
