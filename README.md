# TC Kimlik Numarası    

TC Kimlik numarası oluşturucu



## Bilgisayarınızda Çalıştırın

```bash
go get -u github.com/yisleyen/tckimlik
```

TC Kimlik numarası üretin

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