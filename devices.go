package ondotori

import "errors"

type Devices struct {
	DeviceList []Device `json:"devices"`
}

type Device struct {
	Num string `json:"num"`
	Serial string `json:"serial"`
	Model string `json:"model"`
	Name string `json:"name"`
	Battery string `json:"battery"`
	Rssi string `json:"rssi"`
	TimeDiff string `json:"time_diff"`
	StdBias string `json:"std_bias"`
	DstBias string `json:"dst_bias"`
	UnixTime string `json:"unixtime"`
	Channel []Record `json:"channel"`
	BaseUnit BaseUnit `json:"baseunit"`
	Group Group `json:"group"`
}

type Record struct {
	Num string    `json:"num"`
	Name string   `json:"name"`
	Value string  `json:"value"`
	Unit string   `json:"unit"`
}

type BaseUnit struct {
	Serial string `json:"serial"`
	Model string  `json:"model"`
	Name string   `json:"name"`
}

type Group struct {
	Num	string	`json:"num"`
	Name string `json:"name"`
}

func (devices *Devices) GetDevice(serial string) (*Device, error) {
	for _, d := range devices.DeviceList {
		if d.Serial == serial {
			return &d, nil
		}
	}

	return nil, errors.New("not found")
}

