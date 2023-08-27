package client

import (
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_RealTimeAudioVideoTransmission(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.RealTimeAVRequest(deviceImei, DeviceModelJC450, &RealTimeCmdContent{
		DataType:       AudioVideoDataType,
		CodeStreamType: MainStream,
		VideoUDPPort:   "0",
		VideoIP:        "192.168.1.1",
		VideoTCPPort:   "10002",
		Channel:        "1",
	})
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}

func TestIotHubClient_RealTimeAVControlRequest(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.RealTimeAVControlRequest(deviceImei, DeviceModelJC450, &RealTimeControlCmdContent{
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
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iotHubCli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iotHubCli.ListAVResourcesRequest(deviceImei, DeviceModelJC450, &AVResourceListCmdContent{
		Channel:       0,
		AlarmFlag:     0,
		ResourceType:  ResourceAudioAndVideo,
		CodeType:      CodeTypeAllStream,
		StorageType:   StorageTypeAllStorage,
		InstructionID: GenerateUniqueInstructionID(),
		BeginTime:     "230826113555",
		EndTime:       "230826113854",
	})
	assert.NilError(t, err)
	resp, err := iotHubCli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}

func TestIotHubClient_HistoryVideoPlaybackRequest(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.HistoryVideoPlaybackRequest(deviceImei, DeviceModelJC450, &PlaybackCmdContent{
		InstructionID: GenerateUniqueInstructionID(),
		TCPPort:       "10003",
		UDPPort:       "0",
		Channel:       "1",
		ResourceType:  PlaybackResourceAudioAndVideo,
		CodeType:      PlaybackAllStream,
		StorageType:   PlaybackStorageAll,
		ForwardRewind: ForwardRewindInvalid,
		PlayMethod:    PlayNormal,
		BeginTime:     "230826113555",
		EndTime:       "230826113854",
		ServerAddress: "192.168.1.1",
	})
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}

func TestIotHubClient_HistoryPlaybackControlRequest(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.HistoryPlaybackControlRequest(deviceImei, DeviceModelJC450, &PlaybackControlCmdContent{
		InstructionID: GenerateUniqueInstructionID(),
		Channel:       1,
		ForwardRewind: PlaybackSpeedInvalid,
		BeginTime:     "230826113555",
	})
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}
