package ondotori

import "testing"

func TestCurrentParams(t *testing.T) {
	baseParam := makeTestBaseParam()

	cp := GetCurrentParam{
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

	cp := GetCurrentParam{
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

	cp := GetCurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{"baseserial"},
	}

	r := cp.MakeJsonMap(baseParam)
	testMapExpect(t, "token123", r["api-key"])
	testMapExpect(t, []string{"baseserial"}, r["base-serial"])
	testMapExpect(t, nil, r["remote-serial"])
}

func TestMakeUri(t *testing.T) {
	b := makeTestBaseParam()

	p := GetCurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{},
	}

	if p.MakeUri(b) != "https://api.webstorage.jp/v1/devices/current" {
		t.Fatal("make uri error")
	}
}
