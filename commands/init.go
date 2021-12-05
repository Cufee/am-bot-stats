package commands

import commandutils "aftermath.link/repo/am-bot-stats/command-utils"

func Init() {
	// This is a dummy command to make sure this module is imported.
}

var defaultRunOptions = commandutils.RunOptions{AdminOnly: false, GuildOnly: false, DirectOnly: false}
var guildOnlyRunOptions = commandutils.RunOptions{AdminOnly: false, GuildOnly: true, DirectOnly: false}
var adminOnlyRunOptions = commandutils.RunOptions{AdminOnly: true, GuildOnly: false, DirectOnly: false}
