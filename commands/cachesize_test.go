/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/cachesize_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CacheSizeCommandTestSuite struct {
	Command CacheSizeCommand
	suite.Suite
}

func (suite *CacheSizeCommandTestSuite) SetupSuite() {
	viper.Set("aliases.cachesize", []string{"cachesize", "cs"})
	viper.Set("descriptions.cachesize", "cachesize")
	viper.Set("permissions.cachesize", true)
}

func (suite *CacheSizeCommandTestSuite) TestAliases() {
	suite.Equal([]string{"cachesize", "cs"}, suite.Command.Aliases())
}

func (suite *CacheSizeCommandTestSuite) TestDescription() {
	suite.Equal("cachesize", suite.Command.Description())
}

func (suite *CacheSizeCommandTestSuite) TestIsAdmin() {
	suite.True(suite.Command.IsAdmin())
}

func (suite *CacheSizeCommandTestSuite) TestExecuteWhenCachingIsDisabled() {
	viper.Set("cache.enabled", false)
	state, message, isPrivateMessage, err := suite.Command.Execute(nil, nil)

	suite.Nil(state, "An error occurred so no state should be returned.")
	suite.Equal("", message, "An error occurred so no message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned because caching is disabled.")
}

// TODO: Implement TestExecuteWhenCachingIsEnabled() test.
func (suite *CacheSizeCommandTestSuite) TestExecuteWhenCachingIsEnabled() {

}

func TestCacheSizeCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CacheSizeCommandTestSuite))
}
