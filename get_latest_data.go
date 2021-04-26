package ondotori

import (
	"encoding/json"
	"io"
)

type GetLatestDataParam struct {
	RemoteSerial string
}

func (param GetLatestDataParam) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
	p := make(map[string]interface{})

	baseParam.AddParams(p)

	if len(param.RemoteSerial) > 0 {
		p["remote-serial"] = param.RemoteSerial
	}

	return p
}

func (param GetLatestDataParam) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "latest-data"
}

func (param GetLatestDataParam) ParseResponse(reader io.Reader) (interface{}, error) {
	var body DeviceData
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
