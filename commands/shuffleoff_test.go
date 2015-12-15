/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffleoff_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ShuffleOffCommandTestSuite struct {
	Command ShuffleOffCommand
	suite.Suite
}

func (suite *ShuffleOffCommandTestSuite) SetupSuite() {
	viper.Set("aliases.shuffleoff", []string{"shuffleoff", "shufoff", "shoff"})
}

func (suite *ShuffleOffCommandTestSuite) TestAliases() {
	suite.Equal([]string{"shuffleoff", "shufoff", "shoff"}, suite.Command.Aliases())
}

func (suite *ShuffleOffCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *ShuffleOffCommandTestSuite) TestExecuteWhenAutomaticShuffleOn() {
	viper.Set("general.automaticshuffleon", true)

	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil)

	suite.Nil(state, "This command shouldn't return a state.")
	suite.NotEqual("", message, "A message should be returned as automatic shuffling has been successfully toggled off.")
	suite.False(isPrivateMessage, "This message should not be private.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *ShuffleOffCommandTestSuite) TestExecuteWhenAutomaticShuffleOff() {
	viper.Set("general.automaticshuffleon", false)

	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil)

	suite.Nil(state, "This command shouldn't return a state.")
	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since automatic shuffling is already toggled off.")
}

func TestShuffleOffCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ShuffleOffCommandTestSuite))
}
