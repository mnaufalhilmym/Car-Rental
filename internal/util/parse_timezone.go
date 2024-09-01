package util

import (
	"time"

	"github.com/mnaufalhilmym/gotracing"
)

func ParseTimezone(timezone string) (*time.Location, error) {
	timezone = "UTC" + timezone

	location, err := time.LoadLocation(timezone)
	if err != nil {
		gotracing.Error("Invalid timezone", err)
		return nil, err
	}

	return location, nil
}
