package main

import (
	"github.com/evansmwendwa/uxp/db"
	"github.com/evansmwendwa/uxp/router"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	session := db.Session()
	defer session.Close()

	e := router.Echo()

	e.Logger.Fatal(e.Start(":8080"))
}
