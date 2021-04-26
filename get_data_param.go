package ondotori

import (
	"encoding/json"
	"io"
	"time"
)

type GetDataParam struct {
	RemoteSerial string
	From         *time.Time
	To           *time.Time
	Number       *uint16
}

func (param GetDataParam) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
	p := make(map[string]interface{})

	baseParam.AddParams(p)

	if len(param.RemoteSerial) > 0 {
		p["remote-serial"] = param.RemoteSerial
	}

	if param.From != nil && param.To != nil {
		p["unixtime-from"] = param.From.Unix()
		p["unixtime-to"] = param.To.Unix()
	}

	if param.Number != nil {
		p["number"] = *param.Number
	}

	return p
}

func (param GetDataParam) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "data"
}

func (param GetDataParam) ParseResponse(reader io.Reader) (interface{}, error) {
	var body DeviceData
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
