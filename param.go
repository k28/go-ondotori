package ondotori

type makeParam interface {
	MakeJsonMap(baseParam BaseParam) map[string]interface{}
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

func (param *BaseParam) AddParams(src map[string]interface{}) {
	src["token"] = param.Token
	src["login-id"] = param.LoginId
	src["login-pass"] = param.LoginPass
}

func (param *CurrentParam) MakeJsonMap(baseParam *BaseParam) map[string]interface{} {
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
