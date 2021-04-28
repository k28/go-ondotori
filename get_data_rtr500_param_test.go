package ondotori

import (
	"testing"
	"time"
)

func TestMakeUriGetDataRTR500(t *testing.T) {
	b := makeTestBaseParam()
	p := GetDataRTR500Param{
		RemoteSerial: "5236184A",
		BaseSerial:   "B1234567",
	}
	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/data-rtr500" {
		t.Fatal("make uri error")
	}
}

func TestMakeJsonMapGetDataRTR500(t *testing.T) {

	b := makeTestBaseParam()
	var p = GetDataRTR500Param{
		RemoteSerial: "5236184A",
		BaseSerial:   "B1234567",
	}

	var m = p.MakeJsonMap(b)
	testEquals(t, "5236184A", m["remote-serial"])
	testEquals(t, "B1234567", m["base-serial"])
	testEquals(t, nil, m["unixtime-from"])
	testEquals(t, nil, m["unixtime-to"])

	from := time.Unix(1234567, 0)
	to := time.Unix(2234567, 0)
	p = GetDataRTR500Param{
		RemoteSerial: "12345678",
		BaseSerial:   "B1234567",
		From:         &from,
		To:           &to,
	}
	m = p.MakeJsonMap(b)
	testEquals(t, "12345678", m["remote-serial"])
	testEquals(t, "B1234567", m["base-serial"])
	testEquals(t, int64(1234567), m["unixtime-from"])
	testEquals(t, int64(2234567), m["unixtime-to"])

	limit := uint16(10)
	p = GetDataRTR500Param{
		RemoteSerial: "12345678",
		BaseSerial:   "B1234567",
		Number:       &limit,
	}
	m = p.MakeJsonMap(b)
	testEquals(t, "12345678", m["remote-serial"])
	testEquals(t, "B1234567", m["base-serial"])
	testEquals(t, uint16(10), m["number"])
}
