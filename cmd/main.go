package main

import (
	"fmt"
	"go-tweets/internal/config"
	"go-tweets/pkg/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	userHandlerAlias "go-tweets/internal/handler/user"
	userRepoAlias "go-tweets/internal/repository/user"
	userServiceAlias "go-tweets/internal/service/user"
)

func main() {
	r := gin.Default()
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's working",
		})
	})

	userRepo := userRepoAlias.NewRepository(db)
	userService := userServiceAlias.NewService(cfg, userRepo)
	handler := userHandlerAlias.NewHandler(r, userService)
	handler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)

	r.Run(server)
}
