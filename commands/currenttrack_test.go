/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/currenttrack_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CurrentTrackCommandTestSuite struct {
	Command CurrentTrackCommand
	State   *state.BotState
	suite.Suite
}

func (suite *CurrentTrackCommandTestSuite) SetupSuite() {
	viper.Set("aliases.currenttrack", []string{"currenttrack", "current"})
	viper.Set("permissions.currenttrack", false)
}

func (suite *CurrentTrackCommandTestSuite) SetupTest() {
	suite.State = new(state.BotState)
	suite.State.Queue = state.NewAudioQueue()
}

func (suite *CurrentTrackCommandTestSuite) TestAliases() {
	suite.Equal([]string{"currenttrack", "current"}, suite.Command.Aliases())
}

func (suite *CurrentTrackCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *CurrentTrackCommandTestSuite) TestExecuteWhenQueueIsEmpty() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since an error occurred.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since the queue is empty.")
}

// TODO: Implement TestExecuteWhenQueueNotEmpty() test.
func (suite *CurrentTrackCommandTestSuite) TestExecuteWhenQueueNotEmpty() {

}

func TestCurrentTrackCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CurrentTrackCommandTestSuite))
}
