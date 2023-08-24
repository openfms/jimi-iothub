package client

import (
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_RealTimeAudioVideoTransmission(t *testing.T) {
	endPoint := os.Getenv("IOTHUB_ENDPOINT")
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	videoIPAddr := os.Getenv("IOTHUB_VIDEO_IP")
	iothubcli, err := NewIotHubClient(endPoint, "", "123456")
	assert.NilError(t, err)
	req := iothubcli.NewRealTimeAudioVideoRequest(deviceImei, &RealTimeCmdContent{
		DataType:       AudioVideoDataType,
		CodeStreamType: MainStream,
		VideoUDPPort:   "0",
		VideoIP:        videoIPAddr,
		VideoTCPPort:   "10002",
		Channel:        "1",
	})
	resp, err := iothubcli.RealTimeAudioVideoTransmission(req)
	assert.NilError(t, err)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}
