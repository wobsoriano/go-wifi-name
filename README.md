# go-wifi-name

Get current wifi name in Go.

## Install

```bash
$ go get github.com/wobsoriano/go-wifi-name
```

## Usage

```go
package main

import (
	"fmt"
	wifiname "github.com/wobsoriano/go-wifi-name"
)

func main() {
	fmt.Println("wifi name ", wifiname.WifiName())
}
```

## License

MIT
