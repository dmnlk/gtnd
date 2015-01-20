package main

import (
	"github.com/dmnlk/gtnd"
	"fmt"
	"github.com/k0kubun/pp"
)

func main() {
	param := gtnd.SearchParam{"golang"}
	result, err := gtnd.Search(param)
	if err != nil {
		fmt.Println(err)
	}

	pp.Print(result)
}
