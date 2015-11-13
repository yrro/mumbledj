/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/commander_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CommanderTestSuite struct {
	Commander *Commander
	suite.Suite
}

func (suite *CommanderTestSuite) SetupTest() {
	viper.Set("aliases.add", []string{"add", "a"})
	viper.Set("aliases.cachesize", []string{"cachesize", "cs"})
	suite.Commander = NewCommander()
}

func (suite *CommanderTestSuite) TestNewCommander() {
	suite.True(len(suite.Commander.Commands) > 0, "The command list should be populated.")
}

// TODO: This test is currently broken. Update it to work with the new way commands are handled.
func (suite *CommanderTestSuite) TestFindAndExecuteCommand() {
	result, err := suite.Commander.FindCommand("Add this should find the add command!")
	suite.Equal(viper.GetStringSlice("aliases.add"), result.Aliases(), "The add command should be returned.")
	suite.Nil(err, "There shouldn't be an error.")

	result, err = suite.Commander.FindCommand("cachesize this should find the cache size command!")
	suite.Equal(viper.GetStringSlice("aliases.cachesize"), result.Aliases(), "The cachesize command should be returned.")
	suite.Nil(err, "There shouldn't be an error.")

	result, err = suite.Commander.FindCommand("This shouldn't find a command.")
	suite.Nil(result, "Result should be nil.")
	suite.NotNil(err, "There should be an error.")
}

func TestCommanderTestSuite(t *testing.T) {
	suite.Run(t, new(CommanderTestSuite))
}
