package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"go-tweets/internal/config"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()

	if err != nil {
	log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)


	r.Run(server) 
}