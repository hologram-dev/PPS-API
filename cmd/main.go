package main

import (
	"log"

	"gorm-template/bootstrap"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	route "gorm-template/api/route"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()
	router.Use(cors.Default())

	route.Setup(env, timeout, router)
	log.Fatal(router.Run(env.ServerAddress))
}
