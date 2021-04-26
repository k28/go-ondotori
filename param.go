package ondotori

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

type makeParam interface {
	MakeJsonMap(baseParam BaseParam) map[string]interface{}
	MakeUri(baseParam BaseParam) string
	ParseResponse(reader io.Reader) (interface{}, error)
}

type BaseParam struct {
	Token     string
	LoginId   string
	LoginPass string
}

type CurrentParam struct {
	RemoteSerial []string
	BaseSerial   []string
}

type LatestDataParam struct {
	RemoteSerial string
}

type GetDataParam struct {
	RemoteSerial string
	From         *time.Time
	To           *time.Time
	Number       *uint16
}

func (param BaseParam) AddParams(src map[string]interface{}) {
	src["api-key"] = param.Token
	src["login-id"] = param.LoginId
	src["login-pass"] = param.LoginPass
}

func (param BaseParam) GetBaseURI() string {
	id := param.LoginId
	if strings.HasPrefix(id, "td") || strings.HasPrefix(id, "rd") {
		return "https://api.webstorage-service.com/v1/devices/"
	}
	return "https://api.webstorage.jp/v1/devices/"
}

func (param CurrentParam) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
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

func (param CurrentParam) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "current"
}

func (param CurrentParam) ParseResponse(reader io.Reader) (interface{}, error) {
	var body Devices
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func (param LatestDataParam) MakeJsonMap(baseParam BaseParam) map[string]interface{} {
	p := make(map[string]interface{})

	baseParam.AddParams(p)

	if len(param.RemoteSerial) > 0 {
		p["remote-serial"] = param.RemoteSerial
	}

	return p
}

func (param LatestDataParam) MakeUri(baseParam BaseParam) string {
	u := baseParam.GetBaseURI()
	return u + "latest-data"
}

func (param LatestDataParam) ParseResponse(reader io.Reader) (interface{}, error) {
	var body DeviceData
	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
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
