# Introduction
Go package with embedded csv of Indian pinodes.


### Install
```
    go get github.com/dropdevrahul/pincodes-india/pincodes-in-go 
```


### Usage
```
package main

import (
	"github.com/dropdevrahul/pincodes-india/pincodes-in-go"
)

func main() {
	result, err := pincodes.Load()

    fmt.Println(result[110001])
}
```
