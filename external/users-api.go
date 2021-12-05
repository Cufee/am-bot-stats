package external

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"
)

// UserProfile - reponse from a user check
type UserProfile struct {
	DefaultPID int    `json:"player_id"`
	Locale     string `json:"locale"`

	Premium  bool `json:"premium"`
	Verified bool `json:"verified"`

	CustomBgURL string `json:"bg_url"`
	SelectedUI  string `json:"selected_ui"`

	Banned      bool   `json:"banned"`
	BanReason   string `json:"ban_reason,omitempty"`
	BanNotified bool   `json:"ban_notified,omitempty"`

	Error string `json:"error"`

	ShadowBanned    bool   `json:"shadow_banned"`
	ShadowBanReason string `json:"shadow_ban_reason"`
}

// BanData - request for ban route
type BanData struct {
	Reason     string    `json:"reason"`
	Notified   bool      `json:"notified"`
	Expiration time.Time `json:"expiration"`
}

var usersApiURI string

func init() {
	usersApiURI = os.Getenv("USERS_API_URI")
	if usersApiURI == "" {
		panic("USERS_API_URI is not set")
	}
}

// CheckUserByPID - Check user profile by player id
func CheckUserByPID(pid int) (userData UserProfile, err error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/players/id/%v", usersApiURI, pid))
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = applicationApiKey

	// Send request
	err = DecodeHTTPResponse("GET", headers, requestURL, nil, &userData)
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Check error
	if userData.Error != "" {
		err = fmt.Errorf("users api error: %s", userData.Error)
	}
	return userData, err
}

// BanUserByID - Ban user
func BanUserByID(userID string, reason string, notified bool, expiration time.Time) error {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/users/id/%v/ban", usersApiURI, userID))
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Prepare ban data
	var banData BanData
	banData.Reason = reason
	banData.Notified = notified
	banData.Expiration = expiration

	// Marshal ban data
	banDataRaw, err := json.Marshal(banData)
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = applicationApiKey
	headers["Content-Type"] = "application/json"

	// Send request
	var errorData UserProfile
	err = DecodeHTTPResponse("POST", headers, requestURL, banDataRaw, &errorData)
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Check for returned error
	if errorData.Error != "" {
		err = fmt.Errorf("users api error: %s", errorData.Error)
	}
	return err
}

// CheckUserByUserID - Check user profile by Discord ID
func CheckUserByUserID(userIDStr string) (userData UserProfile, err error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/users/id/%s", usersApiURI, userIDStr))
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = applicationApiKey

	// Send request
	err = DecodeHTTPResponse("GET", headers, requestURL, nil, &userData)
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Check for returned error
	if userData.Error != "" {
		err = fmt.Errorf("users api error: %s", userData.Error)
	}
	return userData, err
}

// CheckUserByName - Check user profile by player nickname
func CheckUserByName(name string) (userData UserProfile, err error) {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/players/name/%s", usersApiURI, name))
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = applicationApiKey

	// Send request
	err = DecodeHTTPResponse("GET", headers, requestURL, nil, &userData)
	if err != nil {
		return userData, fmt.Errorf("users api error: %s", err.Error())
	}

	// Check for returned error
	if userData.Error != "" {
		err = fmt.Errorf("users api error: %s", userData.Error)
	}
	return userData, err
}

// SetNewDefaultPID - Set new default account id for user ID
func SetNewDefaultPID(pid int, userID string) error {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/users/id/%s/newdef/%v", usersApiURI, userID, pid))
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = applicationApiKey

	// Send request
	var errorData statsApiResponse
	err = DecodeHTTPResponse("PATCH", headers, requestURL, nil, &errorData)
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Check for returned error
	if errorData.Error != "" {
		return fmt.Errorf("users api error: %s", errorData.Error)
	}
	return err
}

// SetNewDefaultBG - Set new default background for user ID
func SetNewDefaultBG(userID string, imgURL string) error {
	// Make URL
	requestURL, err := url.Parse(fmt.Sprintf("%s/background/%v?bgurl=%s", usersApiURI, userID, imgURL))
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Make headers
	headers := make(map[string]string)
	headers["x-api-key"] = applicationApiKey

	// Send request
	var errorData statsApiResponse
	err = DecodeHTTPResponse("PATCH", headers, requestURL, nil, &errorData)
	if err != nil {
		return fmt.Errorf("users api error: %s", err.Error())
	}

	// Check for returned error
	if errorData.Error != "" {
		return fmt.Errorf("users api error: %s", errorData.Error)
	}
	return err
}
