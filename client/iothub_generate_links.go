package client

import (
	"fmt"
	"github.com/openfms/jimi-iothub/utils"
	"net"
)

type VideoLinks struct {
	RtmpLink, FlvLink string
}

func (cli *IotHubClient) GenerateHttpFlvLink(secure bool, prefix string, channel int, imei string) (string, error) {
	return utils.GenerateHttpFLVLink(secure, cli.endPointURL.String(), prefix, channel, imei)
}

func (cli *IotHubClient) GenerateRtmpLink(secure bool, prefix string, channel int, imei string) (string, error) {
	return utils.GenerateRtmpLink(secure, cli.endPointURL.String(), prefix, channel, imei)
}

func (cli *IotHubClient) GenerateVideoLinks(secure bool, prefix string, channel int, imei string) (*VideoLinks, error) {
	rtmpLink, err := utils.GenerateRtmpLink(secure, cli.endPointURL.String(), prefix, channel, imei)
	if err != nil {
		return nil, err
	}
	flvLink, err := utils.GenerateHttpFLVLink(secure, cli.endPointURL.String(), prefix, channel, imei)
	if err != nil {
		return nil, err
	}
	return &VideoLinks{
		RtmpLink: rtmpLink,
		FlvLink:  flvLink,
	}, nil
}

type DeviceConfigLinks struct {
	RtmpAddress    string
	HttpUploadLink string
}

func (cli *IotHubClient) GenerateDeviceConfigLinks(rtmpPrefix string) *DeviceConfigLinks {
	rtmpUrl := net.JoinHostPort(cli.GetEndpointHost(), cli.config.RtmpMediaServerPort)
	if len(rtmpPrefix) > 0 {
		rtmpUrl = rtmpUrl + "/" + rtmpPrefix
	}
	httpUploadUrl := fmt.Sprintf("http://%s/upload", net.JoinHostPort(cli.GetEndpointHost(), cli.config.FileStoragePort))
	return &DeviceConfigLinks{
		RtmpAddress:    rtmpUrl,
		HttpUploadLink: httpUploadUrl,
	}
}
