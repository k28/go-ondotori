package ondotori

type DeviceData struct {
	Serial   string    `json:"serial"`
	Model    string    `json:"model"`
	TimeDiff string    `json:"time_diff"`
	StdBias  string    `json:"std_bias"`
	DstBias  string    `json:"dst_bias"`
	Name     string    `json:"name"`
	Channel  []Channel `json:"channel"`
	Data     []Data    `json:"data"`
}

type Channel struct {
	Name string `json:"name"`
	Num  string `json:"num"`
	Unit string `json:"unit"`
}

type Data struct {
	Unixtime string `json:"unixtime"`
	DataId   string `json:"data-id"`
	Ch1      string `json:"ch1"`
	Ch2      string `json:"ch2"`
}
