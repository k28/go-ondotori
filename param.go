package ondotori

import (
	"io"
	"strings"
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
