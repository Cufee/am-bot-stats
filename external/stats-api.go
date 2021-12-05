package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type statsApiResponse struct {
	Image       []byte `json:"image"`
	UniqueCards int    `json:"unique_cards"`
	Analytics   struct {
		TimeToComplete float64 `json:"request_time_sec"`
	} `json:"analytics"`
	Error string `json:"error"`
}

// // GetPlayerStatsImage - Get stats image for a player
// func GetPlayerStatsImage(reqData mongodbapi.StatsRequest) (response apiResponse, imgReader *bytes.Reader, err error) {
// 	// Marshal request
// 	byteData, err := json.Marshal(reqData)
// 	if err != nil {
// 		return response, imgReader, fmt.Errorf("stats api error: %s", err.Error())
// 	}

// 	// Make URL
// 	requestURL, err := url.Parse((config.StatsAPIURL + "/stats/image"))
// 	if err != nil {
// 		return response, imgReader, fmt.Errorf("stats api error: %s", err.Error())
// 	}
// 	log.Print(requestURL)

// 	return getImageLegacy(requestURL, byteData)
// }

func getImage(url *url.URL, payload []byte) (responseData statsApiResponse, imgReader *bytes.Reader, err error) {
	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = ""
	headers["Content-Type"] = "application/json"

	// Send request
	rawRes, err := RawHTTPResponse("POST", headers, url, payload)
	if err != nil {
		return responseData, imgReader, fmt.Errorf("stats api error: %s", err.Error())
	}

	// Unmarshal error
	err = json.Unmarshal(rawRes, &responseData)
	if err != nil {
		return responseData, imgReader, err
	}

	// Check for returned error
	if responseData.Error != "" {
		return responseData, imgReader, fmt.Errorf("stats api error: %s", responseData.Error)
	}

	// Decode image
	imgReader = bytes.NewReader(responseData.Image)
	return responseData, imgReader, err
}
