package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "RestApI/docs"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

func main() {
	client := gin.Default()

	client.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Invalid route"})
	})

	client.GET("/ping", ping)
	client.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	client.GET("/tasks", getTasks)
	client.POST("/tasks", addTask)
	client.GET("/tasks/:id", getTasksById)
	client.DELETE("/tasks/:id", delTasksById)
	client.PUT("/tasks/:id", updateByID)
	err := client.Run(":8080")
	if err != nil {
		return
	}
}

// @Summary      Health check
// @Description  Responds with pong to show the server is running
// @Tags         health
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func ping(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}

// @Summary      Get all tasks
// @Description  Returns a list of all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   Task
// @Router       /tasks [get]
func getTasks(context *gin.Context) {
	context.JSON(200, tasks)
}

// @Summary      Add a new task
// @Description  Adds a new task to the list
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      Task  true  "Task to add"
// @Success      201   {object}  Task
// @Failure      400   {object}  map[string]string
// @Router       /tasks [post]
func addTask(context *gin.Context) {
	var task Task
	if err := context.BindJSON(&task); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, t := range tasks {
		if t.ID == task.ID {
			context.JSON(400, gin.H{"error": "Task ID already exists"})
			return
		}
	}

	tasks = append(tasks, task)
	context.JSON(201, task)
}

// @Summary      Get task by ID
// @Description  Returns a single task by its ID
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  Task
// @Failure      404  {object}  map[string]string
// @Router       /tasks/{id} [get]
func getTasksById(context *gin.Context) {
	idParam := context.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(404, gin.H{"error": "Invalid task ID"})
	}

	for _, task := range tasks {
		if task.ID == id {
			context.JSON(200, task)
			return
		}
	}

	context.JSON(404, gin.H{"error": "Invalid task ID"})
}

// @Summary      Delete task by ID
// @Description  Deletes a task by its ID
// @Tags         tasks
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /tasks/{id} [delete]
func delTasksById(context *gin.Context) {
	idParam := context.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(404, gin.H{"error": "Invalid task ID"})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			context.JSON(200, gin.H{"message": "Task deleted successfully"})
			log.Printf("Current tasks: %+v", tasks)
			return
		}
	}

	context.JSON(404, gin.H{"error": "Invalid task ID"})
	log.Printf("Current tasks: %+v", tasks)
}

// @Summary      Update task by ID
// @Description  Updates a task by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int   true  "Task ID"
// @Param        task  body      Task  true  "Updated task"
// @Success      200   {object}  Task
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /tasks/{id} [put]
func updateByID(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(404, gin.H{"error": "Invalid task ID"})
		return
	}

	var updated Task
	if err := context.BindJSON(&updated); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updated.ID = id

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updated
			context.JSON(200, updated)
			log.Printf("Current tasks: %+v", tasks)
			return
		}
	}
	context.JSON(404, gin.H{"error": "Task not found"})
	log.Printf("Current tasks: %+v", tasks)
}
