/*
Copyright © 2022 purush7

*/
package main

import (
	"log"

	"github.com/purush7/project/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
