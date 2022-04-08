package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"aftermath.link/repo/am-bot-stats/config"
	"github.com/byvko-dev/am-types/api/v1"
	"github.com/byvko-dev/am-types/stats/v1"
)

// GetPlayerStatsImage - Get stats image for a player
func GetPlayerStatsImage(reqData stats.StatsRequest) (imgReader *bytes.Reader, err error) {
	// Marshal request
	byteData, err := json.Marshal(reqData)
	if err != nil {
		return imgReader, fmt.Errorf("render api error: %s", err.Error())
	}

	// Make URL
	requestURL, err := url.Parse((config.RenderApiUrl + "/stats/options"))
	if err != nil {
		return imgReader, fmt.Errorf("render api error: %s", err.Error())
	}

	return getImage(requestURL, byteData)
}

func getImage(url *url.URL, payload []byte) (imgReader *bytes.Reader, err error) {
	// Send request
	res, err := http.DefaultClient.Post(url.String(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return imgReader, fmt.Errorf("render api error: %s", err.Error())
	}

	defer res.Body.Close()

	if res.Header.Get("Content-Type") != "image/png" {
		var data api.ResponseWithError
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			return imgReader, fmt.Errorf("render api error: %s", err.Error())
		}
		return imgReader, fmt.Errorf("render api error: %v", data.Error.Message)
	}

	img, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return imgReader, fmt.Errorf("render api error: %s", err.Error())
	}

	return bytes.NewReader([]byte(img)), err
}
