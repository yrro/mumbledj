/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/nexttrack_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/state"
	"github.com/matthieugrieger/mumbledj/testutils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type NextTrackCommandTestSuite struct {
	Command NextTrackCommand
	State   *state.BotState
	suite.Suite
}

func (suite *NextTrackCommandTestSuite) SetupSuite() {
	viper.Set("aliases.nexttrack", []string{"nexttrack", "nextsong", "nt"})
	viper.Set("permissions.nexttrack", false)
}

func (suite *NextTrackCommandTestSuite) SetupTest() {
	suite.State = new(state.BotState)
	suite.State.Queue = state.NewAudioQueue()
}

func (suite *NextTrackCommandTestSuite) TestAliases() {
	suite.Equal([]string{"nexttrack", "nextsong", "nt"}, suite.Command.Aliases())
}

func (suite *NextTrackCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *NextTrackCommandTestSuite) TestExecuteWhenQueueIsEmpty() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since an error occurred.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since the queue is empty.")
}

func (suite *NextTrackCommandTestSuite) TestExecuteWhenQueueHasOneItem() {
	suite.State.Queue.AddTrack(new(testutils.MockedAudioTrack))

	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since an error occurred.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since the queue only has one track.")
}

func (suite *NextTrackCommandTestSuite) TestExecuteWhenQueueHasTwoOrMoreItems() {
	suite.State.Queue.AddTrack(new(testutils.MockedAudioTrack))
	suite.State.Queue.AddTrack(new(testutils.MockedAudioTrack))

	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since the state did not change.")
	suite.NotEqual("", message, "A message should be returned with the next track information.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func TestNextTrackCommandTestSuite(t *testing.T) {
	suite.Run(t, new(NextTrackCommandTestSuite))
}
