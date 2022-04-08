package commands

import (
	"fmt"
	"time"

	"aftermath.link/repo/am-bot-stats/config"
	"aftermath.link/repo/am-bot-stats/external/render"
	"aftermath.link/repo/am-bot-stats/external/users"
	"aftermath.link/repo/am-bot-stats/external/wargaming"
	"aftermath.link/repo/am-bot-stats/stats"

	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-core/discord/router"
	"github.com/byvko-dev/am-core/logs"
	types "github.com/byvko-dev/am-types/stats/v1"
)

func init() {
	logs.Info("Registering stats.go commands")
	router.RegisterCommand(statsOptions) // Register stats command on startup
}

var statsOptions = router.RegisterOptions{
	AddSlashCommand:  true,
	AddPrefixCommand: true,
	Name:             "session",
	Description:      "Check Session Statistics",
	Aliases:          []string{"wr", "we", "stats", "sess", "s"},
	Permissions:      defaultPremission,
	Arguments: []router.ArgumentDescription{{
		Name:        "name",
		Type:        disgord.OptionTypeString,
		Description: "Player in game nickname",
	}, {
		Name:        "server",
		Type:        disgord.OptionTypeString,
		Description: "NA, EU, RU, ASIA",
	}, {
		Name:        "days",
		Type:        disgord.OptionTypeString,
		Description: "The number of days to check",
	}},
	Handler:      StatsHandler,
	ErrorHandler: nil,
}

func StatsHandler(cmd router.Command) error {
	start := time.Now()

	var request types.StatsRequest
	request.Locale = cmd.User.Locale
	request.IncludeRating = true // TODO: Make this configurable

	if len(cmd.Arguments) == 0 && cmd.User.Profiles.Wargaming.PlayerID == 0 {
		// No default account set and no valid args provided
		return cmd.Reply(cmd, fmt.Sprintf("You never told me what your player name is. Please use `%v%v <player name>` or `@mention` another user.", config.DiscordCommandPrefix, cmd.Name))
	} else if len(cmd.Arguments) == 0 {
		request.PID = int(cmd.User.Profiles.Wargaming.PlayerID)
		request.Profile = "" // TODO - Add support for profiles
	} else {
		statsArgs := stats.ParseCommandArguments(cmd.Arguments)
		request.Days = statsArgs.Days
		if statsArgs.PlayerName != "" {
			// Get player ID
			pid, err := wargaming.IDFromName(statsArgs.PlayerName)
			if err != nil {
				return cmd.Reply(cmd, fmt.Sprintf("Error getting player ID: %v", err))
			}
			// Find user
			user, _ := users.CheckUserByPID(pid)
			request.Profile = ""
			request.PID = int(user.Profiles.Wargaming.PlayerID)
		} else if statsArgs.UserID != "" {
			// Get default player ID
			user, err := users.CheckUserByUserID(statsArgs.UserID)
			if err != nil {
				return cmd.Reply(cmd, fmt.Sprintf("Error getting user: %v", err))
			}
			request.Profile = ""
			request.PID = int(user.Profiles.Wargaming.PlayerID)
		}
	}

	if request.PID == 0 {
		return cmd.Reply(cmd, "No player ID found")
	}

	// Find realm
	realm, err := wargaming.RealmFromID(request.PID)
	if err != nil {
		// TODO: Handle error
		return cmd.Reply(cmd, fmt.Sprintf("Error getting realm: %v", err))
	}
	request.Realm = realm

	imageReader, err := render.GetPlayerStatsImage(request)
	if err != nil {
		return cmd.Reply(cmd, fmt.Sprintf("Error getting stats: %v", err))
	}
	cmd.Reply(cmd, fmt.Sprintf("Image generated in %v", time.Since(start).String()))

	var image disgord.CreateMessageFile
	image.FileName = "stats.jpg"
	image.Reader = imageReader

	cmd.Reply(cmd, "args: "+fmt.Sprintf("%+v", request), image)
	cmd.Reply(cmd, time.Since(start).String())
	return nil
}
