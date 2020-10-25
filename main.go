package main

import (
	"fmt"
	"os"
	"os/user"

	"./repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Monkey :)\n", user.Username)
	fmt.Printf("Type some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
