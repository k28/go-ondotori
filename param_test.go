package ondotori

import (
	"reflect"
	"testing"
)

func testMapExpect(t *testing.T, expect interface{}, val interface{}) {
	if !reflect.DeepEqual(expect, val) {
		t.Fatal("expect [", expect, "] but [", val, "]")
	}
}

func makeTestBaseParam() BaseParam {
	return BaseParam{
		Token:     "token123",
		LoginId:   "ond123",
		LoginPass: "pass123",
	}
}

func makeTestBaseParamWithId(id string) BaseParam {
	return BaseParam{
		Token:     "token123",
		LoginId:   id,
		LoginPass: "pass123",
	}
}

func TestAddParams(t *testing.T) {
	cp := &BaseParam{
		Token:     "token123",
		LoginId:   "ond123",
		LoginPass: "pass123",
	}

	m := make(map[string]interface{})
	cp.AddParams(m)

	testMapExpect(t, "token123", m["api-key"])
	testMapExpect(t, "ond123", m["login-id"])
	testMapExpect(t, "pass123", m["login-pass"])
}

func TestCurrentParams(t *testing.T) {
	baseParam := makeTestBaseParam()

	cp := CurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{},
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, nil, r["remote-serial"])
	testMapExpect(t, nil, r["base-serial"])
}

func TestCurrentParams2(t *testing.T) {
	baseParam := makeTestBaseParam()

	cp := CurrentParam{
		RemoteSerial: []string{"remoteserial"},
		BaseSerial:   []string{},
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, []string{"remoteserial"}, r["remote-serial"])
	testMapExpect(t, nil, r["base-serial"])
}

func TestCurrentParams3(t *testing.T) {
	baseParam := makeTestBaseParam()

	cp := CurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{"baseserial"},
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, []string{"baseserial"}, r["base-serial"])
	testMapExpect(t, nil, r["remote-serial"])
}

func TestLatestDataParam(t *testing.T) {
	baseParam := makeTestBaseParam()
	cp := LatestDataParam{
		RemoteSerial: "5F52123C",
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, "5F52123C", r["remote-serial"])
}

func TestGetBaseUri(t *testing.T) {
	testBaseUrl(t, "tbac1234", true)
	testBaseUrl(t, "ondotori", true)
	testBaseUrl(t, "rbxx1234", true)
	testBaseUrl(t, "tdxx1234", false)
	testBaseUrl(t, "rdxx1234", false)
}

func testBaseUrl(t *testing.T, id string, isJapan bool) {
	b := makeTestBaseParamWithId(id)
	u := b.GetBaseURI()
	var expect = "https://api.webstorage.jp/v1/devices/"
	if !isJapan {
		expect = "https://api.webstorage-service.com/v1/devices/"
	}

	if u != expect {
		t.Fatal("expect [", expect, "] but [", u, "] id:[", id, "]")
	}
}

func TestMakeUri(t *testing.T) {
	b := makeTestBaseParam()

	p := CurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{},
	}

	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/current" {
		t.Fatal("make uri error")
	}
}

func TestMakeUriLatestData(t *testing.T) {
	b := makeTestBaseParam()

	p := LatestDataParam{
		RemoteSerial: "5236184A",
	}

	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/latest-data" {
		t.Fatal("make uri error")
	}
}
