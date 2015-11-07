/*
 * MumbleDJ
 * By Matthieu Grieger
 * state/audioqueue.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import (
	"errors"
	"math/rand"
	"time"

	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/spf13/viper"
)

// AudioQueue holds the audio queue itself along with useful methods for
// performing actions on the queue.
type AudioQueue struct {
	Queue []audio.Track
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NewAudioQueue initializes a new queue and returns it.
func NewAudioQueue() *AudioQueue {
	return &AudioQueue{
		Queue: make([]audio.Track, 0),
	}
}

// AddTrack adds a Track to the AudioQueue.
func (q *AudioQueue) AddTrack(t audio.Track) error {
	beforeLen := len(q.Queue)
	q.Queue = append(q.Queue, t)
	if len(q.Queue) == beforeLen+1 {
		return nil
	}
	return errors.New("Could not add Track to the AudioQueue.")
}

// CurrentTrack returns the current Track.
func (q *AudioQueue) CurrentTrack() (audio.Track, error) {
	if len(q.Queue) != 0 {
		return q.Queue[0], nil
	}
	return nil, errors.New("There are no tracks in the AudioQueue.")
}

// PeekNextTrack peeks at the next Track and returns it.
func (q *AudioQueue) PeekNextTrack() (audio.Track, error) {
	if len(q.Queue) > 1 {
		if viper.GetBool("general.automaticshuffleon") {
			q.RandomNextTrack(false)
		}
		return q.Queue[1], nil
	}
	return nil, errors.New("There isn't a Track coming up next.")
}

// Traverse is a traversal function for AudioQueue. Allows a visit function to
// be passed in which performs the specified action on each queue item.
func (q *AudioQueue) Traverse(visit func(i int, t audio.Track)) {
	for tQueue, queueTrack := range q.Queue {
		visit(tQueue, queueTrack)
	}
}

// ShuffleTracks shuffles the AudioQueue using an inside-out algorithm.
func (q *AudioQueue) ShuffleTracks() {
	for i := range q.Queue[1:] { // Don't touch Track that is currently playing.
		j := rand.Intn(i + 1)
		q.Queue[i+1], q.Queue[j+1] = q.Queue[j+1], q.Queue[i+1]
	}
}

// RandomNextTrack sets a random Track as the next Track to be played.
func (q *AudioQueue) RandomNextTrack(queueWasEmpty bool) {
	if len(q.Queue) > 1 {
		nextTrackIndex := 1
		if queueWasEmpty {
			nextTrackIndex = 0
		}
		swapIndex := nextTrackIndex + rand.Intn(len(q.Queue)-1)
		q.Queue[nextTrackIndex], q.Queue[swapIndex] = q.Queue[swapIndex], q.Queue[nextTrackIndex]
	}
}
