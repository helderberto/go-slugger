# 🔗 GO Slugger

> This package generate slug from unicode string, URL-friendly slugify.

## How to Install

```bash
$ go get -u github.com/helderburato/go-slugger
```

## How to Use

```go
import (
	"fmt"
	"github.com/helderburato/go-slugger"
)

func main() {
	text := slugger(Params{text: "Lorem && ** Ipsum is simply dummy"})
	fmt.Println(text) // => lorem-ipsum-is-simply-dummy
}
```
