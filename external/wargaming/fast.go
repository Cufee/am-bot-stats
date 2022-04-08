package wargaming

import (
	"fmt"
	"net/url"
	"strconv"

	"aftermath.link/repo/am-bot-stats/config"
	"github.com/byvko-dev/am-core/helpers/requests"
	"github.com/byvko-dev/am-types/api/v1"
)

func RealmFromID(id int) (string, error) {
	apiUrl, err := url.Parse(fmt.Sprintf("%v/fast/account/id/%v/realm", config.WargamingApiUrl, id))
	if err != nil {
		return "", err
	}

	var response api.ResponseWithError
	_, err = requests.Send(apiUrl.String(), "GET", nil, nil, &response)
	if err != nil {
		return "", err
	}
	if response.Error.Message != "" {
		return "", fmt.Errorf("wargaming api error: %s", response.Error)
	}

	realm, ok := response.Data.(string)
	if !ok {
		return "", fmt.Errorf("wargaming api error: invalid response data")
	}

	return realm, nil
}

func IDFromName(name string) (int, error) {
	apiUrl, err := url.Parse(fmt.Sprintf("%v/fast/account/name/%v/id", config.WargamingApiUrl, name))
	if err != nil {
		return 0, err
	}

	var response api.ResponseWithError
	_, err = requests.Send(apiUrl.String(), "GET", nil, nil, &response)
	if err != nil {
		return 0, err
	}
	if response.Error.Message != "" {
		return 0, fmt.Errorf("wargaming api error: %s", response.Error)
	}

	id, err := strconv.Atoi(fmt.Sprint(response.Data))
	if err != nil {
		return 0, fmt.Errorf("wargaming api error: invalid response data")
	}

	return id, nil
}
