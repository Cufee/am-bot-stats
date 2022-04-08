package errors

import (
	"strings"

	"github.com/byvko-dev/am-core/logs"
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

func ReplyToError(err error, locale string, reply func(...interface{}) error, context ...interface{}) {
	logs.Error("%v", err)
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
