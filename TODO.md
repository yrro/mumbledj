MumbleDJ Refactor To-do List
============================

* Commands
  * add
    * [x] Implementation
    * [x] Unit tests
  * cachesize
    * [x] Implementation
    * [ ] Unit tests
  * currenttrack
    * [x] Implementation
    * [x] Unit tests
  * forceskip
    * [ ] Implementation
    * [ ] Unit tests
  * forceskipplaylist
    * [ ] Implementation
    * [ ] Unit tests
  * help
    * [ ] Implementation
    * [ ] Unit tests
  * kill
    * [ ] Implementation
    * [ ] Unit tests
  * listtracks
    * [ ] Implementation
    * [ ] Unit tests
  * move
    * [ ] Implementation
    * [ ] Unit tests
  * nexttrack
    * [x] Implementation
    * [x] Unit tests
  * numcached
    * [x] Implementation
    * [ ] Unit tests
  * numtracks
    * [x] Implementation
    * [x] Unit tests
  * reload
    * [ ] Implementation
    * [ ] Unit tests
  * reset
    * [ ] Implementation
    * [ ] Unit tests
  * setcomment
    * [ ] Implementation
    * [ ] Unit tests
  * shuffle
    * [x] Implementation
    * [x] Unit tests
  * shuffleoff
    * [x] Implementation
    * [x] Unit tests
  * shuffleon
    * [x] Implementation
    * [x] Unit tests
  * skip
    * [ ] Implementation
    * [ ] Unit tests
  * skipplaylist
    * [ ] Implementation
    * [ ] Unit tests
  * volume
    * [x] Implementation
    * [x] Unit tests
* State
  * AudioCache
    * [x] Implementation
    * [ ] Unit tests
  * AudioQueue
    * [x] Implementation
    * [x] Unit tests
  * BotState
    * [x] Implementation
  * SkipTracker
    * [x] Implementation
    * [x] Unit tests
* DJ
  * Commander
    * [x] Implementation
    * [ ] Unit tests
  * Config
    * [x] Default values
    * [ ] Commandline overrides
    * [ ] Environment variable overrides
    * [ ] Config file overrides
    * [ ] Unit tests
  * MumbleDJ
    * [ ] Startup checks
    * [ ] Mumble server connection
    * [x] Gumble event listener implementation
* Audio
  * YouTube
    * [ ] Tracks
    * [ ] Playlists
  * SoundCloud
    * [ ] Tracks
    * [ ] Playlists
  * AudioHandler
    * [ ] Implement GetTracks
    * [ ] Unit tests
* README
  * [ ] Reorganize
  * [ ] Simplify installation instructions
  * [ ] Simplify command listing
* CONTRIBUTING
  * [ ] Create CONTRIBUTING.md
  * [ ] Create detailed instructions for adding new commands
  * [ ] Create detailed instructions for proper formatting
* Installation
  * [ ] Attempt to support `go install`
  * [ ] Attempt to support Docker installs
  * [ ] Attempt to make compilation process simpler
* Automation
  * [ ] Integrate a continuous integration tool such as Travis CI
  * [ ] Integrate an automated coverage tool such as Coveralls
* Gumble
  * [ ] Update to comply with new API changes
