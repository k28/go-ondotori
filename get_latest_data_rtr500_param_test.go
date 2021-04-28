package ondotori

import "testing"

func TestLatestDataRTR500Param(t *testing.T) {
	baseParam := makeTestBaseParam()
	cp := GetLatestDataRTR500Param{
		RemoteSerial: "R1234567",
		BaseSerial:   "B1234567",
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, "R1234567", r["remote-serial"])
	testMapExpect(t, "B1234567", r["base-serial"])
}

func TestMakeUriLatestDataRTR500(t *testing.T) {
	b := makeTestBaseParam()
	p := GetLatestDataRTR500Param{
		RemoteSerial: "R1234567",
		BaseSerial:   "B1234567",
	}

	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/latest-data-rtr500" {
		t.Fatal("make uri error")
	}
}
