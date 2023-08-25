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
	req := iothubcli.RealTimeAVRequest(deviceImei, &RealTimeCmdContent{
		DataType:       AudioVideoDataType,
		CodeStreamType: MainStream,
		VideoUDPPort:   "0",
		VideoIP:        videoIPAddr,
		VideoTCPPort:   "10002",
		Channel:        "1",
	})
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}

func TestIotHubClient_RealTimeAVControlRequest(t *testing.T) {
	endPoint := os.Getenv("IOTHUB_ENDPOINT")
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(endPoint, "", "123456")
	assert.NilError(t, err)
	req := iothubcli.RealTimeAVControlRequest(deviceImei, &RealTimeControlCmdContent{
		DataType:       TurnOffBothAudioAndVideo,
		CodeStreamType: ControllerMainStream,
		Channel:        1,
		Cmd:            CmdTurnOffAVTransmission,
	})
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}
