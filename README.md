# TC Identity   

TC identity generator

## Run Your Computer

```bash
go get -u github.com/yisleyen/tckimlik
```

Generate Tc identity

```bash
package main

import (
	"fmt"
	"github.com/yisleyen/tckimlik"
)

func main()  {
	fmt.Println(tckimlik.Generate())
}
```