package route

import (
	"strconv"

	"github.com/amha-mersha/go_taskmanager/controllers"
	"github.com/gin-gonic/gin"
)

func Run(port int) {
	router := gin.Default()

	router.GET("/api/v1/tasks", controllers.GetTasks)
	router.GET("/api/v1/tasks/:id", controllers.GetTaskByID)
	router.POST("/api/v1/tasks", controllers.PostTask)
	router.PUT("/api/v1/tasks/:id", controllers.UpdateTask)
	router.DELETE("/api/v1/tasks/:id", controllers.DeleteTask)

	router.Run("localhost:" + strconv.Itoa(port))
}
