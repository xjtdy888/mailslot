## GoLang 实现Windows下邮件槽封装

## Usage 使用方法
```
go get github.com/xjtdy888/examples/...

server.exe
client.exe
```
## 说明
- mailSlot.New 实现标准 io.Reader
- mailSlot.Open 实现标准 io.Writer

## Server
```
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/xjtdy888/mailslot"
)

const MAIL_SLOT_NAME = `\\.\mailslot\server001`

func main() {
	ms, err := mailslot.New(MAIL_SLOT_NAME, 0, mailslot.MAILSLOT_WAIT_FOREVER)
	defer ms.Close()

	if err != nil {
		panic(err)
	}

	sz, err := io.Copy(os.Stdout, ms)

	fmt.Println(sz, err)

}
```

## Client

```
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/xjtdy888/mailslot"
)

const MAIL_SLOT_NAME = `\\.\mailslot\server001`

func main() {
	ms, err := mailslot.Open(MAIL_SLOT_NAME)
	defer ms.Close()

	if err != nil {
		panic(err)
	}

	n, err := io.Copy(ms, os.Stdin)

	fmt.Println(n , err)

}
```
