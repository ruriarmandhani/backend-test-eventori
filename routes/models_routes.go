package routes

import (
	"backend-task-eventori/controllers"

	"github.com/labstack/echo/v4"
)

func ModelsRoutes(g *echo.Group) {
	// Get list of models
	g.GET("/list", controllers.GetAllModels)

	// Create new model
	g.POST("/create", controllers.CreateModel)

	// Get list of schedules by model id
	g.GET("/schedules/:id", controllers.GetSchedulesByModelId)

	// Create new schedule for model
	g.POST("/schedules/create", controllers.CreateSchedule)

	// Get model data by model id
	g.GET("/:model_id", controllers.GetModelById)

	// Update model data by model id
	g.PATCH("/update/:model_id", controllers.UpdateModelById)

	// Update schedule data by schedule id
	g.PATCH("/schedules/:id", controllers.UpdateScheduleById)

	// Delete schedule data by schedule id
	g.DELETE("/schedules/:id", controllers.DeleteScheduleById)

	// Delete model data by model id
	g.DELETE("/:model_id", controllers.DeleteModelById)
}