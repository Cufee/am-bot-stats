package stats

type statsArguments struct {
	UserID     string `json:"userId"`
	PlayerName string `json:"playerName"`
	Realm      string `json:"realm"`
	Days       int    `json:"days"`
}
