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

	if client.baseParam.Token != token {
		t.Fatal("token is not equal.")
	}

	if client.baseParam.LoginId != id {
		t.Fatal("id is not equal.")
	}

	if client.baseParam.LoginPass != pass {
		t.Fatal("pass is not equal.")
	}
}
