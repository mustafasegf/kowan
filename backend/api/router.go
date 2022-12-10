package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/go-shortener/links"
)

type Route struct {
	router *gin.Engine
}

func (s *Server) setupRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	s.router.LoadHTMLGlob("templates/*")
	s.router.Static("/static", "./static")

	linkRepo := links.NewRepo(s.db, s.rdb)
	linkSvc := links.NewService(linkRepo)
	linkCtlr := links.NewController(linkSvc)

	s.router.GET("/:url", linkCtlr.Redirect)
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", "")
	})

	s.router.POST("/api/link/create", linkCtlr.CreateLink)
}
