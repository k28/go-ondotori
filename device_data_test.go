package ondotori

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestLatestData(t *testing.T) {
	js, err := ioutil.ReadFile("./test_resources/device_data.json")
	if err != nil {
		t.Fatal("jsondevice.json not found.", err.Error())
	}

	var ld DeviceData
	json.Unmarshal([]byte(js), &ld)

	testLatestData(t, &ld)
}

func testLatestData(t *testing.T, ld *DeviceData) {
	testEquals(t, "5236XXXX", ld.Serial)
	testEquals(t, "TR-72wb", ld.Model)
	testEquals(t, "Living TR-72wb", ld.Name)
	testEquals(t, "540", ld.TimeDiff)
	testEquals(t, "60", ld.StdBias)
	testEquals(t, "30", ld.DstBias)
}

func testChannel(t *testing.T, ld *DeviceData) {
	var ch = ld.Channel[0]
	testEquals(t, "1", ch.Num)
	testEquals(t, "temp", ch.Name)
	testEquals(t, "C", ch.Unit)
}

func testData(t *testing.T, ld *DeviceData) {
	var data = ld.Data[0]
	testEquals(t, "1619040000", data.Unixtime)
	testEquals(t, "20660", data.DataId)
	testEquals(t, "22.7", data.Ch1)
	testEquals(t, "27", data.Ch2)
}

func TestLatestDataRTR500(t *testing.T) {
	js, err := ioutil.ReadFile("./test_resources/device_data_rtr500.json")
	if err != nil {
		t.Fatal("jsondevice.json not found.", err.Error())
	}

	var ld DeviceDataRTR500
	json.Unmarshal([]byte(js), &ld)

	testEquals(t, "52800010", ld.RemoteSerial)
	testEquals(t, "RTR501B", ld.RemoteModel)
	testEquals(t, "外気", ld.RemoteName)
	testEquals(t, "5858001E", ld.BaseSerial)
	testEquals(t, "RTR500BW", ld.BaseModel)
	testEquals(t, "k28home", ld.BaseName)
	testEquals(t, "540", ld.TimeDiff)
	testEquals(t, "60", ld.StdBias)
	testEquals(t, "30", ld.DstBias)

	ch := ld.Channel[0]
	testEquals(t, "ch1", ch.Name)
	testEquals(t, "1", ch.Num)
	testEquals(t, "C", ch.Unit)

	dt := ld.Data[0]
	testEquals(t, "1619463974", dt.Unixtime)
	testEquals(t, "34165", dt.DataId)
	testEquals(t, "0.6", dt.Ch1)
}
