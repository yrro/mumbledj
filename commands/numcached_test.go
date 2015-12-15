/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numcached_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type NumCachedCommandTestSuite struct {
	Command NumCachedCommand
	suite.Suite
}

func (suite *NumCachedCommandTestSuite) SetupSuite() {
	viper.Set("aliases.numcached", []string{"numcached", "nc"})
	viper.Set("permissions.numcached", true)
}

func (suite *NumCachedCommandTestSuite) TestAliases() {
	suite.Equal([]string{"numcached", "nc"}, suite.Command.Aliases())
}

func (suite *NumCachedCommandTestSuite) TestIsAdmin() {
	suite.True(suite.Command.IsAdmin())
}

func (suite *NumCachedCommandTestSuite) TestExecuteWhenCachingIsDisabled() {
	viper.Set("cache.enabled", false)
	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil)

	suite.Nil(state, "An error occurred so no state should be returned.")
	suite.Equal("", message, "An error occurred so no message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned because caching is disabled.")
}

// TODO: Implement TestExecuteWhenCachingIsEnabled() test.
func (suite *NumCachedCommandTestSuite) TestExecuteWhenCachingIsEnabled() {

}

func TestNumCachedCommandTestSuite(t *testing.T) {
	suite.Run(t, new(NumCachedCommandTestSuite))
}
