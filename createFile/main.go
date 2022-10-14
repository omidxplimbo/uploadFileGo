package main

import (
	"fmt"
	"os"
	"toolkit"
)

func main() {

	var tools toolkit.Tools

	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		fmt.Println("Please set path...")
	}

	tools.CreateFileIfNotExists(path)
}
