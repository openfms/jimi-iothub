package client

import (
	"context"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestIotHubClient_RealTimeAudioVideoTransmission(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	ctx := context.Background()
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.RealTimeAVRequest(ctx, deviceImei, DeviceModelJC450, &RealTimeCmdContent{
		DataType:       AudioVideoDataType,
		CodeStreamType: MainStream,
		VideoUDPPort:   0,
		VideoIP:        "192.168.1.1",
		VideoTCPPort:   "10002",
		Channel:        "1",
	})
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(ctx, req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == ResponseCodeSuccess)
}
