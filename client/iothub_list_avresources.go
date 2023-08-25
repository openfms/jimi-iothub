package client

import (
	"encoding/json"
	"fmt"
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

func (cli *IotHubClient) ListAVResourcesRequest(imei string, cmdContent *AVResourceListCmdContent) *InstructRequest {
	if cmdContent == nil {
		cmdContent = &AVResourceListCmdContent{
			Channel:       0,
			AlarmFlag:     0,
			ResourceType:  ResourceAudioAndVideo,
			CodeType:      CodeTypeAllStream,
			StorageType:   StorageTypeAllStorage,
			InstructionID: fmt.Sprintf("%d", GenerateUniqueInstructionID()),
		}
	}
	jsonData, _ := json.Marshal(cmdContent)
	req := cli.DeviceInstructionRequest(imei, string(jsonData))
	req.ProNo = ProNoQueryAudioVideoResourceList
	return req
}
