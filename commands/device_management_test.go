package commands

import (
	"fmt"
	"testing"
)

func TestParseParam(t *testing.T) {
	input := "IMEI:862798050982731,GPS Time:10,TIMEZONE:+04:30"

	parsedData := ParseParam(input)

	fmt.Printf("IMEI: %s\n", parsedData.IMEI)
	fmt.Printf("GPS Time: %s\n", parsedData.GPSTime)
	fmt.Printf("TIMEZONE: %s\n", parsedData.TimeZone)
}
