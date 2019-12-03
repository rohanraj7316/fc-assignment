package main

import (
	"freecharge/router"
	"log"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var err error

	err = router.New(8080)
	if err != nil {
		panic("error in router connection")
	}
}
