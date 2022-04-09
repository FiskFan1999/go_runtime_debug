package main

import (
	"log"
	"runtime/debug"
)

func main() {
	log.Println(debug.ReadBuildInfo())
}
