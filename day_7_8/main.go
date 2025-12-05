package day7

import (
	"golang-training/day_7_8/controllers"
	"golang-training/day_7_8/middlewares"
	"golang-training/day_7_8/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
)

func Day7() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.RequestFilter())
	router.Use(middlewares.Logger())

	newStore := models.NewBookStore()

	router.GET("/books", func(ctx *gin.Context) {
		controllers.GetBooksController(ctx, newStore)
	})
	router.POST("/books", func(ctx *gin.Context) {
		controllers.AddBookController(ctx, newStore)
	})
	router.PUT("/books", func(ctx *gin.Context) {
		controllers.UpdateBookController(ctx, newStore)
	})
	router.DELETE("/books", func(ctx *gin.Context) {
		controllers.DeleteBookController(ctx, newStore)
	})

	pprof.Register(router)

	router.Run("localhost:8080")
}
