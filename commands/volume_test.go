/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/volume_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"strconv"
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleffmpeg"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type VolumeCommandTestSuite struct {
	Command VolumeCommand
	State   *state.BotState
	User    *gumble.User
	suite.Suite
}

func (suite *VolumeCommandTestSuite) SetupSuite() {
	viper.Set("volume.lowest", 0.01)
	viper.Set("volume.highest", 0.8)
	viper.Set("volume.default", 0.4)
	viper.Set("aliases.volume", []string{"volume", "v"})
	viper.Set("descriptions.volume", "volume")
	viper.Set("permissions.volume", false)
}

func (suite *VolumeCommandTestSuite) SetupTest() {
	suite.State = new(state.BotState)
	suite.State.AudioStream = new(gumbleffmpeg.Stream)
	suite.State.AudioStream.Volume = float32(viper.GetFloat64("volume.default"))

	suite.User = new(gumble.User)
	suite.User.Name = "Test"
}

func (suite *VolumeCommandTestSuite) TestAliases() {
	suite.Equal([]string{"volume", "v"}, suite.Command.Aliases())
}

func (suite *VolumeCommandTestSuite) TestDescription() {
	suite.Equal("volume", suite.Command.Description())
}

func (suite *VolumeCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *VolumeCommandTestSuite) TestExecuteWithValidVolume() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, suite.User, "0.3")

	suite.NotNil(state, "No error occurred, so the returned state should not be nil.")
	suite.NotEqual("", message, "No error occurred, so the returned message should not be empty.")
	suite.False(isPrivateMessage, "This message should not be private.")
	suite.Nil(err, "No error occurred, so the returned error should be nil.")
	suite.Equal(float32(0.3), state.AudioStream.Volume, "The returned state should have the new volume assigned to it.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithExtraArguments() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, suite.User, "0.3", "extra")

	suite.NotNil(state, "No error occurred, so the returned state should be nil.")
	suite.NotEqual("", message, "No error occurred, so the returned message should not be empty.")
	suite.False(isPrivateMessage, "This message should not be private.")
	suite.Nil(err, "No error occurred, so the returned error should be nil.")
	suite.Equal(float32(0.3), state.AudioStream.Volume, "The returned state should have the new volume assigned to it.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithNoArguments() {
	state, message, isPrivateMessage, err := suite.Command.Execute(suite.State, suite.User)

	suite.Nil(state, "This command does not alter the bot's state, so a new state shouldn't be returned.")
	suite.NotEqual("", message, "No error occurred, so the returned message should not be empty.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error occurred, so the returned error should be nil.")
	suite.Contains(message, strconv.FormatFloat(viper.GetFloat64("volume.default"), 'f', 2, 32), "The returned message should contain the current volume.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithVolumeTooLow() {
	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil, "0.0001")

	suite.Nil(state, "An error occurred so no state should be returned.")
	suite.Equal("", message, "An error occurred so no message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned for providing a volume that is too low.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithVolumeTooHigh() {
	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil, "1.0")

	suite.Nil(state, "An error occurred so no state should be returned.")
	suite.Equal("", message, "An error occurred so no message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned for providing a volume that is too high.")
}

func TestVolumeCommandTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeCommandTestSuite))
}
