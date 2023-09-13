package commands

import (
	"fmt"
	"regexp"
)

const DefaultAdminPassword = "666666"

// ChangePassword changes the password of device
// Example: PASSWORD,666666,123456
func ChangePassword(oldPassword, newPassword string) string {
	return fmt.Sprintf("%s,%s,%s", PASSWORD, oldPassword, newPassword)
}

type VolumeLevel uint8

const (
	VolumeMute VolumeLevel = 0
	VolumeLow  VolumeLevel = 1
	VolumeMid  VolumeLevel = 2
	VolumeHigh VolumeLevel = 3
)

// SetDeviceVolume defines the volume level of the device.
func SetDeviceVolume(level VolumeLevel) string {
	return fmt.Sprintf("%s,%d", VOLUME, level)
}

// SetLEDOnOff defines whether to light on LEDs (all-day)
func SetLEDOnOff(state OnOffState) string {
	return fmt.Sprintf("%s,%s", LED, state)
}

type DeviceParams struct {
	IMEI     string
	TimeZone string
	GPSTime  string
}

func ParseParam(param string) *DeviceParams {
	// Define regular expressions to match IMEI, GPS Time, and TIMEZONE values
	imeiRegex := regexp.MustCompile(`IMEI:(\d+)`)
	gpsTimeRegex := regexp.MustCompile(`GPS Time:(\d+)`)
	timezoneRegex := regexp.MustCompile(`TIMEZONE:([+\-]\d+:\d+)`)

	// Find matches in the input string
	imeiMatch := imeiRegex.FindStringSubmatch(param)
	gpsTimeMatch := gpsTimeRegex.FindStringSubmatch(param)
	timezoneMatch := timezoneRegex.FindStringSubmatch(param)

	// Extract values from the matches
	result := &DeviceParams{}
	if len(imeiMatch) == 2 {
		result.IMEI = imeiMatch[1]
	}
	if len(gpsTimeMatch) == 2 {
		result.GPSTime = gpsTimeMatch[1]
	}
	if len(timezoneMatch) == 2 {
		result.TimeZone = timezoneMatch[1]
	}
	return result
}
