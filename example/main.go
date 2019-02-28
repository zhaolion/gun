package main

import (
	"fmt"

	"github.com/zhaolion/gun/util/consoleutil"
)

func main() {
	fmt.Print("please confirm: ")
	c := consoleutil.AskForConfirmation()
	fmt.Printf("response: %v", c)
}
