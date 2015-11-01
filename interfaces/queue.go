/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/queue.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package interfaces

// Queue is an interface that allows access to audio tracks within a queue.
type Queue interface {
	AddTrack(t Track) error
	CurrentTrack() (Track, error)
	PeekNextTrack() (Track, error)
	Traverse(visit func(i int, t Track))
	ShuffleTracks()
}
