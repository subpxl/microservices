package main

import  (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(http.StatusOK)
	})

	r.POST("/todos", CreateTodo)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)
	r.GET("/todos", GetAllTodos)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}
