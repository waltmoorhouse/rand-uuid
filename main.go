package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func main() {
	var useV6 bool
	var useV7 bool
	flag.BoolVar(&useV6, "v6", false, "Generate UUIDv6 instead of UUIDv4")
	flag.BoolVar(&useV7, "v7", false, "Generate UUIDv7 instead of UUIDv4")
	flag.Parse()

	var uuidStr string
	var err error

	if useV7 {
		var uuid7 uuid.UUID
		uuid7, err = uuid.NewV7()
		if err == nil {
			uuidStr = uuid7.String()
		}
	} else if useV6 {
		var uuid6 uuid.UUID
		uuid6, err = uuid.NewV6()
		if err == nil {
			uuidStr = uuid6.String()
		}
	} else {
		uuidStr = uuid.NewString()
	}

	if err != nil {
		log.Fatalf("Error generating UUID: %v", err)
	}

	fmt.Println(uuidStr)
}
