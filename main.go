package main

import (
	"design/command"
	"fmt"
)

func main() {
	err := command.Do()
	if err != nil {
		fmt.Println(err)
	}
}
