package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	rows, err := DB.Query(context.Background(), "select * from todo;")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Todo); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func CreateTodo(c *gin.Context) {
	if c.Request.Method == http.MethodPost {

		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo.ID = rand.Intn(55)
		_, err := DB.Exec(context.Background(), "INSERT INTO todo (Todo) values ($1)", todo.Todo)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{
			"todos": "created successfully todo",
		})
	}
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Request.Method == http.MethodPut {

		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := DB.Exec(context.Background(), "UPDATE todo SET todo = $1 WHERE id = $2", todo.Todo, id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{
			"todos": "updated successfully todo",
		})
	}

}

func DeleteTodo(c *gin.Context) {
	if c.Request.Method == http.MethodDelete {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err = DB.Exec(context.Background(), "DELETE FROM todo WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "deleted successfully todo"})
	}
}
