package main

import (
	"github.com/cesc1802/auth-service/cmd"
	"log"
)

// @title Auth Service API
// @version 1.0
// @description This is Auth Service API.

// @contact.name Cesc Nguyen
// @contact.email thuocnv1802@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
