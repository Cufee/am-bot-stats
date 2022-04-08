package users

import (
	"fmt"
	"net/url"

	"aftermath.link/repo/am-bot-stats/config"
	"aftermath.link/repo/am-bot-stats/external"
	legacy "github.com/byvko-dev/am-types/users/v1"
	"github.com/byvko-dev/am-types/users/v2"
)

// CheckUserByPID - Check user profile by player id
func CheckUserByPID(pid int) (users.CompleteProfile, error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/players/id/%v", config.UsersApiUrl, pid))
	if err != nil {
		return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = config.InternalApiKey

	// Send request
	var response legacy.User
	err = external.DecodeHTTPResponse("GET", headers, requestURL, nil, &response)
	if err != nil {
		return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
	}

	return userToProfile(response)
}

// CheckUserByUserID - Check user profile by Discord ID
func CheckUserByUserID(userIDStr string) (users.CompleteProfile, error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/users/id/%s", config.UsersApiUrl, userIDStr))
	if err != nil {
		return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = config.InternalApiKey

	// Send request
	var response legacy.User
	err = external.DecodeHTTPResponse("GET", headers, requestURL, nil, &response)
	if err != nil {
		return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
	}

	return userToProfile(response)
}

// CheckUserByName - Check user profile by player nickname
func CheckUserByName(name string) (users.CompleteProfile, error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/players/name/%s", config.UsersApiUrl, name))
	if err != nil {
		return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = config.InternalApiKey

	// Send request
	var response legacy.User
	err = external.DecodeHTTPResponse("GET", headers, requestURL, nil, &response)
	if err != nil {
		return users.CompleteProfile{}, fmt.Errorf("users api error: %s", err.Error())
	}

	return userToProfile(response)
}
