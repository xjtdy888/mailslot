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
