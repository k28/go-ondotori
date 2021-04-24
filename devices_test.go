package ondotori

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func testEquals(t *testing.T, expect interface{}, value interface{}) {
	if expect != value {
		t.Fatal("expect [", expect, "] but [", value, "]")
	}
}

func TestDevices(t *testing.T) {
	js, err := ioutil.ReadFile("./devices.json")
	if err != nil {
		t.Fatal("jsondevice.json not found.", err.Error())
	}

	var device Devices
	json.Unmarshal([]byte(js), &device)

	if len(device.DeviceList) != 10 {
		t.Fatal("device.DeviceList", device.DeviceList)
	}

	testTR7Device(t, &device)
	testTR4Device(t, &device)
	testRTR500Device(t, &device)
}

func testTR7Device(t *testing.T, device *Devices) {
	nagaimo, _ := device.GetDevice("3214XXXX")
	testEquals(t, "1", nagaimo.Num)
	testEquals(t, "TR-72wf", nagaimo.Model)
	testEquals(t, "5", nagaimo.Battery)
	testEquals(t, "", nagaimo.Rssi)
	testEquals(t, "540", nagaimo.TimeDiff)
	testEquals(t, "0", nagaimo.StdBias)
	testEquals(t, "60", nagaimo.DstBias)
	testEquals(t, "1617526828", nagaimo.UnixTime)

	channels := nagaimo.Channel
	ch1 := channels[0]
	testEquals(t, "1", ch1.Num)
	testEquals(t, "Temp", ch1.Name)
	testEquals(t, "13.8", ch1.Value)
	testEquals(t, "C", ch1.Unit)

	ch2 := channels[1]
	testEquals(t, "2", ch2.Num)
	testEquals(t, "Humi", ch2.Name)
	testEquals(t, "69", ch2.Value)
	testEquals(t, "%", ch2.Unit)
}

func testTR4Device(t *testing.T, device *Devices) {
	xtrail, _ := device.GetDevice("582CXXXX")
	testEquals(t, "1", xtrail.Num)
	testEquals(t, "TR41", xtrail.Model)
	testEquals(t, "X-Trail", xtrail.Name)
	testEquals(t, "5", xtrail.Battery)
	testEquals(t, "", xtrail.Rssi)
	testEquals(t, "540", xtrail.TimeDiff)
	testEquals(t, "0", xtrail.StdBias)
	testEquals(t, "0", xtrail.DstBias)
	testEquals(t, "1617428258", xtrail.UnixTime)

	ch := xtrail.Channel[0]
	testEquals(t, "1", ch.Num)
	testEquals(t, "32.4", ch.Value)
	testEquals(t, "C", ch.Unit)
}

func testRTR500Device(t *testing.T, device *Devices) {
	rtr500, _ := device.GetDevice("5280XXXX")
	testEquals(t, "1", rtr500.Num)
	testEquals(t, "RTR501B", rtr500.Model)
	testEquals(t, "外気", rtr500.Name)
	testEquals(t, "5", rtr500.Battery)
	testEquals(t, "5", rtr500.Rssi)
	testEquals(t, "540", rtr500.TimeDiff)
	testEquals(t, "0", rtr500.StdBias)
	testEquals(t, "1617527147", rtr500.UnixTime)

	ch := rtr500.Channel[0]
	testEquals(t, "1", ch.Num)
	testEquals(t, "", ch.Name)
	testEquals(t, "13.3", ch.Value)
	testEquals(t, "C", ch.Unit)

	baseUnit := rtr500.BaseUnit
	testEquals(t, "5858XXXX", baseUnit.Serial)
	testEquals(t, "RTR500BW", baseUnit.Model)
	testEquals(t, "k28home", baseUnit.Name)

	gp := rtr500.Group
	testEquals(t, "1", gp.Num)
	testEquals(t, "Group1", gp.Name)
}
