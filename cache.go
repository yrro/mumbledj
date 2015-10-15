/*
 * MumbleDJ
 * By Matthieu Grieger
 * cache.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"io/ioutil"
	"time"
)

// AudioCache keeps track of the filesize of the audio cache and provides
// methods for pruning the cache.
type AudioCache struct {
	NumAudioFiles int
	TotalFileSize int64
	MaximumSize   int64
	Directory     string
	ExpireTime    float64
	CheckInterval int
}

// NewAudioCache creates an empty AudioCache.
func NewAudioCache(maximumSize int64, directory string, expireTime float64, checkInterval int) *AudioCache {
	newCache := &AudioCache{
		NumAudioFiles: 0,
		TotalFileSize: 0,
		MaximumSize:   maximumSize,
		Directory:     directory,
		ExpireTime:    expireTime,
		CheckInterval: checkInterval,
	}
	return newCache
}

// GetCurrentStatistics retrieves the total file size and number of files stored in the
// cache and updates the member variables accordingly.
func (c *AudioCache) GetCurrentStatistics() (int, int64) {
	var totalSize int64
	files, _ := ioutil.ReadDir(c.Directory)
	for _, file := range files {
		totalSize += file.Size()
	}
	return len(files), totalSize
}

// CheckDirectorySize checks the cache directory to determine if the filesize
// of the files within exceed the user-specified size limit. If so, the oldest files
// are cleared until it is no longer exceeding the limit.
func (c *AudioCache) CheckDirectorySize() {
	c.UpdateStats()
	for c.TotalFileSize > (c.MaximumSize * BytesInMebibyte) {
		if err := c.RemoveOldest(); err != nil {
			break
		}
	}
}

// UpdateStats updates the statistics relevant to the cache (number of audio files cached,
// total current size of the cache).
func (c *AudioCache) UpdateStats() {
	c.NumAudioFiles, c.TotalFileSize = c.GetCurrentStatistics()
}

// CleanPeriodically loops forever, cleaning expired cached audio files as necessary.
func (c *AudioCache) CleanPeriodically() {
	for range time.Tick(time.Duration(c.CheckInterval) * time.Minute) {
		files, _ := ioutil.ReadDir(c.Directory)
		for _, file := range files {
			// NOTE: Some sort of safe-guard must be put in place to make
			// sure that a file isn't deleted when it is being actively played.
			hours := time.Since(file.ModTime()).Hours()
			// TODO: Finish the rest of this function.
		}
	}
}

// RemoveOldest deletes the oldest file in the cache.
func (c *AudioCache) RemoveOldest() error {
	files, _ := ioutil.ReadDir(c.Directory)
	// TODO: Finish the rest of this function, add sort type.
}
