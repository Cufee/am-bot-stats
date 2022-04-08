package disgord

// Code generated - This file has been automatically generated by generate/intents/main.go - DO NOT EDIT.
// Warning: This file is overwritten at "go generate", instead adapt internal/gateway/intents.go and run go generate

import (
	"github.com/andersfylling/disgord/internal/gateway"
)

type Intent = gateway.Intent

const (
	IntentDirectMessageReactions = gateway.IntentDirectMessageReactions
	IntentDirectMessageTyping    = gateway.IntentDirectMessageTyping
	IntentDirectMessages         = gateway.IntentDirectMessages
	IntentGuildBans              = gateway.IntentGuildBans
	IntentGuildEmojisAndStickers = gateway.IntentGuildEmojisAndStickers
	IntentGuildIntegrations      = gateway.IntentGuildIntegrations
	IntentGuildInvites           = gateway.IntentGuildInvites
	IntentGuildMembers           = gateway.IntentGuildMembers
	IntentGuildMessageReactions  = gateway.IntentGuildMessageReactions
	IntentGuildMessageTyping     = gateway.IntentGuildMessageTyping
	IntentGuildMessages          = gateway.IntentGuildMessages
	IntentGuildPresences         = gateway.IntentGuildPresences
	IntentGuildScheduledEvents   = gateway.IntentGuildScheduledEvents
	IntentGuildVoiceStates       = gateway.IntentGuildVoiceStates
	IntentGuildWebhooks          = gateway.IntentGuildWebhooks
	IntentGuilds                 = gateway.IntentGuilds
)

func AllIntents() Intent {
	return AllIntentsExcept()
}

func AllIntentsExcept(exceptions ...Intent) Intent {
	IntentsMap := map[Intent]int8{
		IntentDirectMessageReactions: 0,
		IntentDirectMessageTyping:    0,
		IntentDirectMessages:         0,
		IntentGuildBans:              0,
		IntentGuildEmojisAndStickers: 0,
		IntentGuildIntegrations:      0,
		IntentGuildInvites:           0,
		IntentGuildMembers:           0,
		IntentGuildMessageReactions:  0,
		IntentGuildMessageTyping:     0,
		IntentGuildMessages:          0,
		IntentGuildPresences:         0,
		IntentGuildScheduledEvents:   0,
		IntentGuildVoiceStates:       0,
		IntentGuildWebhooks:          0,
		IntentGuilds:                 0,
	}

	for i := range exceptions {
		delete(IntentsMap, exceptions[i])
	}

	var intents gateway.Intent
	for intent := range IntentsMap {
		intents |= intent
	}
	return intents
}