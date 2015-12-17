/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/add_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/matthieugrieger/mumbledj/testutils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type AddCommandTestSuite struct {
	Command AddCommand
	State   *state.BotState
	suite.Suite
}

func (suite *AddCommandTestSuite) SetupSuite() {
	viper.Set("aliases.add", []string{"add", "a"})
	viper.Set("permissions.add", false)
}

func (suite *AddCommandTestSuite) SetupTest() {
	suite.State = new(state.BotState)
	suite.State.Queue = state.NewAudioQueue()
	suite.State.Handler = new(testutils.MockedAudioHandler)
}

func (suite *AddCommandTestSuite) TestAliases() {
	suite.Equal([]string{"add", "a"}, suite.Command.Aliases())
}

func (suite *AddCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *AddCommandTestSuite) TestExecuteWithNoArgs() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since an error occurred.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned for attempting to add a track without providing a URL.")
}

func (suite *AddCommandTestSuite) TestExecuteWhenNoTracksFound() {
	// NOTE: The third argument ("0") tells the MockedAudioHandler to return 0 tracks
	// during the execution of the add command.
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil, "0")

	suite.Nil(state, "No state should be returned since an error occurred.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned for providing a URL that did not have any tracks associated with it.")
}

func (suite *AddCommandTestSuite) TestExecuteWhenTrackFound() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, new(gumble.User), "1")

	suite.NotNil(state, "A new state should be returned since it changed.")
	suite.Equal(1, len(state.Queue.Queue), "There should be one item in the queue.")
	suite.NotEqual("", message, "A message should be returned.")
	suite.False(isPrivateMessage, "This message shouldn't be private.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *AddCommandTestSuite) TestExecuteWhenMultipleTracksFound() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, new(gumble.User), "3")

	suite.NotNil(state, "A new state should be returned since it changed.")
	suite.Equal(3, len(state.Queue.Queue), "There should be three items in the queue.")
	suite.NotEqual("", message, "A message should be returned.")
	suite.False(isPrivateMessage, "This message shouldn't be private.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *AddCommandTestSuite) TestExecuteWithMultipleURLs() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, new(gumble.User), "0", "1", "2")

	suite.NotNil(state, "A new state should be returned since it changed.")
	suite.Equal(3, len(state.Queue.Queue), "There should be three items in the queue.")
	suite.NotEqual("", message, "A message should be returned.")
	suite.False(isPrivateMessage, "This message shouldn't be private.")
	suite.Nil(err, "No error should be returned.")
}

func TestAddCommandTestSuite(t *testing.T) {
	suite.Run(t, new(AddCommandTestSuite))
}
