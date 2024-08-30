package services

import (
	"api/config"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "primeira api com go rodando!",
	})
}

func GetAllTasks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title FROM tasks")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		tasks = append(tasks, task)

	}

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No tasks available",
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid task id",
		})
		return
	}

	var task models.Task

	row := config.DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", idInt)

	if err := row.Scan(&task.Id, &task.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateNewTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := config.DB.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask.Id = int(id)

	c.JSON(http.StatusCreated, newTask)
}

func UpdateTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task not valid",
		})
		return
	}

	var updatedTask models.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "task not found",
		})
		return
	}

	result, err := config.DB.Exec("UPDATE tasks SET title = ? WHERE id = ?", updatedTask.Title, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})
		return
	}
	updatedTask.Id = id
	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid task id",
		})
		return
	}

	result, err := config.DB.Exec("DELETE FROM tasks WHERE id = ?", idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
