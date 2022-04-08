package users

import (
	"fmt"

	"aftermath.link/repo/am-bot-stats/external/wargaming"
	legacy "github.com/byvko-dev/am-types/users/v1"
	"github.com/byvko-dev/am-types/users/v2"
	accounts "github.com/byvko-dev/am-types/wargaming/v1/accounts"
)

func userToProfile(user legacy.User) (users.CompleteProfile, error) {
	profile := users.CompleteProfile{
		Locale: user.Locale,
		Ban: users.UserBan{
			Active:   user.Banned,
			Reason:   user.BanReason,
			Notified: user.BanNotified,
		},
		Profiles: users.ExternalProfiles{
			Wargaming: users.WargamingProfile{
				PlayerID: accounts.PlayerID(user.DefaultPID),
				Verified: user.Verified,
			},
		},
	}

	if user.ShadowBanned {
		profile.Features.Disabled = append(profile.Features.Disabled, users.CustomizeBackgrounds)
	}

	if profile.Profiles.Wargaming.PlayerID != 0 {
		realm, err := wargaming.RealmFromID(user.DefaultPID)
		if err != nil {
			return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
		}
		profile.Profiles.Wargaming.Realm = realm
	}
	return profile, nil
}
