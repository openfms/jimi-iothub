package client

import (
	"context"
	"encoding/json"
	"fmt"
)

type PlaybackResourceType string

const (
	PlaybackResourceAudioAndVideo PlaybackResourceType = "0"
	PlaybackResourceAudio         PlaybackResourceType = "1"
	PlaybackResourceVideo         PlaybackResourceType = "2"
	PlaybackResourceAudioOrVideo  PlaybackResourceType = "3"
)

type PlaybackCodeType string

const (
	PlaybackAllStream  PlaybackCodeType = "0"
	PlaybackMainStream PlaybackCodeType = "1"
	PlaybackSubStream  PlaybackCodeType = "2"
)

type PlaybackStorageType string

const (
	PlaybackStorageAll  PlaybackStorageType = "0"
	PlaybackStorageMain PlaybackStorageType = "1"
	StorageDisaster     PlaybackStorageType = "2"
)

type PlayMethod string

const (
	PlayNormal         PlayMethod = "0"
	PlayFastForward    PlayMethod = "1"
	PlayKeyframeRewind PlayMethod = "2"
	PlayKeyframe       PlayMethod = "3"
	PlaySingleFrame    PlayMethod = "4"
)

type ForwardRewind string

const (
	ForwardRewindInvalid ForwardRewind = "0"
	ForwardRewindX1      ForwardRewind = "1"
	ForwardRewindX2      ForwardRewind = "2"
	ForwardRewindX4      ForwardRewind = "3"
	ForwardRewindX8      ForwardRewind = "4"
	ForwardRewindX16     ForwardRewind = "5"
)

type PlaybackCmdContent struct {
	ServerAddress string               `json:"serverAddress"`
	TCPPort       string               `json:"tcpPort"`
	UDPPort       string               `json:"udpPort"`
	Channel       string               `json:"channel"`
	ResourceType  PlaybackResourceType `json:"resourceType"`
	CodeType      PlaybackCodeType     `json:"codeType"`
	StorageType   PlaybackStorageType  `json:"storageType"`
	PlayMethod    PlayMethod           `json:"playMethod"`
	ForwardRewind ForwardRewind        `json:"forwardRewind"`
	BeginTime     string               `json:"beginTime"`
	EndTime       string               `json:"endTime"`
	InstructionID string               `json:"instructionID"`
}

func (cli *IotHubClient) HistoryVideoPlaybackRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *PlaybackCmdContent) (*InstructRequest, error) {
	if cmdContent == nil {
		return nil, ErrEmptyCmdContent
	}
	if deviceModel < DeviceModelJC450 {
		return nil, ErrUnsupportedRequest
	}
	if len(cmdContent.ResourceType) == 0 {
		cmdContent.ResourceType = PlaybackResourceAudioAndVideo
	}
	if len(cmdContent.CodeType) == 0 {
		cmdContent.CodeType = PlaybackAllStream
	}
	if len(cmdContent.StorageType) == 0 {
		cmdContent.StorageType = PlaybackStorageAll
	}
	if len(cmdContent.Channel) == 0 {
		cmdContent.Channel = "1"
	}
	if len(cmdContent.TCPPort) == 0 {
		cmdContent.TCPPort = cli.config.HistoryVideoPort
	}
	if len(cmdContent.ForwardRewind) == 0 {
		cmdContent.ForwardRewind = ForwardRewindInvalid
	}
	if len(cmdContent.PlayMethod) == 0 {
		cmdContent.PlayMethod = PlayNormal
	}
	if len(cmdContent.ServerAddress) == 0 {
		cmdContent.ServerAddress = cli.GetEndpointHost()
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
	req.ProNo = ProNoRemoteVideoPlaybackRequest
	return req, nil
}
