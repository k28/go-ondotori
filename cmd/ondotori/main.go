package main

import (
	"fmt"
	"context"
	"io/ioutil"
	"encoding/json"
	ondotori "github.com/k28/go-ondotori"
)

type AccessInfo struct {
	Token	string `json:"api_key"`
	Id		string `json:"user_id"`
	Pass	string `json:"user_pass"`
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

	res, err := client.Get(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res.DeviceList)
}

