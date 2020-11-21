## Kittenclient-go

Kittenclient is a client to [kittenhouse](https://github.com/VKCOM/kittenhouse) written in go. He makes requests to kittenhouse easier. 

## Example Usage
```go
package main

import "github.com/EpicStep/kittenclient-go"

func main() {
	config := kittenclient.ClientConfig{
		Addr: "localhost:13338",
		UDP:  false,
	}

	klogger := kittenclient.NewLogger(&config)

	err := klogger.Log("internal_logs_buffer(time,server,port)", "('2020-11-20 14:40:00','%s', 8080)", "myfavoriteserver")
	if err != nil {
		panic(err)
	}
}
```

## Features
Client support:
* HTTP transport
* UDP transport

## TODO
* RowBinary support

