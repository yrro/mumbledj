/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/help_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type HelpCommandTestSuite struct {
	Command HelpCommand
	suite.Suite
}

func (suite *HelpCommandTestSuite) SetupSuite() {
	viper.Set("aliases.help", []string{"help", "h"})
	viper.Set("descriptions.help", "help")
	viper.Set("permissions.help", false)
	viper.Set("permissions.adminsenabled", true)
	viper.Set("permissions.admins", []string{"Admin"})
}

func (suite *HelpCommandTestSuite) TestAliases() {
	suite.Equal([]string{"help", "h"}, suite.Command.Aliases())
}

func (suite *HelpCommandTestSuite) TestDescription() {
	suite.Equal("help", suite.Command.Description())
}

func (suite *HelpCommandTestSuite) TestIsAdmin() {
	suite.False(suite.Command.IsAdmin())
}

func (suite *HelpCommandTestSuite) TestExecute() {
	adminUser := new(gumble.User)
	adminUser.Name = "Admin"

	regularUser := new(gumble.User)
	regularUser.Name = "NotAdmin"

	state1, message1, isPrivateMessage1, err1 := suite.Command.Execute(nil, regularUser)

	suite.Nil(state1, "No state should be returned since the state did not change.")
	suite.NotEqual("", message1, "A message should be returned with the command descriptions.")
	suite.True(isPrivateMessage1, "This should be a private message.")
	suite.Nil(err1, "No error should be returned.")

	state2, message2, isPrivateMessage2, err2 := suite.Command.Execute(nil, adminUser)

	suite.Nil(state2, "No state should be returned since the state did not change.")
	suite.NotEqual("", message2, "A message should be returned with the command descriptions.")
	suite.True(isPrivateMessage2, "This should be a private message.")
	suite.Nil(err2, "No error should be returned.")

	suite.True(len(message2) > len(message1), "The response for an admin should have more commands listed.")
}

func TestHelpCommandTestSuite(t *testing.T) {
	suite.Run(t, new(HelpCommandTestSuite))
}
