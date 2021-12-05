package commands

import (
	commandutils "aftermath.link/repo/am-bot-stats/command-utils"
	"aftermath.link/repo/am-bot-stats/logs"
)

func init() {
	logs.Info("Registering ping command")
	commandutils.RegisterCommand("ping", defaultRunOptions, PingCommandHandler, nil, "p") // Register ping command on startup
}

func PingCommandHandler(cmd commandutils.Command) error {
	return cmd.Reply("Pong!")
}
