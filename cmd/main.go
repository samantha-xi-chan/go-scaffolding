package main

import (
	"context"
	"fmt"
	"go-scaffolding/util/util_etcd"
	"go-scaffolding/util/util_etcd_v2"
	_ "go.uber.org/automaxprocs"
	"log"
	"runtime"
	"time"
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
	ctx := context.TODO()
	log.Printf("ctx: %#v", ctx)

	log.Println("main start")
	defer log.Println("main end")

	// test util_etcd
	go func() {
		util_etcd.StartReadTest(ctx)
	}()

	time.Sleep(time.Second)

	go func() {
		for true {
			util_etcd.WriteTest(ctx)
			time.Sleep(time.Second)
		}
	}()

	// test util_etcd_v2
	for true {
		util_etcd_v2.Test()
		time.Sleep(time.Second)
	}

	select {}
}
