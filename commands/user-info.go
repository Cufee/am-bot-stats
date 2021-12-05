package commands

import (
	commandutils "aftermath.link/repo/am-bot-stats/command-utils"
	"aftermath.link/repo/am-bot-stats/logs"
)

func init() {
	logs.Info("Registering ping command")
	commandutils.RegisterCommand("info", defaultRunOptions, UserInfoHandler, nil) // Register ping command on startup
}

func UserInfoHandler(cmd commandutils.Command) error {
	cmd.Reply("User info command")
	cmd.Reply(cmd.User.DefaultPID)
	return nil
}
