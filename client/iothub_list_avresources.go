package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openfms/jimi-iothub/utils"
)

type AVResourceListCmdContent struct {
	Channel       uint8                   `json:"channel"`
	BeginTime     string                  `json:"beginTime"`
	EndTime       string                  `json:"endTime"`
	AlarmFlag     uint64                  `json:"alarmFlag"`
	ResourceType  ResourcesType           `json:"resourceType"`
	CodeType      ResourceListCodeType    `json:"codeType"`
	StorageType   ResourceListStorageType `json:"storageType"`
	InstructionID string                  `json:"instructionID"`
}
type ResourcesType uint8

const (
	ResourceAudioAndVideo ResourcesType = 0
	ResourceAudio         ResourcesType = 1
	ResourceVideo         ResourcesType = 2
	ResourceVideoOrAudio  ResourcesType = 3
)

type ResourceListCodeType uint8

const (
	CodeTypeAllStream  ResourceListCodeType = 0
	CodeTypeMainStream ResourceListCodeType = 1
	CodeTypeSubStream  ResourceListCodeType = 2
)

type ResourceListStorageType uint8

const (
	StorageTypeAllStorage              ResourceListStorageType = 0
	StorageTypeMainStorage             ResourceListStorageType = 1
	StorageTypeDisasterRecoveryStorage ResourceListStorageType = 2
)

func (cli *IotHubClient) ListAVResourcesRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *AVResourceListCmdContent) (*InstructRequest, error) {
	if cmdContent == nil {
		return nil, ErrEmptyCmdContent
	}
	if deviceModel < DeviceModelJC450 {
		return nil, ErrUnsupportedRequest
	}
	if len(cmdContent.InstructionID) == 0 {
		cmdContent.InstructionID = utils.GenerateUniqueInstructionID()
	}
	if len(cmdContent.BeginTime) == 0 {
		return nil, fmt.Errorf("field begin_time is empty")
	}
	if len(cmdContent.EndTime) == 0 {
		return nil, fmt.Errorf("field end_time is empty")
	}
	jsonData, _ := json.Marshal(cmdContent)
	req, err := cli.DeviceInstructionRequest(ctx, imei, string(jsonData))
	if err != nil {
		return nil, err
	}
	req.ProNo = ProNoQueryAudioVideoResourceList
	return req, nil
}
