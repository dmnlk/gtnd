package main

import (
	"github.com/dmnlk/gtnd"
	"fmt"
	"github.com/k0kubun/pp"
)

func main() {
	result, err := gtnd.Search("golang")
	if err != nil {
		fmt.Println(err)
	}

	pp.Print(result)
}
