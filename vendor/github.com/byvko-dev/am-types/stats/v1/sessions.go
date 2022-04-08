package stats

import (
	"time"

	"github.com/byvko-dev/am-types/wargaming/v1/statistics"
)

// Session - Will be switching to this format soon
type Session struct {
	Vehicles      []statistics.VehicleStatsFrame `json:"vehicles" bson:"vehicles"`
	Achievements  statistics.AchievementsFrame   `json:"achievements" bson:"achievements"`
	PlayerID      int                            `json:"player_id" bson:"player_id"`
	Timestamp     time.Time                      `json:"timestamp" bson:"timestamp"`
	LastBattle    time.Time                      `json:"last_battle_time" bson:"last_battle_time"`
	BattlesAll    int                            `json:"battles_random" bson:"battles_random"`
	StatsAll      statistics.StatsFrame          `json:"stats_random" bson:"stats_random"`
	BattlesRating int                            `json:"battles_rating" bson:"battles_rating"`
	StatsRating   statistics.StatsFrame          `json:"stats_rating" bson:"stats_rating"`
	SessionRating int                            `json:"session_wn8" bson:"session_wn8"`
}

// RetroSession - Session using old data structure
type RetroSession struct {
	Vehicles      map[string]statistics.VehicleStatsFrame `json:"vehicles" bson:"vehicles"`
	Achievements  statistics.AchievementsFrame            `json:"achievements" bson:"achievements"`
	PlayerID      int                                     `json:"player_id" bson:"player_id"`
	Timestamp     time.Time                               `json:"timestamp" bson:"timestamp"`
	LastBattle    time.Time                               `json:"last_battle_time" bson:"last_battle_time"`
	BattlesAll    int                                     `json:"battles_random" bson:"battles_random"`
	StatsAll      statistics.StatsFrame                   `json:"stats_random" bson:"stats_random"`
	BattlesRating int                                     `json:"battles_rating" bson:"battles_rating"`
	StatsRating   statistics.StatsFrame                   `json:"stats_rating" bson:"stats_rating"`
	SessionRating int                                     `json:"session_wn8" bson:"session_wn8"`
}
