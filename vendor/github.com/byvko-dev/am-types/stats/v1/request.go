package stats

// StatsRequest - Request struct for stats api
type StatsRequest struct {
	PID           int    `bson:"player_id" json:"player_id"`
	Realm         string `bson:"realm" json:"realm"`
	Days          int    `bson:"days" json:"days"`
	Locale        string `json:"locale" bson:"locale"`
	Profile       string `json:"profile" bson:"profile"`
	SortKey       string `bson:"sort_key" json:"sort_key"`
	TankLimit     int    `bson:"detailed_limit" json:"detailed_limit"`
	IncludeRating bool   `bson:"include_rating" json:"include_rating"`
}
