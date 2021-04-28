package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	ondotori "github.com/k28/go-ondotori"
)

type AccessInfo struct {
	Token string `json:"api-key"`
	Id    string `json:"login-id"`
	Pass  string `json:"login-pass"`
}

func main() {
	// do_get_current()
	// do_get_latest_info()
	// do_get_data()
	// do_get_latest_rtr500()
	do_get_data_rtr500()
}

func load_access_info() (*AccessInfo, error) {
	// webstorage.json like this
	// {
	//     "api-key"    : "xxxxxxxxxxxxxx",
	//     "login-id"   : "rbxx1234",
	//     "login-pass" : "password"
	// }
	raw, err := ioutil.ReadFile("/var/tmp/webstorage.json")
	if err != nil {
		return nil, fmt.Errorf("webstorage.json load error %s", err.Error())
	}

	var ac AccessInfo
	json.Unmarshal(raw, &ac)

	return &ac, nil
}

func do_get_current() {
	ac, err := load_access_info()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := ondotori.New(ac.Token, ac.Id, ac.Pass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cp := ondotori.GetCurrentParam{
		RemoteSerial: []string{},
		BaseSerial:   []string{},
	}

	res, err := client.GetCurrent(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res)
}

func do_get_latest_info() {

	ac, err := load_access_info()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := ondotori.New(ac.Token, ac.Id, ac.Pass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cp := ondotori.GetLatestDataParam{
		RemoteSerial: "5236184E",
	}

	res, err := client.GetLatestData(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res)
}

func do_get_data() {
	ac, err := load_access_info()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := ondotori.New(ac.Token, ac.Id, ac.Pass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	now := time.Now()
	from := now.Add(-1 * time.Hour)
	limit := uint16(3)

	cp := ondotori.GetDataParam{
		RemoteSerial: "5236184E",
		From:         &from,
		To:           &now,
		Number:       &limit,
	}

	res, err := client.GetData(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res)
}

func do_get_latest_rtr500() {
	ac, err := load_access_info()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := ondotori.New(ac.Token, ac.Id, ac.Pass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cp := ondotori.GetLatestDataRTR500Param{
		RemoteSerial: "52800010",
		BaseSerial:   "5858001E",
	}

	res, err := client.GetLatestDataRTR500(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res)
}

func do_get_data_rtr500() {
	ac, err := load_access_info()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := ondotori.New(ac.Token, ac.Id, ac.Pass)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	now := time.Now()
	from := now.Add(-1 * time.Hour)
	limit := uint16(3)

	cp := ondotori.GetDataRTR500Param{
		RemoteSerial: "52800010",
		BaseSerial:   "5858001E",
		From:         &from,
		To:           &now,
		Number:       &limit,
	}

	res, err := client.GetDataRTR500(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("response ", res)
}
