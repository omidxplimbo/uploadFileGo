package main

import (
	"fmt"
	"toolkit"
)

func main() {

	var tools toolkit.Tools

	result, err := tools.SlugifyString("    234%$^%$%^asdasdasdasdasdSDASDASDASDAS 345")

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(result)
	}

}
