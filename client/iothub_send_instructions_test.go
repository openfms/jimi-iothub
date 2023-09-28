package client

import (
	"context"
	"github.com/openfms/jimi-iothub/commands"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_SendDeviceInstruction(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	ctx := context.Background()

	assert.NilError(t, err)
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.DeviceInstructionRequest(ctx, deviceImei, commands.GenerateCommand(commands.STATUS))
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(ctx, req)
	assert.NilError(t, err)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == ResponseCodeSuccess)
}
