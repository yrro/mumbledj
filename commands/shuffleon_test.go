/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffleon_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ShuffleOnCommandTestSuite struct {
	Command ShuffleOnCommand
	suite.Suite
}

func (suite *ShuffleOnCommandTestSuite) SetupSuite() {
	viper.Set("aliases.shuffleon", []string{"shuffleon", "shufon", "shon"})
	viper.Set("descriptions.shuffleon", "shuffleon")
	viper.Set("permissions.shuffleon", true)
}

func (suite *ShuffleOnCommandTestSuite) TestAliases() {
	suite.Equal([]string{"shuffleon", "shufon", "shon"}, suite.Command.Aliases())
}

func (suite *ShuffleOnCommandTestSuite) TestDescription() {
	suite.Equal("shuffleon", suite.Command.Description())
}

func (suite *ShuffleOnCommandTestSuite) TestIsAdmin() {
	suite.True(suite.Command.IsAdmin())
}

func (suite *ShuffleOnCommandTestSuite) TestExecuteWhenAutomaticShuffleOn() {
	viper.Set("general.automaticshuffleon", true)

	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil)

	suite.Nil(state, "This command shouldn't return a state.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since automatic shuffling is already toggled on.")
}

func (suite *ShuffleOnCommandTestSuite) TestExecuteWhenAutomaticShuffleOff() {
	viper.Set("general.automaticshuffleon", false)

	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil)

	suite.Nil(state, "This command shouldn't return a state.")
	suite.NotEqual("", message, "A message should be returned as automatic shuffling has been successfully toggled on.")
	suite.False(isPrivateMessage, "This message should not be private.")
	suite.Nil(err, "No error should be returned.")
}

func TestShuffleOnCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ShuffleOnCommandTestSuite))
}
