package client

import (
	"context"
	"github.com/openfms/jimi-iothub/utils"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_HistoryPlaybackControlRequest(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	ctx := context.Background()
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.HistoryPlaybackControlRequest(ctx, deviceImei, DeviceModelJC450, &PlaybackControlCmdContent{
		InstructionID: utils.GenerateUniqueInstructionID(),
		Channel:       1,
		ForwardRewind: PlaybackSpeedInvalid,
		BeginTime:     "230826113555",
	})
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(ctx, req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == ResponseCodeSuccess)
}
