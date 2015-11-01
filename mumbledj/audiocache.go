/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj/audiocache.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package mumbledj

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

// SortFilesByAge is a type that holds file information for cached items.
type SortFilesByAge []os.FileInfo

// Len returns the length of the file slice.
func (a SortFilesByAge) Len() int {
	return len(a)
}

// Swap swaps two elements in the file slice.
func (a SortFilesByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less compares two file modification times to determine if one is less than
// the other. Returns true if the item in index i is older than the item in
// index j, false otherwise.
func (a SortFilesByAge) Less(i, j int) bool {
	return time.Since(a[i].ModTime()) < time.Since(a[j].ModTime())
}

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
			// It is safe to check the modification time because when audio
			// files are played their modification time is updated. This ensures
			// that audio files will not get deleted while they are playing, assuming
			// a reasonable expiry time is set in the configuration.
			hours := time.Since(file.ModTime()).Hours()
			if hours >= c.ExpireTime {
				os.Remove(fmt.Sprintf("%s/%s", c.Directory, file.Name()))
			}
		}
	}
}

// RemoveOldest deletes the oldest file in the cache.
func (c *AudioCache) RemoveOldest() error {
	files, _ := ioutil.ReadDir(c.Directory)
	if len(files) > 0 {
		sort.Sort(SortFilesByAge(files))
		os.Remove(fmt.Sprintf("%s/%s", c.Directory, files[0].Name()))
		return nil
	}
	return errors.New("There are no files currently cached.")
}

// NumFiles returns the number of audio files currently cached.
func (c *AudioCache) NumFiles() int {
	return c.NumAudioFiles
}

// TotalSize returns the total file size of the cache in bytes.
func (c *AudioCache) TotalSize() int64 {
	return c.TotalFileSize
}

// MaximumSize returns the maximum file size of the cache in bytes.
func (c *AudioCache) MaximumSize() int64 {
	return c.MaximumSize
}

// Directory returns the directory in which audio files are cached.
func (c *AudioCache) Directory() string {
	return c.Directory
}
