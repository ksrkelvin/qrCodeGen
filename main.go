package main

import (
	"fmt"
	"log"
	"qrCodeGen/internal/api"
	"qrCodeGen/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	diinoAPI, err := api.New()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	app := gin.Default()

	err = routes.Router(app, diinoAPI)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Fatal(app.Run(":300"))

}
