package main

import (
	"fmt"
	"monkeyLang/repl"
	"os"
	"os/user"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s this is monkey programming language!\n", usr.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
