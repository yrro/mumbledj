/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/command_executor.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import "github.com/matthieugrieger/mumbledj/mumbledj"

// CommandExecutor contains logic that parses user text messages for commands
// and executes the commands.
type CommandExecutor struct {
	Commands []Command
	Bot      *mumbledj.MumbleDJ
}
