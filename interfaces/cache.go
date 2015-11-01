/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/datastore.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package interfaces

// Cache is an interface that defines methods for retrieving audio tracks
// that are stored within a cache.
type Cache interface {
	NumFiles() int
	TotalSize() int64
	MaximumSize() int64
	Directory() string
}
