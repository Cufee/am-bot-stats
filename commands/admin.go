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
	Permissions:      adminOnlyPermissions,
	Handler:          CountGuildsHandler,
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
