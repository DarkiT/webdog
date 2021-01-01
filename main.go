package main

import (
	"edboffical/webdog/config"
	"edboffical/webdog/router"
	"log"
	"net/http"
	"os"
)

func main() {
	dr := router.InitDogRouter()
	cfg := config.ReadCfg()
	// start config monitor
	go config.InitMonitor(dr.RegisterRouter)
	// log pid
	log.Printf("server process pid:%d", os.Getpid())
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, dr.GetRouter()).Error())
}
