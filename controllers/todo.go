package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	guuid "github.com/google/uuid"
	models "github.com/rizperdana/prototype-golang-rest/models"
)

var dbModel *pg.DB

func InitiateDB(db *pg.DB) {
	dbModel = db
}

func GetListTodos(c *gin.Context) {
	var todos = []*models.Todo{}
	err := dbModel.Model(&todos).Select()
	if err != nil {
		log.Printf("Error while getting all todos! Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todos retrieved",
		"data":    todos,
	})
}

func CreateTodo(c *gin.Context) {
	var todo = &models.Todo{}
	c.BindJSON(&todo)

	title := todo.Title
	body := todo.Body
	id := guuid.New().String()

	insertError := dbModel.Insert(&models.Todo{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
}

func GetDetailTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &models.Todo{ID: todoId}
	err := dbModel.Select(todo)

	if err != nil {
		log.Printf("Error unable to found todo! Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo received!",
		"data":    todo,
	})
}
