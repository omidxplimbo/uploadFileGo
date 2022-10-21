package main

import (
	"fmt"
	"os"
	"toolkit"
)

func main() {

	var tools toolkit.Tools

	if len(os.Args) != 3 {
		usage(os.Args[0])
	}

	result := tools.Base64Tools(os.Args[1], os.Args[2])

	fmt.Println(result)
}

func usage(name string) {
	fmt.Fprintf(os.Stdout, "Using: %s encode DATA\n", name)
	fmt.Fprintf(os.Stdout, "OR %s decode DATA\n", name)
	os.Exit(1)
}
