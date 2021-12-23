# TC Identity   

TC identity generator and validate

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

Validate Tc identity

```bash
package main

import (
	"fmt"
	"github.com/yisleyen/tckimlik"
	"log"
)

func main()  {
	valid, err := tckimlik.Validate("11111111111")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(valid)
}
```