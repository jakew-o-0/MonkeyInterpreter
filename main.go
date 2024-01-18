package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jakew-o-0/MonkeyInterpreter/repl"
)



func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("hello %s! This is the monkey programming language!\n", user.Username)
    fmt.Printf("you are now in the REPL:\n")
    repl.Start(os.Stdin, os.Stdout)
}
