/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj/audiocache_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import "github.com/stretchr/testify/suite"

type AudioCacheTestSuite struct {
	Cache *AudioCache
	suite.Suite
}

/*func (suite *AudioCacheTestSuite) SetupSuite() {
	viper.Set("cache.directory", os.TempDir())
	viper.Set("cache.maximumsize", 512)
	viper.Set("cache.checkinterval", 5)
	viper.Set("cache.expiretime", 24)
}

func (suite *AudioCacheTestSuite) SetupTest() {
	suite.Cache = NewAudioCache()
}

func (suite *AudioCacheTestSuite) TestNewAudioCache() {
	suite.Zero(suite.Cache.NumAudioFiles, "There shouldn't be any cached audio files.")
	suite.Zero(suite.Cache.TotalFileSize, "There shouldn't be any cached audio files.")
}

func (suite *AudioCacheTestSuite) TestGetCurrentStatistics() {
	numFiles, totalSize := suite.Cache.GetCurrentStatistics()
	suite.Zero(numFiles, "There shouldn't be any cached audio files.")
	suite.Zero(totalSize, "There shouldn't be any cached audio files.")

	file, _ := ioutil.TempFile(os.TempDir(), "test")
	file.WriteString("Just giving this file some data!")

	numFiles, totalSize = suite.Cache.GetCurrentStatistics()
	suite.Equal(1, numFiles, "There should now be one cached file.")
	suite.True(totalSize > 0, "There should now be one cached file.")

	fileTwo, _ := ioutil.TempFile(os.TempDir(), "test2")
	fileTwo.WriteString("More data!")

	numFiles, totalSize = suite.Cache.GetCurrentStatistics()
	suite.Equal(2, numFiles, "There should now be two cached files.")
	suite.True(totalSize > 0, "There should now be two cached files.")

	os.Remove(file.Name())
	os.Remove(fileTwo.Name())
	file.Close()
	fileTwo.Close()
}

func TestAudioCacheTestSuite(t *testing.T) {
	suite.Run(t, new(AudioCacheTestSuite))
}*/
