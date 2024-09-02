## tdage

`tdage` is a Go package that allows you to calculate the creation date of Telegram accounts based on a dataset of IDs. This package provides functionality to determine whether an account is older or newer than a given ID, as well as to obtain an approximate creation date.

### Usage

```go
package main

import (
	"fmt"

	"gopkg.in/tdage.v1"
)

func main() {
	pool := tdage.NewPool()
	userId := int64(900323135)
	result := pool.GetDate(userId)
    	date := fmt.Sprintf("%02d/%d", result.Date.Month(), r.Date.Year())
	fmt.Printf("%d: %s %s", userId, result.Status, date)
}
```
