package env

import (
	"os"
	"strconv"
)

// Configuration properties
type Configuration struct {
	Port          int    `json:"port"`
	MysqlURL      string `json:"mysqlUrl"`
	RedisURL      string `json:"redisUrl"`
	SessionSecret string `json:"sessionSecret"`
}

var config *Configuration

func new() *Configuration { //default configuration
	return &Configuration{
		Port:          3010,
		MysqlURL:      "root:jose1@/wallet",
		RedisURL:      "host.docker.internal:49153",
		SessionSecret: "ThisIsASecretShh",
	}
}

// Get: Gets environment variables
func Get() *Configuration {
	config = load()
	return config
}

// Load file properties
func load() *Configuration {
	result := new()
	if value := os.Getenv("SESSION_SECRET"); len(value) > 0 {
		result.SessionSecret = value
	}
	if value := os.Getenv("REDIS_URL"); len(value) > 0 {
		result.RedisURL = value
	}
	if value := os.Getenv("MYSQL_URL"); len(value) > 0 {
		result.MysqlURL = value
	}

	if value := os.Getenv("PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err != nil {
			result.Port = intVal
		}
	}

	return result
}
