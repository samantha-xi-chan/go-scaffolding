package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-scaffolding/config"
	"go-scaffolding/internal/app01/handler"
	"go-scaffolding/internal/app01/model"
	"go-scaffolding/pkg/db"
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
	ctx := context.TODO()
	log.Printf("ctx: %#v", ctx)

	log.Println("main start")
	defer log.Println("main end")

	// test util_etcd
	//go func() {
	//	util_etcd.StartReadTest(ctx)
	//}()
	//
	//time.Sleep(time.Second)
	//
	//go func() {
	//	for true {
	//		util_etcd.WriteTest(ctx)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//
	//// test util_etcd_v2
	//for true {
	//	util_etcd_v2.Test()
	//	time.Sleep(time.Second)
	//}

	// db init
	dbConn, err := db.Connect(config.MySQLDSN)
	if err != nil {
		log.Fatalf("db.Connect(config.MySQLDSN): %s", err)
	}
	if err := dbConn.Migrator().DropTable(
		&model.User{},
	); err != nil {
		log.Fatalf("failed to DropTable: %v", err)
	}
	if err := dbConn.AutoMigrate(
		&model.User{},
	); err != nil {
		log.Fatalf("failed to auto migrate database: %v", err)
	}

	r := gin.Default()
	h := handler.NewHandler(dbConn)
	h.RegisterRoutes(r)
	r.Run(":8080")

	log.Println("waiting select {}")
	select {}
}
