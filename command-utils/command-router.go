package commandutils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"aftermath.link/repo/am-bot-stats/external"
	"aftermath.link/repo/am-bot-stats/localization"
	"aftermath.link/repo/am-bot-stats/utils"
	"github.com/andersfylling/disgord"
)

type CommandHandler struct {
	Command string
	Aliases []string
	Handler func(Command) error
	OnError func(Command, error)
	// Run options
	opts RunOptions
}
type RunOptions struct {
	AdminOnly  bool
	GuildOnly  bool
	DirectOnly bool
}

var validCommands []CommandHandler
var boundCommands []string

// Register a command handler to the command router where name is the command prefix (case insensitive)
func RegisterCommand(name string, opts RunOptions, handler func(Command) error, errorHandler func(Command, error), alias ...string) {
	// Check if a command or alias is already registered
	if utils.ContainsString(boundCommands, name) {
		panic(fmt.Sprintf("Command %v is already bound", name))
	}
	for _, a := range alias {
		if utils.ContainsString(boundCommands, a) {
			panic(fmt.Sprintf("Alias %v is already bound", a))
		}
	}

	// Register command
	validCommands = append(validCommands, CommandHandler{strings.ToLower(name), alias, handler, errorHandler, opts})
}

// Main command router
var RouterCtx = context.Background()

type Command struct {
	Name      string
	Arguments []string
	User      external.UserProfile
	session   disgord.Session
	message   *disgord.Message
}

func (c *Command) SetSession(session disgord.Session) {
	c.session = session
}
func (c *Command) SetMessage(message *disgord.Message) {
	c.message = message
}
func (c *Command) Reply(content ...interface{}) error {
	_, err := c.message.Reply(RouterCtx, c.session, content...)
	return err
}

func CommandRouter(s disgord.Session, data *disgord.MessageCreate) {
	var messageCommand string
	var messageArguments []string
	if data.Message.Content == "" {
		return
	}
	messageArgumentsSlice := strings.Split(data.Message.Content, " ")
	messageCommand = strings.ToLower(messageArgumentsSlice[0])
	messageArguments = messageArgumentsSlice[1:]

	var command Command
	command.SetSession(s)
	command.SetMessage(data.Message)
	command.Name = messageCommand
	command.Arguments = messageArguments

	// Find a command handler and execute it
	for _, cmd := range validCommands {
		if command.Name == cmd.Command {
			runCommandHandler(cmd, command)
			return
		}
		if len(cmd.Aliases) > 0 {
			if utils.ContainsString(cmd.Aliases, command.Name) {
				runCommandHandler(cmd, command)
				return
			}
		}
	}

	// Invalid command
	commandNotFoundHandler(command)
}

func runCommandHandler(cmd CommandHandler, command Command) {
	// Get user profile
	user, err := external.CheckUserByUserID(command.message.Author.ID.String())
	command.User = user
	if err != nil {
		localization.ReplyToErorr(err, command.User.Locale, command.Reply)
		return
	}

	// Check if user is banned
	if command.User.Banned || command.User.ShadowBanned {
		userBannedHandler(command)
		return
	}

	// Check for command restrictions
	if cmd.opts.GuildOnly && command.message.GuildID == 0 || cmd.opts.DirectOnly && command.message.GuildID != 0 {
		commandRestrictedError(command, cmd.opts)
		return
	}

	// Execute command handler
	err = cmd.Handler(command)
	if err != nil {
		if cmd.OnError != nil {
			cmd.OnError(command, err)
		} else {
			localization.ReplyToErorr(err, command.User.Locale, command.Reply)
		}
	}
}

func commandNotFoundHandler(command Command) {
	localization.ReplyToErorr(errors.New(localization.ErrUnknownCommandKeyword), command.User.Locale, command.Reply)
}

func commandRestrictedError(command Command, opts RunOptions) {
	if opts.DirectOnly {
		localization.ReplyToErorr(errors.New(localization.ErrCommandIsDirectOnlyKeyword), command.User.Locale, command.Reply)
		return
	}
	if opts.GuildOnly {
		localization.ReplyToErorr(errors.New(localization.ErrCommandIsGuildOnlyKeyword), command.User.Locale, command.Reply)
		return
	}
	// opts.AdminOnly will also fall here
	commandNotFoundHandler(command)
}

func userBannedHandler(command Command) {
	if command.User.BanNotified {
		return
	}
	command.User.BanNotified = true
	// Update user profile

	localization.ReplyToErorr(errors.New(localization.ErrUserBannedKeyword), command.User.Locale, command.Reply, command.User.BanReason)
}
