package ondotori

import (
	"encoding/json"
	"io"
)

type GetLatestDataRTR500Param struct {
	RemoteSerial string
	BaseSerial   string
}

func (param GetLatestDataRTR500Param) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
	p := make(map[string]interface{})

	baseParam.AddParams(p)

	if len(param.RemoteSerial) > 0 {
		p["remote-serial"] = param.RemoteSerial
	}

	if len(param.BaseSerial) > 0 {
		p["base-serial"] = param.BaseSerial
	}

	return p
}

func (param GetLatestDataRTR500Param) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "latest-data-rtr500"
}

func (param GetLatestDataRTR500Param) ParseResponse(reader io.Reader) (interface{}, error) {
	var body DeviceDataRTR500
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
