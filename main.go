package main

import (
	"log"

	"skill-review/cmd"
	_ "skill-review/docs"
)

// @title SkillReviewApp
// @version 1.0
// @description Rest and Grpc endpoints.
// @BasePath /
func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(r)
		}
	}()

	cmd.Execute()
}
