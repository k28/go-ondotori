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
