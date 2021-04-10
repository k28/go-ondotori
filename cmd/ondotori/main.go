package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	ondotori "github.com/k28/go-ondotori"
)

type AccessInfo struct {
	Token string `json:"api_key"`
	Id    string `json:"user_id"`
	Pass  string `json:"user_pass"`
}

func main() {
	do_access_ondotori()
}

func do_access_ondotori() {
	fmt.Println("Hello World")

	raw, err := ioutil.ReadFile("/var/tmp/webstorage.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var ac AccessInfo
	json.Unmarshal(raw, &ac)

	fmt.Println(ac.Token)

	client, err := ondotori.New(ac.Token, ac.Id, ac.Pass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cp := ondotori.CurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{},
	}

	res, err := client.Get(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res.DeviceList)
}
