package main

import (
	"log"

	"skill-review/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(r)
		}
	}()

	cmd.Execute()
}
