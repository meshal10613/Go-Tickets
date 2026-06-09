package auth

import (
	"fmt"
	"go-tickets/internel/config"
	"strconv"
	"strings"
	"time"
)

func ParseDuration() (time.Duration, error) {
	cfg, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Sprintf("failed to load environment variables: %v", err))
	}
	duration := strings.TrimSpace(strings.ToLower(cfg.JwtDuration))

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
