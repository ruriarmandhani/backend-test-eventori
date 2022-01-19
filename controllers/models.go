package controllers

import (
	"backend-task-eventori/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllModels(c echo.Context) error {
	res, err := models.GetAllModels()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func CreateModel(c echo.Context) error {
	model_name := c.FormValue("model_name")
	image_url := c.FormValue("image_url")
	description := c.FormValue("description")
	res, err := models.CreateModel(model_name, image_url, description)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func GetSchedulesByModelId(c echo.Context) error {
	model_id := c.Param("id")
	conv_model_id, err := strconv.Atoi(model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.GetSchedulesByModelId(conv_model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func CreateSchedule(c echo.Context) error {
	model_id := c.FormValue("model_id")
	schedule_date := c.FormValue("schedule_date")
	conv_model_id, err := strconv.Atoi(model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.CreateSchedule(conv_model_id, schedule_date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func GetModelById(c echo.Context) error {
	model_id := c.Param("model_id")
	conv_model_id, err := strconv.Atoi(model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.GetModelById(conv_model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateModelById(c echo.Context) error {
	model_id := c.Param("model_id")
	model_name := c.FormValue("model_name")
	image_url := c.FormValue("image_url")
	description := c.FormValue("description")

	conv_model_id, err := strconv.Atoi(model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.UpdateModelById(conv_model_id, model_name, image_url, description)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateScheduleById(c echo.Context) error {
	schedule_id := c.Param("id")
	model_id := c.FormValue("model_id")
	schedule_date := c.FormValue("schedule_date")
	conv_schedule_id, err := strconv.Atoi(schedule_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	conv_model_id, err := strconv.Atoi(model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.UpdateScheduleById(conv_schedule_id, conv_model_id, schedule_date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteScheduleById(c echo.Context) error {
	schedule_id := c.Param("id")
	conv_schedule_id, err := strconv.Atoi(schedule_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.DeleteScheduleById(conv_schedule_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteModelById(c echo.Context) error {
	model_id := c.Param("model_id")
	conv_model_id, err := strconv.Atoi(model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	res, err := models.DeleteModelById(conv_model_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message":err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}