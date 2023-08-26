package client

import "encoding/json"

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
	Invalid ForwardRewind = "0"
	X1      ForwardRewind = "1"
	X2      ForwardRewind = "2"
	X4      ForwardRewind = "3"
	X8      ForwardRewind = "4"
	X16     ForwardRewind = "5"
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

func (cli *IotHubClient) HistoryVideoPlaybackRequest(imei string, cmdContent *PlaybackCmdContent) *InstructRequest {
	if cmdContent == nil {
		cmdContent = &PlaybackCmdContent{
			InstructionID: GenerateUniqueInstructionID(),
			TCPPort:       "10003",
			UDPPort:       "0",
			Channel:       "1",
			ResourceType:  PlaybackResourceAudioAndVideo,
			CodeType:      PlaybackAllStream,
			StorageType:   PlaybackStorageAll,
			ForwardRewind: Invalid,
			PlayMethod:    PlayNormal,
		}
	}
	jsonData, _ := json.Marshal(cmdContent)
	req := cli.DeviceInstructionRequest(imei, string(jsonData))
	req.ProNo = ProNoRemoteVideoPlaybackRequest
	return req
}
