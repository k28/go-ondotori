package ondotori

import (
	"encoding/json"
	"io"
	"time"
)

type GetDataRTR500Param struct {
	RemoteSerial string
	BaseSerial   string
	From         *time.Time
	To           *time.Time
	Number       *uint16
}

func (param GetDataRTR500Param) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
	p := make(map[string]interface{})

	baseParam.AddParams(p)

	if len(param.RemoteSerial) > 0 {
		p["remote-serial"] = param.RemoteSerial
	}

	if len(param.BaseSerial) > 0 {
		p["base-serial"] = param.BaseSerial
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

func (param GetDataRTR500Param) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "data-rtr500"
}

func (param GetDataRTR500Param) ParseResponse(reader io.Reader) (interface{}, error) {
	var body DeviceDataRTR500
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
