package client

import (
	"github.com/openfms/jimi-iothub/commands"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_SendDeviceInstruction(t *testing.T) {
	endPoint := os.Getenv("IOTHUB_ENDPOINT")
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	iothubcli, err := NewIotHubClient(endPoint, "", "123456")
	assert.NilError(t, err)
	req := iothubcli.NewDeviceInstructionRequest(deviceImei, commands.GenerateCommand(commands.STATUS))
	resp, err := iothubcli.SendDeviceInstruction(req)
	assert.NilError(t, err)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == Success)
}
