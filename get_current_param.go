package ondotori

import (
	"encoding/json"
	"io"
)

type GetCurrentParam struct {
	RemoteSerial []string
	BaseSerial   []string
}

func (param GetCurrentParam) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
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

func (param GetCurrentParam) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "current"
}

func (param GetCurrentParam) ParseResponse(reader io.Reader) (interface{}, error) {
	var body Devices
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
