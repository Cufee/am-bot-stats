package commands

import (
	"aftermath.link/repo/am-bot-stats/logs"
	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-core/discord/router"
)

func init() {
	logs.Info("Registering admin.go commands")
	router.RegisterCommand(countGuildsOptions) // Register ping command on startup
}

var countGuildsOptions = router.RegisterOptions{
	AddSlashCommand:  false,
	AddPrefixCommand: true,
	Name:             "guilds",
	Description:      "Counts the number of guilds the bot is in",
	Aliases:          nil,
	Permissions:      adminOnlyPermissions,
	Arguments: []router.ArgumentDescription{{
		Name:        "Name",
		Type:        disgord.OptionTypeString,
		Description: "Player name to check",
	}},
	Handler:      CountGuildsHandler,
	ErrorHandler: nil,
}

func CountGuildsHandler(cmd router.Command) error {
	session := cmd.UnsafeGetSession()
	params := disgord.GetCurrentUserGuilds{}
	guilds, err := session.CurrentUser().GetGuilds(&params)
	if err != nil {
		return err
	}
	return cmd.Reply(cmd, "I'm in ", len(guilds), " guilds")
}
