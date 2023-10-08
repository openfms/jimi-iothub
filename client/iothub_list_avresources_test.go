package client

import (
	"context"
	"github.com/openfms/jimi-iothub/utils"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestIotHubClient_ListAVResourcesRequest(t *testing.T) {
	env, err := ReadIotHubEnvironments()
	assert.NilError(t, err)
	deviceImei := os.Getenv("IOTHUB_DEVICE_IMEI")
	ctx := context.Background()
	iotHubCli, err := NewIotHubClient(env)
	assert.NilError(t, err)
	req, err := iotHubCli.ListAVResourcesRequest(ctx, deviceImei, DeviceModelJC450, &AVResourceListCmdContent{
		Channel:       0,
		AlarmFlag:     0,
		ResourceType:  ResourceAudioAndVideo,
		CodeType:      CodeTypeAllStream,
		StorageType:   StorageTypeAllStorage,
		InstructionID: utils.GenerateUniqueInstructionID(),
		BeginTime:     "230826113555",
		EndTime:       "230826113854",
	})
	assert.NilError(t, err)
	resp, err := iotHubCli.SendDeviceInstruction(ctx, req)
	assert.NilError(t, err)
	t.Log(resp)
	assert.Assert(t, resp.Code == 0)
	assert.Assert(t, resp.Data.Code == ResponseCodeSuccess)
}
