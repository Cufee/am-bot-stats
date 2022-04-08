package config

import "os"

func getEnv(required bool, keyAndDefault ...string) string {
	var (
		key          string
		defaultValue string
	)

	switch len(keyAndDefault) {
	case 1:
		key = keyAndDefault[0]
	case 2:
		key = keyAndDefault[0]
		defaultValue = keyAndDefault[1]
	default:
		panic("Invalid number of arguments")
	}

	value := os.Getenv(key)
	if value == "" {
		if required {
			panic("Required environment variable " + key + " not set")
		} else {
			return defaultValue
		}
	}
	return value
}
