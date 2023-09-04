package client

import (
	"context"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_RealTimeAVControlRequest(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	ctx := context.Background()
	iothubcli, err := NewIotHubClient(env, nil)
	assert.NilError(t, err)
	req, err := iothubcli.RealTimeAVControlRequest(ctx, deviceImei, DeviceModelJC450, &RealTimeControlCmdContent{
		DataType:       TurnOffBothAudioAndVideo,
		CodeStreamType: ControllerMainStream,
		Channel:        1,
		Cmd:            CmdTurnOffAVTransmission,
	})
	resp, err := iothubcli.SendDeviceInstruction(ctx, req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == ResponseCodeSuccess)
}
