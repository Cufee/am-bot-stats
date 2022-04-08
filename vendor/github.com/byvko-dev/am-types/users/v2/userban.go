package users

import "github.com/byvko-dev/am-types/users/generic"

type UserBan struct {
	Active     bool              `json:"active" bson:"active"`
	Reason     string            `json:"reason" bson:"reason"`
	Notified   bool              `json:"notified" bson:"notified"`
	Expiration generic.Timestamp `json:"expiration" bson:"expiration"`
}
