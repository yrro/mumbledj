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
    * [x] Implementation
    * [ ] Unit tests
  * forceskipplaylist
    * [x] Implementation
    * [ ] Unit tests
  * help
    * [x] Implementation
    * [x] Unit tests
  * kill
    * [x] Implementation
    * [ ] Unit tests
  * listtracks
    * [x] Implementation
    * [ ] Unit tests
  * move
    * [x] Implementation
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
    * [x] Implementation
    * [ ] Unit tests
  * reset
    * [x] Implementation
    * [ ] Unit tests
  * setcomment
    * [x] Implementation
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
    * [x] Implementation
    * [ ] Unit tests
  * skipplaylist
    * [x] Implementation
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
    * [x] Commandline overrides
    * [ ] Environment variable overrides
    * [x] Config file overrides
    * [ ] Unit tests
  * MumbleDJ
    * [ ] Startup checks
    * [x] Mumble server connection
    * [x] Gumble event listener implementation
* Audio
  * YouTube
    * [x] Tracks
    * [x] Playlists
  * SoundCloud
    * [x] Tracks
    * [x] Playlists
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
