package ondotori

import "testing"

func TestLatestDataParam(t *testing.T) {
	baseParam := makeTestBaseParam()
	cp := GetLatestDataParam{
		RemoteSerial: "5F52123C",
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, "5F52123C", r["remote-serial"])
}

func TestMakeUriLatestData(t *testing.T) {
	b := makeTestBaseParam()

	p := GetLatestDataParam{
		RemoteSerial: "5236184A",
	}

	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/latest-data" {
		t.Fatal("make uri error")
	}
}
