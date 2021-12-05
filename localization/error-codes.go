package localization

import (
	"strings"
)

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

func newErrorWithCode(prefix ErrorPrefix, code, keyword, message string, report bool) ErrorWithCode {
	e := ErrorWithCode{ErrorCode(string(prefix) + code), report, keyword, message}
	validErrors = append(validErrors, e)
	return e
}

func parseError(err error) ErrorWithCode {
	if err == nil {
		return errorNil
	}
	for _, e := range validErrors {
		if strings.Contains(err.Error(), e.search) {
			return e
		}
	}
	return errorUnknown
}

func (e ErrorWithCode) Localized(locale string) string {
	return string(e.Code) + ": " + e.internalMessage
}

type errorReport struct {
	userMessage string
	err         error
}

func ReplyToErorr(err error, locale string, reply func(...interface{}) error, context ...interface{}) {
	if err == nil {
		return
	}

	e := parseError(err)
	var report = e.report
	var r errorReport
	r.err = err
	defer func(r errorReport) {
		if report {
			// TODO: report error
		}
	}(r)
	var message string = e.Localized(locale)
	r.userMessage = message

	newErr := reply(message)
	if newErr != nil {
		report = true
	}
}
