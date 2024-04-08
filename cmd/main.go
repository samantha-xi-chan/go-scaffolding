package main

import (
	_ "go.uber.org/automaxprocs"
	"log"
)

func init() {
	log.Println("init")
}

func main() {
	log.Println("main")
}
