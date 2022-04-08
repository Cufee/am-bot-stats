package router

import (
	"context"

	"github.com/byvko-dev/am-core/discord/router/errors"
	"github.com/byvko-dev/am-types/users/v2"
)

func runCommandHandler(command Command) {
	// Add a reaction to the message
	go func() {
		if command.message != nil {
			command.message.React(context.Background(), command.session, "☑️")
		}
	}()

	// Get user profile
	user, err := command.userCheck(command.userID)
	command.User = user
	if err != nil {
		sendError(command, errors.ErrUserCheckFailed)
		return
	}

	// Check if user is banned
	if command.User.Ban.Active {
		sendError(command, errors.ErrUserBanned)
		return
	}

	// Check for command restrictions
	if command.options.Permissions.GuildOnly && (command.message == nil || command.message.GuildID == 0) {
		sendError(command, errors.ErrCommandIsGuildOnly)
		return
	}
	if command.options.Permissions.DirectOnly && (command.message == nil || command.message.GuildID != 0) {
		sendError(command, errors.ErrCommandIsDirectOnly)
		return
	}
	if command.options.Permissions.AdminOnly && !command.User.Features.Enabled.Includes(users.DiscordAdminCommands) {
		sendError(command, errors.ErrCommandIsAdminOnly)
		return
	}

	// Execute command handler
	err = command.options.Handler(command)
	if err != nil {
		sendError(command, err)
	}
}

func sendError(command Command, err error) {
	if command.reportError != nil {
		var guildId string = "0"
		var channelId string = "0"
		if command.message != nil {
			guildId = command.message.GuildID.String()
			channelId = command.message.ChannelID.String()
		}
		ctx := ErrorContext{
			Command:   command.Name,
			Arguments: command.Arguments,
			UserID:    command.userID,
			GuildID:   guildId,
			ChannelID: channelId,
			User:      command.User,
		}
		defer command.reportError(ctx, err)
	}

	if command.options.ErrorHandler != nil {
		command.options.ErrorHandler(command, err)
	} else {
		command.Reply(command, command.errorPrinter(command.User.Locale, err))
	}
}
