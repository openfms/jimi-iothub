package client

import "encoding/json"

type PlaybackControl byte

const (
	PlaybackStart       PlaybackControl = 0
	PlaybackPause       PlaybackControl = 1
	PlaybackEnd         PlaybackControl = 2
	PlaybackFastForward PlaybackControl = 3
	PlaybackKeyframe    PlaybackControl = 4
	PlaybackDrag        PlaybackControl = 5
)

type PlaybackSpeed uint8

const (
	PlaybackSpeedInvalid PlaybackSpeed = 0
	PlaybackSpeed1x      PlaybackSpeed = 1
	PlaybackSpeed2x      PlaybackSpeed = 2
	PlaybackSpeed4x      PlaybackSpeed = 3
	PlaybackSpeed8x      PlaybackSpeed = 4
	PlaybackSpeed16x     PlaybackSpeed = 5
)

type PlaybackControlCmdContent struct {
	Channel       uint8           `json:"channel"` //1-5
	PlayCtrl      PlaybackControl `json:"playCtrl"`
	ForwardRewind PlaybackSpeed   `json:"forwardRewind"`
	BeginTime     string          `json:"beginTime"`
	//Drag position(YYMMDDHHMMSS, when playback control is 5, this field is valid)
	InstructionID string `json:"instructionID"`
}

func (cli *IotHubClient) HistoryPlaybackControlRequest(imei string, cmdContent *PlaybackControlCmdContent) *InstructRequest {
	if cmdContent == nil {
		cmdContent = &PlaybackControlCmdContent{
			InstructionID: GenerateUniqueInstructionID(),
			Channel:       1,
			ForwardRewind: PlaybackSpeedInvalid,
			PlayCtrl:      PlaybackStart,
		}
	}
	jsonData, _ := json.Marshal(cmdContent)
	req := cli.DeviceInstructionRequest(imei, string(jsonData))
	req.ProNo = ProNoRemoteVideoPlaybackControl
	return req
}
