package errors

type ErrorCode string
type ErrorPrefix string
type ErrorWithCode struct {
	Code            ErrorCode
	report          bool
	search          string
	internalMessage string
}

var validErrors []ErrorWithCode

var (
	// No keyword error
	errorKeywordNone = "!@#$ - not a valid keyword"

	// Generic errors - prefix - "0"
	errorNil     = newErrorWithCode("0", "00", errorKeywordNone, "Invalid error", false) // Invalid / nil error
	errorUnknown = newErrorWithCode("0", "01", errorKeywordNone, "Unknown Error", true)  // Unknown error

	// prefix for general, bot related errors - "1"
	generalErrorPrefix            ErrorPrefix = "1"
	ErrUnknownCommandKeyword                  = "unknown command"
	ErrUnknownCommand                         = newErrorWithCode(generalErrorPrefix, "001", ErrUnknownCommandKeyword, "Unknown Command", false)
	ErrCommandIsGuildOnlyKeyword              = "guild only command ran in direct messages"
	ErrCommandIsGuildOnly                     = newErrorWithCode(generalErrorPrefix, "002", ErrCommandIsGuildOnlyKeyword, "This command cannot be executed in Direct Messages", false)
	ErrCommandIsDirectOnlyKeyword             = "dm only command ran in guild"
	ErrCommandIsDirectOnly                    = newErrorWithCode(generalErrorPrefix, "003", ErrCommandIsGuildOnlyKeyword, "This command cannot be executed in a Server", false)

	// prefix for discord API related errors - "2"
	discordErrorPrefix ErrorPrefix = "2"
	ErrDiscordError                = newErrorWithCode(discordErrorPrefix, "001", "discord", "Discord Error", false)

	// prefix for users api related errors - "3"
	usersApiErrorPrefix  ErrorPrefix = "3"
	ErrUserBannedKeyword             = "user banned from Aftermath"
	ErrUserBanned                    = newErrorWithCode(usersApiErrorPrefix, "001", ErrUserBannedKeyword, "User is banned", false)
)
