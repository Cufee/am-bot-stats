package router

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-core/logs"
	"github.com/byvko-dev/am-types/users/v2"
)

// Proxy types
type ArgumentDescription disgord.ApplicationCommandOption

type Command struct {
	Name      string                `json:"name" bson:"name"`
	Arguments []string              `json:"arguments" bson:"arguments"`
	User      users.CompleteProfile `json:"user" bson:"user"`

	userID       string                                      `json:"-" bson:"-"`
	options      RegisterOptions                             `json:"-" bson:"-"`
	session      disgord.Session                             `json:"-" bson:"-"`
	message      *disgord.Message                            `json:"-" bson:"-"`
	userCheck    func(string) (users.CompleteProfile, error) `json:"-" bson:"-"`
	errorPrinter func(string, error) string                  `json:"-" bson:"-"`
	reportError  func(ErrorContext, error)                   `json:"-" bson:"-"`

	Reply func(c Command, content ...interface{}) error `json:"-" bson:"-"`
}

type Permissions struct {
	AdminOnly  bool
	GuildOnly  bool
	DirectOnly bool
}

type RegisterOptions struct {
	AddSlashCommand  bool `json:"add_slash_command" bson:"add_slash_command"`
	AddPrefixCommand bool `json:"add_prefix_command" bson:"add_prefix_command"`

	Name        string                `json:"name" bson:"name"`
	Description string                `json:"description" bson:"description"`
	Aliases     []string              `json:"aliases" bson:"aliases"`
	Arguments   []ArgumentDescription `json:"arguments" bson:"arguments"`

	Permissions Permissions `json:"permissions" bson:"permissions"`

	Handler      func(Command) error  `json:"-" bson:"-"`
	ErrorHandler func(Command, error) `json:"-" bson:"-"`
}

type ErrorContext struct {
	Command   string   `json:"command" bson:"command"`
	Arguments []string `json:"arguments" bson:"arguments"`

	UserID    string `json:"user_id" bson:"user_id"`
	GuildID   string `json:"guild_id" bson:"guild_id"`
	ChannelID string `json:"channel_id" bson:"channel_id"`

	User users.CompleteProfile `json:"user" bson:"user"`
}

type RouterOptions struct {
	ErrorReportHandler func(ErrorContext, error)
	UserCheckHandler   func(string) (users.CompleteProfile, error)
	ErrorPrinter       func(string, error) string
}

func (o RouterOptions) check() {
	if o.ErrorPrinter == nil {
		panic(fmt.Errorf("CommandRouter: options.ErrorPrinter is nil"))
	}
	if o.ErrorReportHandler == nil {
		logs.Info("CommandRouter: options.ErrorReportHandler is nil")
	}
	if o.UserCheckHandler == nil {
		panic(fmt.Errorf("CommandRouter: options.UserCheckHandler is nil"))
	}
}
