package commands

import (
	api "aftermath.link/repo/am-bot-stats/external/users"
	"aftermath.link/repo/am-bot-stats/logs"
	"github.com/byvko-dev/am-core/discord/router"
	"github.com/byvko-dev/am-types/users/v2"
)

func Init() {
	// This is a dummy command to make sure this package is imported.
}

var defaultPremission = router.Permissions{AdminOnly: false, GuildOnly: false, DirectOnly: false}
var adminOnlyPermissions = router.Permissions{AdminOnly: true, GuildOnly: false, DirectOnly: false}
var guildOnlyPermissions = router.Permissions{AdminOnly: false, GuildOnly: true, DirectOnly: false}

func errorPrinter(locale string, err error) string {
	if err == nil {
		return "invalid error"
	}
	return err.Error()
}

func reportError(ctx router.ErrorContext, err error) {
	logs.Error("Error in command %v for %v: ", ctx.Command, ctx.UserID, err)
	// message := err.Error()
	// // Fix blank error message
	// if message == "" {
	// 	message = "Unknown error"
	// }
	// // Mesage on the dev server
	// reportMsg := fmt.Sprintf("**Error Report:\n**Guild ID: %v\nChannel ID: %v\nUser ID: %v\nMessaged: %v\nError:```%v```User Message:```%v```", guildID, channelID, userID, messagedChan, errMsg, userMsgRaw)
	// s.ChannelMessageSend(config.ErrorsChanID, reportMsg)
}

func getUserInfo(id string) (users.CompleteProfile, error) {
	user, err := api.CheckUserByUserID(id)
	if err != nil {
		logs.Error("users api error: ", err)
	}
	return user, nil
}

var Options = router.RouterOptions{
	ErrorReportHandler: reportError,
	ErrorPrinter:       errorPrinter,
	UserCheckHandler:   getUserInfo,
}
