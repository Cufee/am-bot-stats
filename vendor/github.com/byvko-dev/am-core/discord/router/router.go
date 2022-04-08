package router

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-core/helpers/slices"
)

var validCommands []RegisterOptions
var boundCommands []string

// Main command router
var RouterCtx = context.Background()

func CommandRouter(options RouterOptions) func(s disgord.Session, data *disgord.MessageCreate) {
	options.check() // Panic if some required handlers are missing

	return func(s disgord.Session, data *disgord.MessageCreate) {
		var messageCommand string
		var messageArguments []string
		if data.Message.Content == "" {
			return
		}
		messageArgumentsSlice := strings.Split(data.Message.Content, " ")
		messageCommand = strings.ToLower(messageArgumentsSlice[0])
		if len(messageArgumentsSlice) > 1 {
			messageArguments = messageArgumentsSlice[1:]
		}

		var command Command
		command.SetSession(s)
		command.SetMessage(data.Message)
		command.Name = messageCommand
		command.Arguments = messageArguments
		command.userID = data.Message.Author.ID.String()
		command.userCheck = options.UserCheckHandler
		command.reportError = options.ErrorReportHandler
		command.errorPrinter = options.ErrorPrinter

		command.Reply = func(c Command, content ...interface{}) error {
			_, err := c.message.Reply(RouterCtx, c.session, content...)
			return err
		}

		// Find a command handler and execute it
		for _, cmd := range validCommands {
			if !cmd.AddPrefixCommand {
				continue
			}
			if command.Name == strings.ToLower(cmd.Name) || (len(cmd.Aliases) > 0 && slices.Contains(cmd.Aliases, strings.ToLower(command.Name)) > -1) {
				command.options = cmd
				runCommandHandler(command)
			}
		}
	}
	// If a command was not found, it is likely registered in another module
}
