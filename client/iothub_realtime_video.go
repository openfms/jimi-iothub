package client

import (
	"context"
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

func (cli *IotHubClient) RealTimeAVRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *RealTimeCmdContent) (*InstructRequest, error) {
	if cmdContent == nil {
		return nil, ErrEmptyCmdContent
	}
	if deviceModel < DeviceModelJC450 {
		return nil, ErrUnsupportedRequest
	}
	if len(cmdContent.DataType) == 0 {
		cmdContent.DataType = AudioVideoDataType
	}
	if len(cmdContent.CodeStreamType) == 0 {
		cmdContent.CodeStreamType = MainStream
	}
	if len(cmdContent.Channel) == 0 {
		cmdContent.Channel = "0"
	}
	if len(cmdContent.VideoTCPPort) == 0 {
		cmdContent.VideoTCPPort = cli.config.LiveVideoPort
	}
	if len(cmdContent.VideoIP) == 0 {
		cmdContent.VideoIP = cli.endPointURL.Host
	}
	jsonData, _ := json.Marshal(cmdContent)
	req, err := cli.DeviceInstructionRequest(ctx, imei, string(jsonData))
	if err != nil {
		return nil, err
	}
	req.ProNo = ProNoRealTimeAudioVideoRequest
	return req, nil
}
