# go-ondotori
Ondotori Web Storage API client  
(T&D WebStorage Service API client)

## Installation
```
go get -u github.com/k28/go-ondotori
```

## Example

### Get Current Readings

```golang
package main

import (
	"context"
	"fmt"

	"github.com/k28/go-ondotori"
)

func main() {
	client, err := ondotori.New("API Token here", "rbxx1234", "password")
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

	fmt.Println("response ", res.DeviceList)
}
```

## Get Latest Data (300 most recent readings) [TR-7wb/nw/wf, TR4]

```golang
	client, err := ondotori.New("API Token here", "rbxx1234", "password")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cp := ondotori.GetLatestDataParam{
		RemoteSerial: "Device Serial here",
	}

	res, err := client.GetLatestData(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// res will be DeviceData
	fmt.Println("response ", res)
```

## Get Data by Specific Period and Number [TR-7wb/nw/wf, TR4]

```golang
	client, err := ondotori.New("API Token here", "rbxx1234", "password")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cp := ondotori.GetDataParam{
		RemoteSerial: "Device Serial here",
	}

	res, err := client.GetData(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// res will be DeviceData
	fmt.Println("response ", res)
```

## RTR500BW

### Get Latest Data (300 most recent readings) [RTR500BW]

```golang
	cp := ondotori.GetLatestDataRTR500Param{
		RemoteSerial: "remote unit serial here",
		BaseSerial:   "base unit serial here",
	}

	res, err := client.GetLatestDataRTR500(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// res will be DeviceData
	fmt.Println("response ", res)
```

## License

The package is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).


