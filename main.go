package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aftermath.link/repo/am-bot-stats/commands"

	"aftermath.link/repo/am-bot-stats/config"
	"aftermath.link/repo/am-bot-stats/logs"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/byvko-dev/am-core/discord/router"
	"github.com/byvko-dev/am-core/helpers/env"
)

type botSettings struct {
	Prefix string
	token  string
	status string
}

func (b *botSettings) loadToken() {
	b.token = config.DiscordToken
}

func (b *botSettings) loadSettings() {
	b.Prefix = config.DiscordCommandPrefix
	b.status = config.DiscordBotStatus
}

func loadShardingConfig() disgord.ShardConfig {
	envs := env.MustGet("SHARDS_COUNT", "SHARDS_LIST")
	shards, err := strconv.Atoi(envs[0].(string))
	if err != nil {
		panic(err)
	}
	servingShardsStr := strings.Split(envs[1].(string), ",")
	servingShards := make([]uint, len(servingShardsStr))
	for i, v := range servingShardsStr {
		s, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		servingShards[i] = uint(s)
	}
	return disgord.ShardConfig{
		ShardCount: uint(shards),
		ShardIDs:   servingShards,
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
		BotToken:           bot.token,
		Logger:             logs.Logger,
		LoadMembersQuietly: true,
		ShardConfig:        loadShardingConfig(),
		Intents:            disgord.AllIntentsExcept(disgord.IntentGuildPresences, disgord.IntentGuildMembers, disgord.IntentGuildMessageTyping, disgord.IntentDirectMessageTyping),
	})

	defer client.Gateway().StayConnectedUntilInterrupted()
	// logFilter, _ := std.NewLogFilter(client) // Disabled to not log message content
	filter, _ := std.NewMsgFilter(router.RouterCtx, client)
	filter.SetPrefix(bot.Prefix)

	// create a handler and bind it to new message events
	client.Gateway().WithMiddleware(
		filter.NotByBot,  // ignore bot messages
		filter.HasPrefix, // message must have the given prefix
		// logFilter.LogMsg,   // log command message
		filter.StripPrefix, // remove the command prefix from the message
	).MessageCreate(router.CommandRouter(commands.Options))

	client.Gateway().BotReady(func() {
		// router.RegisterSlashCommands(client)

		client.UpdateStatusString(bot.status)
		logs.Info("Bot is ready!")

		go func() {
			time.Sleep(time.Second * 30)
			guilds := client.GetConnectedGuilds()
			logs.Info(fmt.Sprintf("Connected to %v guilds", len(guilds)))
		}()
	})
	// router.AddSlashCommandsHandlers(commands.Options)(client) // Registering slash commands does not work correctly
}
