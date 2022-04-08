package config

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DiscordToken         string
	DiscordBotName       string
	DiscordBotStatus     string
	DiscordCommandPrefix string
)

func init() {
	DiscordToken = getEnv(true, "DISCORD_TOKEN")
	DiscordCommandPrefix = getEnv(true, "DISCORD_COMMAND_PREFIX")

	DiscordBotStatus = fmt.Sprintf(getEnv(false, "DISCORD_BOT_STATUS", "%[1]vhelp"), DiscordCommandPrefix)
}
