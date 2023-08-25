package client

import (
	"encoding/json"
)

type RealTimeAVDataType string

const (
	AudioVideoDataType              RealTimeAVDataType = "0"
	VideoDataType                   RealTimeAVDataType = "1"
	TwoWayIntercomDataType          RealTimeAVDataType = "2"
	MonitorDataType                 RealTimeAVDataType = "3"
	CenterBroadcastDataType         RealTimeAVDataType = "4"
	TransparentTransmissionDataType RealTimeAVDataType = "5"
)

type RealTimeCodeStreamType string

const (
	MainStream RealTimeCodeStreamType = "0"
	SubStream  RealTimeCodeStreamType = "1"
)

type RealTimeCmdContent struct {
	DataType       RealTimeAVDataType     `json:"dataType"`
	CodeStreamType RealTimeCodeStreamType `json:"codeStreamType"`
	Channel        string                 `json:"channel"`
	VideoIP        string                 `json:"videoIP"`
	VideoTCPPort   string                 `json:"videoTCPPort"`
	VideoUDPPort   string                 `json:"videoUDPPort"`
}

func (cli *IotHubClient) RealTimeAVRequest(imei string, cmdContent *RealTimeCmdContent) *InstructRequest {
	if cmdContent == nil {
		cmdContent = &RealTimeCmdContent{
			DataType:       AudioVideoDataType,
			CodeStreamType: MainStream,
			Channel:        "1",
			VideoTCPPort:   "10002",
			VideoUDPPort:   "0",
		}
	}
	jsonData, _ := json.Marshal(cmdContent)
	req := cli.DeviceInstructionRequest(imei, string(jsonData))
	req.ProNo = ProNoRealTimeAudioVideoRequest
	return req
}
