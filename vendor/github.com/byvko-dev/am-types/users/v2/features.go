package users

import (
	"github.com/byvko-dev/am-core/helpers/slices"
)

type Feature string
type Features []Feature

type UserFeatures struct {
	Enabled  Features `json:"enabled" bson:"enabled"`
	Disabled Features `json:"disabled" bson:"disabled"`
}

const PreviewAccess Feature = "preview_access"
const EnhancedSecurity Feature = "enhanced_security"
const DiscordAdminCommands Feature = ("discord_admin_commands")

const PremiumBypass Feature = "premium_bypass"
const CustomizeProfile Feature = "customize_profile"
const CustomizeBackgrounds Feature = "customize_backgrounds"

func (all *Features) Includes(f Feature) bool {
	return slices.Contains(all, f) > -1
}
