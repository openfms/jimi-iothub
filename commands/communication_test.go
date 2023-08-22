package commands

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCoreKitSwitch(t *testing.T) {
	expected := "COREKITSW,0"
	result := CoreKitSwitch(CoreKitIntegrated)
	assert.Equal(t, expected, result)
}

func TestModifyHttpUploadServer(t *testing.T) {
	expected := "UPLOAD,http://www.baidu.com/upload"
	result := ModifyHttpUploadServer("http://www.baidu.com/upload")
	assert.Equal(t, expected, result)
}

func TestSetHttpUploadLimit(t *testing.T) {
	expected := "HTTPUPLOADLIMIT,5,3"
	result := SetHttpUploadLimit(5, 3)
	assert.Equal(t, expected, result)
}

func TestSetRtmpServer(t *testing.T) {
	expected := "RSERVICE,192.168.0.1:1935/live"
	result := SetRtmpServer("192.168.0.1:1935/live")
	assert.Equal(t, expected, result)
}

func TestUploadEventFile(t *testing.T) {
	expected := "UPLOADFILE,EVENT_357730090564767_00000000_2021_01_29_07_28_18_F_05.mp4"
	result := UploadEventFile("EVENT_357730090564767_00000000_2021_01_29_07_28_18_F_05.mp4")
	assert.Equal(t, expected, result)
}

func TestCapturePicture(t *testing.T) {
	expected := "Picture,in"
	result := CapturePicture(CameraTypeInward)
	assert.Equal(t, expected, result)
}

func TestReplayVideoList(t *testing.T) {
	expected := "REPLAYLIST,video1.mp4,video2.mp4,video3.mp4"
	videoNames := []string{"video1.mp4", "video2.mp4", "video3.mp4"}
	result := ReplayVideoList(videoNames)
	assert.Equal(t, expected, result)
}

func TestRtmpLogin(t *testing.T) {
	expected := "RLOGIN,jimi,88888888"
	result := RtmpLogin("jimi", "88888888")
	assert.Equal(t, expected, result)
}

func TestUploadHistoryVideo(t *testing.T) {
	expected := "HVIDEO,2020_01_01_24_05_06,2"
	result := UploadHistoryVideo("2020_01_01_24_05_06", HistoryCameraTypeInward)
	assert.Equal(t, expected, result)
}

func TestSetFileListServer(t *testing.T) {
	expected := "FILELIST,http://www.baidu.com"
	result := SetFileListServer("http://www.baidu.com")
	assert.Equal(t, expected, result)
}

func TestSetTCPServer(t *testing.T) {
	expected := "SERVER,1,dvrdev.tracksolidpro.com,21100"
	result := SetTCPServer(HostTypeDomain, "dvrdev.tracksolidpro.com", 21100)
	assert.Equal(t, expected, result)
}

func TestCaptureVideo(t *testing.T) {
	expected := "Video,in,3s"
	result := CaptureVideo(CameraTypeInward, 3)
	assert.Equal(t, expected, result)
}

func TestRtmpLiveStreamWithPushDuration(t *testing.T) {
	expected := "RTMP,ON,IN,10"
	result := RtmpLiveStream(OnState, CameraTypeInward, 10)
	assert.Equal(t, expected, result)
}

func TestRtmpLiveStreamWithoutPushDuration(t *testing.T) {
	expected := "RTMP,ON,IN"
	result := RtmpLiveStream(OnState, CameraTypeInward, 0)
	assert.Equal(t, expected, result)
}

func TestUploadPlaybackVideosList(t *testing.T) {
	expected := "FILELIST"
	result := UploadPlaybackVideosList()
	assert.Equal(t, expected, result)
}

func TestUploadEventVideo(t *testing.T) {
	expected := "EVIDEO,2020-06-15 12:12:12,1,30"
	result := UploadEventVideo("2020-06-15 12:12:12", HistoryCameraTypeFront, 30)
	assert.Equal(t, expected, result)
}
