package utils

import (
	"fmt"

	"github.com/centrex/webcore/core/env"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "pgsql":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			env.GetEnv("DB_HOST", "127.0.0.1"),
			env.GetEnv("DB_PORT", "5432"),
			env.GetEnv("DB_USERNAME", "postgres"),
			env.GetEnv("DB_PASSWORD", "postgres"),
			env.GetEnv("DB_DATABASE", "webcore"),
			env.GetEnv("DB_SSL_MODE", "disable"),
		)
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			env.GetEnv("DB_USERNAME", "root"),
			env.GetEnv("DB_PASSWORD", ""),
			env.GetEnv("DB_HOST", "127.0.0.1"),
			env.GetEnv("DB_PORT", "3306"),
			env.GetEnv("DB_DATABASE", "webcore"),
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			env.GetEnv("REDIS_HOST", "127.0.0.1"),
			env.GetEnv("REDIS_PORT", "6379"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			env.GetEnv("SERVER_HOST", "127.0.0.1"),
			env.GetEnv("SERVER_PORT", "8080"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
