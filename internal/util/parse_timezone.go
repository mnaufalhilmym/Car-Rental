package util

import (
	"time"

	"github.com/mnaufalhilmym/gotracing"
)

func ParseTimezone(timezone string) (time.Duration, error) {
	timezone = "UTC" + timezone

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		gotracing.Error("Invalid timezone", err)
		return 0, err
	}

	refTime := time.Now().UTC()

	_, offset := refTime.In(loc).Zone()

	return time.Duration(offset) * time.Second, nil
}
