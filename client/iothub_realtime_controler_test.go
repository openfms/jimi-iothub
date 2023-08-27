package client

import (
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

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
