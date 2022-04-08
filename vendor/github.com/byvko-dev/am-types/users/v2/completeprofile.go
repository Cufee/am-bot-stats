package users

import (
	"time"
)

type CompleteProfile struct {
	ID     string  `json:"id" firestore:"id" bson:"id"`
	Locale string  `json:"locale" firestore:"locale" bson:"locale"`
	Ban    UserBan `json:"ban" firestore:"ban" bson:"ban"`

	Features UserFeatures     `json:"features" firestore:"features" bson:"features"`
	Profiles ExternalProfiles `json:"profiles" firestore:"profiles" bson:"profiles"`

	LastSeen  time.Time `json:"last_seen" firestore:"last_seen" bson:"last_seen"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at" bson:"updated_at"`
}

func (u *CompleteProfile) Verify() bool {
	return u.ID != ""
}
