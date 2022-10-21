package main

import (
	"fmt"
	"os"
	"toolkit"
)

func main() {

	var tools toolkit.Tools

	if len(os.Args) == 1 || os.Args[1] == "-h" {
		usage(os.Args[0])
	}

	host := os.Args[1]
	result, err := tools.GetServerName(host)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func usage(name string) {
	fmt.Fprintf(os.Stdout, "Usage:\t%v hostname\n", name)
	fmt.Printf("Looking up nameservers\n")
	os.Exit(1)
}
