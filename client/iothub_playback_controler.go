package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openfms/jimi-iothub/utils"
)

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

func (cli *IotHubClient) HistoryPlaybackControlRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *PlaybackControlCmdContent) (*InstructRequest, error) {
	if cmdContent == nil {
		return nil, ErrEmptyCmdContent
	}
	if deviceModel < DeviceModelJC450 {
		return nil, ErrUnsupportedRequest
	}
	if len(cmdContent.InstructionID) == 0 {
		cmdContent.InstructionID = utils.GenerateUniqueInstructionID()
	}
	if cmdContent.Channel == 0 {
		cmdContent.Channel = 1
	}
	if len(cmdContent.BeginTime) == 0 {
		return nil, fmt.Errorf("field begin_time is empty")
	}
	jsonData, _ := json.Marshal(cmdContent)
	req, err := cli.DeviceInstructionRequest(ctx, imei, string(jsonData))
	if err != nil {
		return nil, err
	}
	req.ProNo = ProNoRemoteVideoPlaybackControl
	return req, nil
}
