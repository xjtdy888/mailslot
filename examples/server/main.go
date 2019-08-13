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
