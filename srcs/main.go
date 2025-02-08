package main

import (
	"fmt"
	"os/user"
	"shell/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", user.Username)
	repl.Start()
}
