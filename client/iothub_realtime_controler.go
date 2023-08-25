package client

import "encoding/json"

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

func (cli *IotHubClient) RealTimeAVControlRequest(imei string, cmdContent *RealTimeControlCmdContent) *InstructRequest {
	if cmdContent == nil {
		cmdContent = &RealTimeControlCmdContent{
			Channel:        1,
			Cmd:            CmdTurnOffAVTransmission,
			DataType:       TurnOffBothAudioAndVideo,
			CodeStreamType: ControllerMainStream,
		}
	}
	jsonData, _ := json.Marshal(cmdContent)
	req := cli.DeviceInstructionRequest(imei, string(jsonData))
	req.ProNo = ProNoAudioVideoTransmissionControl
	return req
}
