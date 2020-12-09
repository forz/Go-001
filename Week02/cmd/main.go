// Package main ...
package main

import (
	"Go-000/Week02/api"
	"fmt"
)

func main() {
	isSame, err := api.SameUserName("0001", "0002")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	if isSame {
		fmt.Println("0001 and 0002 have same name!")
	}
}
