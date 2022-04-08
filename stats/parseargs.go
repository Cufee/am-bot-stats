package stats

import (
	"regexp"
	"strconv"
	"strings"
)

var mentionRegex = regexp.MustCompile(`<@!?(\d+)>`)

func ParseCommandArguments(args []string) statsArguments {
	var parsed statsArguments
	if len(args) == 0 {
		return parsed
	}

	for _, arg := range args {
		// Check mentions
		if mentionRegex.MatchString(arg) {
			parsed.UserID = mentionRegex.FindStringSubmatch(arg)[1]
		}
		// Check for days (number)
		if days, err := strconv.Atoi(arg); err == nil {
			parsed.Days = days
		}
	}

	// Check for player name and realm
	if len(args) > 0 && parsed.UserID == "" {
		parsed.PlayerName = args[0]
	}
	if strings.Contains(parsed.PlayerName, "-") {
		parsed.Realm = strings.ToUpper(strings.Split(parsed.PlayerName, "-")[1])
		parsed.PlayerName = strings.Split(parsed.PlayerName, "-")[0]
	}
	return parsed
}
