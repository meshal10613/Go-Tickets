package auth

import (
	"fmt"
	"go-tickets/internel/config"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(jwtDuration *string) (time.Duration, error) {
	var duration string

	if jwtDuration != nil {
		duration = *jwtDuration
	} else {
		cfg, err := config.LoadEnv()
		if err != nil {
			return 0, fmt.Errorf("failed to load environment variables: %w", err)
		}

		duration = cfg.JwtDuration
	}

	duration = strings.TrimSpace(strings.ToLower(duration))

	if strings.HasSuffix(duration, "d") {
		daysStr := strings.TrimSuffix(duration, "d")

		days, err := strconv.Atoi(daysStr)
		if err != nil {
			return 0, fmt.Errorf("invalid day duration: %s", duration)
		}

		return time.Duration(days) * 24 * time.Hour, nil
	}

	return time.ParseDuration(duration)
}
