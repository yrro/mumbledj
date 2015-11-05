/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/commander_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommander(t *testing.T) {
	commander := NewCommander()
	assert.True(t, len(commander.Commands) > 0, "The command list should be populated.")
}
