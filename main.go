package main

import (
	"log"
)

func main() {
	if err := cmdFlags(); err != nil {
		log.Fatalf("%v", err)
	}
}
