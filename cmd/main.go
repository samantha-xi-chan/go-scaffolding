package main

import (
	"fmt"
	"go-scaffolding/util/util_etcd"
	_ "go.uber.org/automaxprocs"
	"log"
	"runtime"
)

func printVersionInfo() string {
	return fmt.Sprintf(" Version:   %s\n BuildTime: %s\n GitBranch: %s\n GitCommit: %s",
		Version, BuildTime, GitBranch, GitCommit)
}

func init() {
	log.Println("init start")
	defer log.Println("init end")

	versionInfo := printVersionInfo()
	log.Printf("printVersionInfo: \n%s\n", versionInfo)
	log.Println("Current GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}

func main() {
	log.Println("main start")
	defer log.Println("main end")

	util_etcd.Test()
}
