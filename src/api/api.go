// Package main that runs tha main process for the api executable
package main

import (
	"log"

	_ "notifier/src/api/routes"
)

func main() {
	// Just a typical recover mechanism here.
	defer func() {
		if recover := recover(); recover != nil {
			log.Println(recover)
			panic(recover)
		}
	}()
}
