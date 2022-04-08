package stats

import (
	types "github.com/byvko-dev/am-types/stats/v1"
)

func (s *statsArguments) BuildRequest() (types.StatsRequest, error) {
	var request types.StatsRequest
	request.Days = s.Days

	return types.StatsRequest{}, nil
}
