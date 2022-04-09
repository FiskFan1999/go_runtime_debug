package main

import (
	"log"
	"runtime/debug"
)

func main() {
	deb, _ := debug.ReadBuildInfo()
	log.Println(deb)
	log.Println("second commit")
}
