package client

import (
	"context"
	"encoding/json"
)

type RealTimeControlCmd string

const (
	CmdTurnOffAVTransmission RealTimeControlCmd = "0"
	CmdSwitchStream          RealTimeControlCmd = "1"
	CmdPauseAllStreams       RealTimeControlCmd = "2"
	CmdResumeStream          RealTimeControlCmd = "3"
	CmdTurnOffIntercom       RealTimeControlCmd = "4"
)

type TurnOffAVType string

const (
	TurnOffBothAudioAndVideo TurnOffAVType = "0"
	TurnOffAudioOnly         TurnOffAVType = "1"
	TurnOffVideoOnly         TurnOffAVType = "2"
)

type RealTimeControllerCodeStreamType int

const (
	ControllerMainStream RealTimeControllerCodeStreamType = 0
	ControllerSubStream  RealTimeControllerCodeStreamType = 1
)

type RealTimeControlCmdContent struct {
	Channel        int                              `json:"channel"`
	Cmd            RealTimeControlCmd               `json:"cmd"`
	DataType       TurnOffAVType                    `json:"dataType"`
	CodeStreamType RealTimeControllerCodeStreamType `json:"codeStreamType"`
}

func (cli *IotHubClient) RealTimeAVControlRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *RealTimeControlCmdContent) (*InstructRequest, error) {
	if cmdContent == nil {
		return nil, ErrEmptyCmdContent
	}
	if deviceModel < DeviceModelJC450 {
		return nil, ErrUnsupportedRequest
	}
	if len(cmdContent.DataType) == 0 {
		cmdContent.DataType = TurnOffBothAudioAndVideo
	}
	if cmdContent.Channel == 0 {
		cmdContent.Channel = 1
	}
	if len(cmdContent.Cmd) == 0 {
		cmdContent.Cmd = CmdResumeStream
	}
	jsonData, _ := json.Marshal(cmdContent)
	req, err := cli.DeviceInstructionRequest(ctx, imei, string(jsonData))
	if err != nil {
		return nil, err
	}
	req.ProNo = ProNoAudioVideoTransmissionControl
	return req, nil
}
