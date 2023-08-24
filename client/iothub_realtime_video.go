package client

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

type RealTimeCmdContent struct {
	DataType       RealTimeAudioVideoDataType `url:"dataType"`
	CodeStreamType RealTimeCodeStreamType     `url:"codeStreamType"`
	Channel        string                     `url:"channel"`
	VideoIP        string                     `url:"videoIP"`
	VideoTCPPort   string                     `url:"videoTCPPort"`
	VideoUDPPort   string                     `url:"videoUDPPort"`
}

type RealTimeAudioVideoRequest struct {
	BaseInstructRequest
	CmdContent *RealTimeCmdContent `url:"cmdContent,required"`
}

func (cli *IotHubClient) NewRealTimeAudioVideoRequest(imei string, cmdContent *RealTimeCmdContent) *RealTimeAudioVideoRequest {
	if cmdContent == nil {
		cmdContent = &RealTimeCmdContent{
			DataType:       AudioVideoDataType,
			CodeStreamType: MainStream,
			Channel:        "1",
			VideoTCPPort:   "10002",
			VideoUDPPort:   "0",
		}
	}
	return &RealTimeAudioVideoRequest{
		CmdContent: cmdContent,
		BaseInstructRequest: BaseInstructRequest{
			DeviceIMEI:   imei,
			ProNo:        ProNoRealTimeAudioVideoRequest,
			Platform:     RequestPlatformWeb,
			CmdType:      NormallnsCommandType,
			Token:        cli.apiToken,
			OfflineFlag:  true,
			Timeout:      30,
			Sync:         true,
			RequestID:    getRequestID(),
			ServerFlagID: getServerFlagID(),
		},
	}
}
func (cli *IotHubClient) RealTimeAudioVideoTransmission(request *RealTimeAudioVideoRequest) (*Response, error) {
	values, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	// Send the POST request with x-www-form-urlencoded data
	resp, err := cli.client.R().
		SetBody(values.Encode()).
		Post(cli.endPointURL.String() + "/api/device/sendInstruct")

	if err != nil {
		return nil, err
	}
	apiResponse := &Response{}
	err = json.Unmarshal(resp.Body(), apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, nil
}
