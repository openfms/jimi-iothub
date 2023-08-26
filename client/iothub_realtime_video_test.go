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

func TestIotHubClient_ListAVResourcesRequest(t *testing.T) {
	endPoint := os.Getenv("IOTHUB_ENDPOINT")
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(endPoint, "", "123456")
	assert.NilError(t, err)
	req := iothubcli.ListAVResourcesRequest(deviceImei, &AVResourceListCmdContent{
		Channel:       0,
		AlarmFlag:     0,
		ResourceType:  ResourceAudioAndVideo,
		CodeType:      CodeTypeAllStream,
		StorageType:   StorageTypeAllStorage,
		InstructionID: GenerateUniqueInstructionID(),
		BeginTime:     "230826113555",
		EndTime:       "230826113854",
	})
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}

func TestIotHubClient_HistoryVideoPlaybackRequest(t *testing.T) {
	endPoint := os.Getenv("IOTHUB_ENDPOINT")
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(endPoint, "", "123456")
	assert.NilError(t, err)
	req := iothubcli.HistoryVideoPlaybackRequest(deviceImei, &PlaybackCmdContent{
		InstructionID: GenerateUniqueInstructionID(),
		TCPPort:       "10003",
		UDPPort:       "0",
		Channel:       "1",
		ResourceType:  PlaybackResourceAudioAndVideo,
		CodeType:      PlaybackAllStream,
		StorageType:   PlaybackStorageAll,
		ForwardRewind: Invalid,
		PlayMethod:    PlayNormal,
		BeginTime:     "230826113555",
		EndTime:       "230826113854",
		ServerAddress: "192.168.1.1",
	})
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}

func TestIotHubClient_HistoryPlaybackControlRequest(t *testing.T) {
	endPoint := os.Getenv("IOTHUB_ENDPOINT")
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(endPoint, "", "123456")
	assert.NilError(t, err)
	req := iothubcli.HistoryPlaybackControlRequest(deviceImei, &PlaybackControlCmdContent{
		InstructionID: GenerateUniqueInstructionID(),
		Channel:       1,
		ForwardRewind: PlaybackSpeedInvalid,
		BeginTime:     "230826113555",
	})
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}
