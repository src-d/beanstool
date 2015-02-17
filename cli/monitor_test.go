package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	m := &Monitor{Host: "localhost:11300"}
	err := m.Init()
	assert.Nil(t, err)

	m.PrintStats()

}
