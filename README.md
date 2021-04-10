# go-ondotori
Ondotori Web Storage API client  
(T&D WebStorage Service API client)

## Installation
```
go get -u github.com/k28/go-ondotori
```

## Example

### Get Current Readings
```
package main

import (
	"context"
	"fmt"

	"github.com/k28/go-ondotori"
)

func main() {
	client, err := ondotori.New("xxxxxx", "rbxx1234", "password")
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
```

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).


