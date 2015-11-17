/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/commander_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/commands"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CommanderTestSuite struct {
	Commander *Commander
	suite.Suite
}

type MockedAddCommand struct {
	commands.AddCommand
	mock.Mock
}

func (suite *CommanderTestSuite) SetupTest() {
	viper.Set("aliases.add", []string{"add", "a"})
	viper.Set("aliases.cachesize", []string{"cachesize", "cs"})
	suite.Commander = NewCommander()
}

func (suite *CommanderTestSuite) TestNewCommander() {
	suite.True(len(suite.Commander.Commands) > 0, "The command list should be populated.")
}

func (suite *CommanderTestSuite) TestFindCommand() {
	command, err := suite.Commander.FindCommand("add this should return the add command!")
	suite.Equal(new(commands.AddCommand), command, "This input should return the add command.")
	suite.Nil(err, "A command was found, so an error shouldn't be returned.")

	command, err = suite.Commander.FindCommand("CacheSize this should return the cachesize command.")
	suite.Equal(new(commands.CacheSizeCommand), command, "This input should return the cachesize command.")
	suite.Nil(err, "A command was found, so an error shouldn't be returned.")

	command, err = suite.Commander.FindCommand("fjdkasfjaskf this should not return a command.")
	suite.Nil(command, "This input should not return any command, so command should be nil.")
	suite.NotNil(err, "No command was found, so an error should be returned.")
}

// TODO: Implement TestExecuteCommand() test.
func (suite *CommanderTestSuite) TestExecuteCommand() {

}

func TestCommanderTestSuite(t *testing.T) {
	suite.Run(t, new(CommanderTestSuite))
}
