package client

import (
	"github.com/openfms/jimi-iothub/commands"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_SendDeviceInstruction(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iothubcli.DeviceInstructionRequest(deviceImei, commands.GenerateCommand(commands.STATUS))
	assert.NilError(t, err)
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}
