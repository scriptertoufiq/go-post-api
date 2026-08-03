package main

import (
	"fmt"
	"go-tweets/internal/config"
	"go-tweets/pkg/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	postHandlerAlias "go-tweets/internal/handler/post"
	userHandlerAlias "go-tweets/internal/handler/user"
	postRepoAlias "go-tweets/internal/repository/post"
	userRepoAlias "go-tweets/internal/repository/user"
	postServiceAlias "go-tweets/internal/service/post"
	userServiceAlias "go-tweets/internal/service/user"
)

func main() {
	r := gin.Default()
	validate := validator.New()
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
	postRepo := postRepoAlias.NewRepository(db)

	userService := userServiceAlias.NewService(cfg, userRepo)
	postService := postServiceAlias.NewPostService(cfg, postRepo)

	userHandler := userHandlerAlias.NewHandler(r, validate, userService)
	userHandler.RouteList()

	postHandler := postHandlerAlias.NewHandler(r, validate, postService)
	postHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)

	r.Run(server)
}
