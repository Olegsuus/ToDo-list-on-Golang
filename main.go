package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db := InitDB("tasks.db")
	defer db.Close()
	CreateTable(db)

	router := gin.Default()
	router.LoadHTMLGlob("*.html")

	router.GET("/", func(c *gin.Context) {
		tasks, err := GetAllTasks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": tasks,
		})
	})

	router.GET("/api/tasks", func(c *gin.Context) {
		tasks, err := GetAllTasks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})

	router.POST("/tasks", func(c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")
		category_IDStr := c.PostForm("category_id")
		category_ID, _ := strconv.Atoi(category_IDStr)

		err := CreateTask(db, title, description, category_ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/tasks/update", func(c *gin.Context) {
		idStr := c.PostForm("id")
		id, _ := strconv.Atoi(idStr)
		title := c.PostForm("title")
		description := c.PostForm("description")
		categoryIDStr := c.PostForm("category_id")
		categoryID, _ := strconv.Atoi(categoryIDStr)
		completedStr := c.PostForm("completed")
		completed := completedStr == "on"

		if err := UpdateTask(db, id, title, description, categoryID, completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/tasks/delete", func(c *gin.Context) {
		idStr := c.PostForm("id")
		id, _ := strconv.Atoi(idStr)

		if err := DeleteTask(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	router.Run(":8080")
}
