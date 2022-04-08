package config

import _ "github.com/joho/godotenv/autoload"

var (
	InternalApiKey string

	RenderApiUrl    string
	UsersApiUrl     string
	WargamingApiUrl string
)

func init() {
	WargamingApiUrl = getEnv(true, "WARGAMING_API_URL")
	InternalApiKey = getEnv(true, "INTERNAL_API_KEY")
	RenderApiUrl = getEnv(true, "RENDER_API_URL")
	UsersApiUrl = getEnv(true, "USERS_API_URL")
}
