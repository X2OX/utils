package timezone

import (
	"os"
	"strconv"
	"time"
)

const (
	FixedZoneName     = "X2OX_TIMEZONE_FIXED_ZONE_NAME"
	FixedZoneOffset   = "X2OX_TIMEZONE_FIXED_ZONE_OFFSET"
	Name              = "X2OX_TIMEZONE_NAME"
	NameUnpredictable = "X2OX_TIMEZONE_NAME_UNPREDICTABLE"
)

func init() {
	switch {
	case setFixedZone():
	case setTimeZone():
	case setTimeZoneEnv():
	}
}

func setFixedZone() bool {
	name := os.Getenv(FixedZoneName)
	if name == "" {
		return false
	}
	sOffset := os.Getenv(FixedZoneOffset)
	if sOffset == "" {
		return false
	}
	offset, err := strconv.Atoi(sOffset)
	if err != nil {
		return false
	}
	time.Local = time.FixedZone(name, offset)
	return true
}
func setTimeZone() bool {
	name := os.Getenv(Name)
	if name == "" {
		return false
	}
	loc, err := time.LoadLocation(name)
	if err == nil {
		time.Local = loc
	}
	return err == nil
}

func setTimeZoneEnv() bool {
	name := os.Getenv(NameUnpredictable)
	if name == "" {
		return false
	}
	return os.Setenv("TZ", name) == nil
}
