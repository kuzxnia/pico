package main

import (
	"fmt"
	"os"
	"os/user"
	"pico/executor"
	"pico/repl"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		executor.ExecuteFile(args[0], os.Stdout)
		os.Exit(0)
	}

	fmt.Printf("%v", os.Args[1:])
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Pico programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
