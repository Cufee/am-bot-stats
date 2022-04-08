package router

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-core/helpers/slices"
	"github.com/byvko-dev/am-core/logs"
)

// Register a command handler to the command router where name is the command prefix (case insensitive)
func RegisterCommand(opts RegisterOptions) {
	// Check if a command or alias is already registered
	if slices.Contains(boundCommands, fmt.Sprintf("%v-prefix", opts.Name)) > -1 {
		logs.Error(fmt.Sprintf("Command %v is already bound", fmt.Sprintf("%v-prefix", opts.Name)))
		return
	}
	for _, a := range opts.Aliases {
		if slices.Contains(boundCommands, fmt.Sprintf("%v-prefix", a)) > -1 {
			logs.Error(fmt.Sprintf("Alias %v is already bound", fmt.Sprintf("%v-prefix", a)))
			return
		}
		boundCommands = append(boundCommands, fmt.Sprintf("%v-prefix", a))
	}
	boundCommands = append(boundCommands, fmt.Sprintf("%v-prefix", opts.Name))

	// Register command
	validCommands = append(validCommands, opts)
}

func (c *Command) SetSession(session disgord.Session) {
	c.session = session
}
func (c *Command) SetMessage(message *disgord.Message) {
	c.message = message
}

func (c *Command) UnsafeGetSession() disgord.Session {
	return c.session
}

func (c *Command) UnsafeGetMessage() *disgord.Message {
	return c.message
}
