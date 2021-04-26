package ondotori

import (
	"testing"
	"time"
)

func TestMakeUriGetData(t *testing.T) {
	b := makeTestBaseParam()
	p := GetDataParam{
		RemoteSerial: "5236184A",
	}
	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/data" {
		t.Fatal("make uri error")
	}
}

func TestMakeJsonMapGetData(t *testing.T) {

	b := makeTestBaseParam()
	var p = GetDataParam{
		RemoteSerial: "5236184A",
	}

	var m = p.MakeJsonMap(b)
	testEquals(t, "5236184A", m["remote-serial"])
	testEquals(t, nil, m["unixtime-from"])
	testEquals(t, nil, m["unixtime-to"])

	from := time.Unix(1234567, 0)
	to := time.Unix(2234567, 0)
	p = GetDataParam{
		RemoteSerial: "12345678",
		From:         &from,
		To:           &to,
	}
	m = p.MakeJsonMap(b)
	testEquals(t, "12345678", m["remote-serial"])
	testEquals(t, int64(1234567), m["unixtime-from"])
	testEquals(t, int64(2234567), m["unixtime-to"])

	limit := uint16(10)
	p = GetDataParam{
		RemoteSerial: "12345678",
		Number:       &limit,
	}
	m = p.MakeJsonMap(b)
	testEquals(t, "12345678", m["remote-serial"])
	testEquals(t, uint16(10), m["number"])
}
