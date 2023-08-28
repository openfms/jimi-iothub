package commands

import "fmt"

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
