/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numtracks_test.go
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

type NumTracksCommandTestSuite struct {
	Command NumTracksCommand
	State   *state.BotState
	suite.Suite
}

func (suite *NumTracksCommandTestSuite) SetupSuite() {
	viper.Set("aliases.numtracks", []string{"numtracks", "numsongs", "num"})
	viper.Set("descriptions.numtracks", "numtracks")
	viper.Set("permissions.numtracks", false)
}

func (suite *NumTracksCommandTestSuite) SetupTest() {
	suite.State = new(state.BotState)
	suite.State.Queue = state.NewAudioQueue()
}

func (suite *NumTracksCommandTestSuite) TestAliases() {
	suite.Equal([]string{"numtracks", "numsongs", "num"}, suite.Command.Aliases())
}

func (suite *NumTracksCommandTestSuite) TestDescription() {
	suite.Equal("numtracks", suite.Command.Description())
}

func (suite *NumTracksCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *NumTracksCommandTestSuite) TestExecute() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since the state did not change.")
	suite.NotEqual("", message, "A message should be returned.")
	suite.Contains(message, "<b>0</b> tracks", "The message should report that there are 0 tracks.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")

	suite.State.Queue.AddTracks(new(testutils.MockedAudioTrack))

	state, message, isPrivateMessage, err = suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since the state did not change.")
	suite.NotEqual("", message, "A message should be returned.")
	suite.Contains(message, "<b>1</b> track", "The message should report that there is 1 track.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")

	suite.State.Queue.AddTracks(new(testutils.MockedAudioTrack))

	state, message, isPrivateMessage, err = suite.Command.Execute(suite.State, nil)

	suite.Nil(state, "No state should be returned since the state did not change.")
	suite.NotEqual("", message, "A message should be returned.")
	suite.Contains(message, "<b>2</b> tracks", "The message should report that there are 2 tracks.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func TestNumTracksCommandTestSuite(t *testing.T) {
	suite.Run(t, new(NumTracksCommandTestSuite))
}
