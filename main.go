package main

import (
	"os"

	commandutils "aftermath.link/repo/am-bot-stats/command-utils"
	"aftermath.link/repo/am-bot-stats/commands"
	"aftermath.link/repo/am-bot-stats/logs"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
)

type botSettings struct {
	Name   string
	Prefix string
	token  string
}

func (b *botSettings) loadToken() {
	b.token = getToken()
}

func getToken() string {
	if os.Getenv("DISCORD_TOKEN") != "" {
		return os.Getenv("DISCORD_TOKEN")
	}
	panic("DISCORD_TOKEN not set")
}

func (b *botSettings) loadSettings() {
	if os.Getenv("BOT_PREFIX") != "" {
		b.Prefix = os.Getenv("BOT_PREFIX")
	} else {
		b.Prefix = "!"
		logs.Info("BOT_PREFIX not set, using default: " + b.Prefix)
	}
	if os.Getenv("BOT_NAME") != "" {
		b.Name = os.Getenv("BOT_NAME")
	} else {
		b.Name = "GoBot"
		logs.Info("BOT_NAME not set, using default: " + b.Name)
	}
}

func init() {
	commands.Init() // Dummy handler to load commands module
}

func main() {
	var bot botSettings
	bot.loadToken()
	bot.loadSettings()

	client := disgord.New(disgord.Config{
		ProjectName: bot.Name,
		BotToken:    bot.token,
		Logger:      logs.Logger,
		RejectEvents: []string{
			// rarely used, and causes unnecessary spam
			disgord.EvtTypingStart,

			// these require special privilege
			// https://discord.com/developers/docs/topics/gateway#privileged-intents
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
		// ! Non-functional due to a current bug, will be fixed.
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "write " + bot.Prefix + "ping",
			},
		},
		DMIntents: disgord.IntentDirectMessages | disgord.IntentDirectMessageReactions | disgord.IntentDirectMessageTyping,
		// also listen for direct messages

	})

	defer client.Gateway().StayConnectedUntilInterrupted()

	logFilter, _ := std.NewLogFilter(client)
	filter, _ := std.NewMsgFilter(commandutils.RouterCtx, client)
	filter.SetPrefix(bot.Prefix)

	// create a handler and bind it to new message events
	// thing about the middlewares are whitelists or passthrough functions.
	client.Gateway().WithMiddleware(
		filter.NotByBot,    // ignore bot messages
		filter.HasPrefix,   // message must have the given prefix
		logFilter.LogMsg,   // log command message
		filter.StripPrefix, // remove the command prefix from the message
	).MessageCreate(commandutils.CommandRouter)

	// create a handler and bind it to the bot init
	// dummy log print
	client.Gateway().BotReady(func() {
		logs.Info("Bot is ready!")
	})
}
