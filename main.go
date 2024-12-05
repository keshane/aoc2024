package main

import (
	"fmt"
	"os"
)

func main() {
	problemNo := os.Args[1]
	fmt.Println(fmt.Sprintf("executing %s", problemNo))

	switch problemNo {
	case "1":
		problem1()
	case "2":
		problem2()
	}
	fmt.Println("done")
}
