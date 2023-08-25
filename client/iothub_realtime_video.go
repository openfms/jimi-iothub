package client

import (
	"encoding/json"
)

type RealTimeCmdContent struct {
	DataType       RealTimeAudioVideoDataType `json:"dataType"`
	CodeStreamType RealTimeCodeStreamType     `json:"codeStreamType"`
	Channel        string                     `json:"channel"`
	VideoIP        string                     `json:"videoIP"`
	VideoTCPPort   string                     `json:"videoTCPPort"`
	VideoUDPPort   string                     `json:"videoUDPPort"`
}

func (cli *IotHubClient) NewRealTimeAudioVideoRequest(imei string, cmdContent *RealTimeCmdContent) *InstructRequest {
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
	req := cli.NewDeviceInstructionRequest(imei, string(jsonData))
	req.ProNo = ProNoRealTimeAudioVideoRequest
	return req
}
