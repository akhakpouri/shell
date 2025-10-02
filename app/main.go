package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	command := bufio.NewScanner(os.Stdin)

	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Error reading input: ", err)
	// 	os.Exit(1)
	// }

	for command.Scan() {
		line := command.Text()
		if line == "exit" {
			break
		}
		fmt.Println(command.Text() + ": command not found")
	}
}
