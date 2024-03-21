package main

import (
	"api/pkg/config"
	"api/pkg/di"
	"log"
)

func main() {
	c,configerr:=config.LoadConfig()
	if configerr != nil{
	log.Fatal("cannot load config",configerr)
	}
	server,dier:=di.InitializeApi(c)
	if dier != nil{
		log.Fatal("cannot intitalize server",dier)
	}
	server.Start()
}