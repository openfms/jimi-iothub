package commands

import "fmt"

/*
UploadEventVideo Generate and upload event videos which store in SD card.

timestamp Format=Year-Month-Day Hour:Minute:Second.
length between 10-60 second default is 15.

Example: EVIDEO,2020-06-15 12:12:12,1,30
*/
func RecordAudio(timeStamp string, cameraType HistoryCameraType, lengthSecond uint8) string {
	return fmt.Sprintf("%s,%s,%d,%d", EVIDEO, timeStamp, cameraType, lengthSecond)
}
