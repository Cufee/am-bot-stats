package users

import (
	"github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

type ExternalProfiles struct {
	Discord   DiscordProfile   `json:"discord" firestore:"discord" bson:"discord"`
	Wargaming WargamingProfile `json:"wargaming" firestore:"wargaming" bson:"wargaming"`
}

type DiscordProfile struct {
	Verified bool   `json:"verified" firestore:"verified" bson:"verified"`
	UserID   string `json:"user_id" firestore:"user_id" bson:"user_id"`
}

type WargamingProfile struct {
	Verified bool              `json:"verified" firestore:"verified" bson:"verified"`
	PlayerID accounts.PlayerID `json:"player_id" firestore:"player_id" bson:"player_id"`
	Realm    string            `json:"realm" firestore:"realm" bson:"realm"`
}
