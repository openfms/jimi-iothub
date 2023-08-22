package commands

import (
	"fmt"
	"strings"
)

type CoreKitMode int

const (
	CoreKitIntegrated  CoreKitMode = 0
	CoreKitDistributed CoreKitMode = 1
)

/*
	CoreKitSwitch switches platform mode

Note: Before switching the device to the integrated version, you must first do the followings in the strict order:
switch the working logic, modify the addresses of HTTP and RTMP servers, and modify the address of the TCP server.
*/
func CoreKitSwitch(mode CoreKitMode) string {
	return fmt.Sprintf("%s,%d", COREKITSW, mode)
}

// ModifyHttpUploadServer modifies http upload server url
// Example: UPLOAD,http://www.baidu.com/upload
func ModifyHttpUploadServer(url string) string {
	return fmt.Sprintf("%s,%s", UPLOAD, url)
}

/*
	SetHttpUploadLimit changes http upload limit for when platform doesn't respond after the device uploads data over HTTP

retryCount[1-10] default: 5
retryInterval[1-30] default: 3minute

Example: HTTPUPLOADLIMIT,5,3
*/
func SetHttpUploadLimit(retryCount, retryInterval uint8) string {
	if retryCount < 1 || retryCount > 10 {
		retryCount = 5
	}
	if retryInterval < 1 || retryCount > 30 {
		retryInterval = 3
	}
	return fmt.Sprintf("%s,%d,%d", HTTPUPLOADLIMIT, retryCount, retryInterval)
}

// SetRtmpServer modifies the address of the RTMP server
// Example: RSERVICE,192.168.0.1:1935/live
func SetRtmpServer(url string) string {
	return fmt.Sprintf("%s,%s", RSERVICE, url)
}

/*
SetFileListServer modifies the url address to receive the history videos name list

Example: FILELIST,http://www.baidu.com
*/
func SetFileListServer(url string) string {
	return fmt.Sprintf("%s,%s", FILELIST, url)
}

type HostType int

const (
	HostTypeIPAddress HostType = 0
	HostTypeDomain    HostType = 1
)

// SetTCPServer modifies the address of the tcp server.
// The device will restart after the address of the TCP server is changed.
// Example: SERVER,1,dvrdev.tracksolidpro.com,21100
func SetTCPServer(host HostType, serverAddr string, port uint16) string {
	return fmt.Sprintf("%s,%d,%s,%d", SERVER, host, serverAddr, port)
}

type CameraType string

const (
	CameraTypeInward CameraType = "IN"
	CameraTypeFront  CameraType = "OUT"
	CameraTypeBoth   CameraType = "INOUT"
)

// CapturePicture Capture the images from the device.
// Example: Picture,in
func CapturePicture(camera CameraType) string {
	return fmt.Sprintf("%s,%s", Picture, strings.ToLower(string(camera)))
}

// CaptureVideo Capture the video from the device.
// duration should be between 3s and 10s.
// Example: Video,in,3s
func CaptureVideo(camera CameraType, seconds uint8) string {
	return fmt.Sprintf("%s,%s,%ds", Video, strings.ToLower(string(camera)), seconds)
}

type OnOffState string

const (
	OffState OnOffState = "OFF"
	OnState  OnOffState = "ON"
)

// RtmpLiveStream enables,disables live video stream
// duration should be between 2 and 180 minute default is 15.
// Example: Video,in,3s
func RtmpLiveStream(OnOff OnOffState, camera CameraType, pushDuration uint8) string {
	if pushDuration > 0 {
		return fmt.Sprintf("%s,%s,%s,%d", RTMP, OnOff, camera, pushDuration)
	}
	return fmt.Sprintf("%s,%s,%s", RTMP, OnOff, camera)
}

// UploadPlaybackVideosList requests to upload the playback video name list
func UploadPlaybackVideosList() string {
	return fmt.Sprintf("%s", FILELIST)
}

// ReplayVideoList Request playback video streaming
// Example: REPLAYLIST,2021_05_31_08_10_45_02.mp4,2021_05_31_08_11_46_02.mp4,2021_05_31_08_12_48_02.mp4
func ReplayVideoList(videoNames []string) string {
	return fmt.Sprintf("%s,%s", REPLAYLIST, strings.Join(videoNames, ","))
}

// RtmpLogin sets login information for rtmp live video streaming
// Example: RLOGIN,jimi,88888888
func RtmpLogin(userName, password string) string {
	return fmt.Sprintf("%s,%s,%s", RLOGIN, userName, password)
}

type HistoryCameraType int

const (
	HistoryCameraTypeInward HistoryCameraType = 2
	HistoryCameraTypeFront  HistoryCameraType = 1
)

// UploadHistoryVideo Upload the playback video which store in memory (each one minute)
// timestamp format: (Year_Month_Day_Hour_Minute_Second)
// Example: HVIDEO,2020_01_01_24_05_06,1
func UploadHistoryVideo(timeStamp string, cameraType HistoryCameraType) string {
	return fmt.Sprintf("%s,%s,%d", HVIDEO, timeStamp, cameraType)
}

/*
UploadEventVideo Generate and upload event videos which store in SD card.

timestamp Format=Year-Month-Day Hour:Minute:Second.
length between 10-60 second default is 15.

Example: EVIDEO,2020-06-15 12:12:12,1,30
*/
func UploadEventVideo(timeStamp string, cameraType HistoryCameraType, lengthSecond uint8) string {
	return fmt.Sprintf("%s,%s,%d,%d", EVIDEO, timeStamp, cameraType, lengthSecond)
}

// UploadEventFile Upload videos of a specific event type (a command to upload video files on demand)
// Example: UPLOADFILE,EVENT_357730090564767_00000000_2021_01_29_07_28_18_F_05.mp4
func UploadEventFile(fileName string) string {
	return fmt.Sprintf("%s,%s", UPLOADFILE, fileName)
}
