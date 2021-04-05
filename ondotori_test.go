package ondotori

import "testing"

func TestNew(t *testing.T) {
	token := "token"
	id := "tbac1234"
	pass := "password"
	client, err := New(token, id, pass)
	if err != nil {
		t.Fatal("err is not null.")
	}

	if client.token != token {
		t.Fatal("toke is not equal.")
	}

	if client.loginId != id {
		t.Fatal("id is not equal.")
	}

	if client.loginPass != pass {
		t.Fatal("pass is not equal.")
	}
}
