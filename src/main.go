package main

import (
	"www/src/router"
)

func main() {
	rt := router.Init()

	rt.Logger.Fatal(rt.Start(":8080"))
}
