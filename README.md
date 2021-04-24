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

	cp := ondotori.CurrentParam{
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

## Get Latest Data

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

	cp := ondotori.LatestDataParam{
		RemoteSerial: "Device Serial here",
	}

	res, err := client.GetLatestData(cp, context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// res will be DeviceData
	fmt.Println("response ", res)
}
```

## License

The package is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).


