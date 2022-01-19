package models

import (
	"backend-task-eventori/database"
	"time"
)

type Schedules struct {
	Id int `json:"id"`
	Model_id int `json:"model_id"`
	Schedule_date time.Time `json:"schedule_date"`
}

type Models struct {
	Id          int    `json:"id"`
	Model_name  string `json:"model_name"`
	Image_url   string `json:"image_url"`
	Description string `json:"description"`
	Schedules	[]Schedules `json:"schedules"`
}

type ModelsList []Models

type Response struct {
	Status 	string `json:"status"`
	Message string `json:"message"`
}

func GetAllModels() (ModelsList, error) {
	var list ModelsList
	
	conn := database.Connect()

	sqlStatement := `SELECT 
	mc.id, 
    mc.model_name, 
    mc.image_url, 
    mc.description,
    ms.id,
    ms.model_id, 
    ms.schedule_date,
	(case
		when ms.deleted_at is null then false
        else true
	end) as is_deleted
FROM model_catalogues as mc 
JOIN model_schedules as ms
ON mc.id = ms.model_id
WHERE mc.deleted_at IS NULL`

	rows, err := conn.Query(sqlStatement)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	ids:= make(map[int]bool)
	for rows.Next(){
		model := Models{}
		schedule := Schedules{}
		var is_deleted bool
		err = rows.Scan(&model.Id, &model.Model_name, &model.Image_url, 
			&model.Description, &schedule.Id, &schedule.Model_id, &schedule.Schedule_date, &is_deleted)
		if err != nil {
			return list, err
		}
		if !ids[model.Id] {
			ids[model.Id] = true
			if !is_deleted {
				model.Schedules = append(model.Schedules, schedule)
			}
			list = append(list, model)
		} else {
			for i, _ := range list {
				if list[i].Id == model.Id{
					list[i].Schedules = append(list[i].Schedules, schedule)
					break
				}
			}
		}
	}

	return list, nil
}

func CreateModel(model_name string, image_url string, description string) (Response, error){
	var res Response
	conn := database.Connect()

	sqlStatement := "INSERT INTO model_catalogues (model_name, image_url, description) VALUES(?, ?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	_, err = stmt.Exec(model_name, image_url, description)
	if err != nil {
		return res, err
	}

	res.Status = "success"
	res.Message = "Model created successfully"

	return res, nil
}

func GetSchedulesByModelId(model_id int) ([]Schedules, error) {
	var list []Schedules
	conn := database.Connect()

	sqlStatement := "SELECT id, model_id, schedule_date FROM model_schedules WHERE model_id = ? AND deleted_at IS NULL"

	rows, err := conn.Query(sqlStatement, model_id)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next(){
		schedule := Schedules{}
		err = rows.Scan(&schedule.Id, &schedule.Model_id, &schedule.Schedule_date)
		if err != nil {
			return list, err
		}
		list = append(list, schedule)
	}
	return list, nil
}

func CreateSchedule(model_id int, schedule_date string)(Response, error) {
	var res Response
	conn := database.Connect()

	sqlStatement := "INSERT INTO model_schedules (model_id, schedule_date) VALUES (?, ?)"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	_, err = stmt.Exec(model_id, schedule_date)
	if err != nil {
		return res, err
	}

	res.Status = "success"
	res.Message = "Schedule created successfully"

	return res, nil
}

func GetModelById(id int)(Models, error) {
	var model Models
	conn := database.Connect()

	sqlStatement := "SELECT id, model_name, image_url, description FROM model_catalogues WHERE id = ? AND deleted_at IS NULL"

	err := conn.QueryRow(sqlStatement, id).Scan(&model.Id, &model.Model_name, &model.Image_url, &model.Description)
	if err != nil {
		return model, err
	}
	
	sqlStatement = "SELECT id, model_id, schedule_date FROM model_schedules WHERE model_id = ?" 

	rows, err := conn.Query(sqlStatement, id)
	if err != nil {
		return model, err
	}

	for rows.Next(){
		schedule := Schedules{}
		err = rows.Scan(&schedule.Id, &schedule.Model_id, &schedule.Schedule_date)
		if err != nil {
			return model, err
		}
		model.Schedules = append(model.Schedules, schedule)
	}

	return model, nil
}

func UpdateModelById(id int, model_name string, image_url string, description string)(Response, error){
	var res Response

	conn := database.Connect()

	sqlStatement := "UPDATE model_catalogues SET model_name = ?, image_url = ?, description = ? WHERE id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	_, err = stmt.Exec(model_name, image_url, description, id)
	if err != nil {
		return res, err
	}

	res.Status = "success"
	res.Message = "Model updated successfully"

	return res, nil
}

func UpdateScheduleById(id int, model_id int, schedule_date string)(Response, error){
	var res Response
	conn := database.Connect()
	
	sqlStatement := "UPDATE model_schedules SET model_id = ?, schedule_date = ? WHERE id = ?" 

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	_, err = stmt.Exec(model_id, schedule_date, id)
	if err != nil {
		return res, err
	}

	res.Status = "success"
	res.Message = "Schedule updated successfully"

	return res, nil
}

func DeleteScheduleById(id int)(Response, error){
	var res Response
	conn := database.Connect()

	sqlStatement := "UPDATE model_schedules SET deleted_at=? WHERE id=?" 

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		return res, err
	}

	res.Status = "success"
	res.Message = "Schedule deleted successfully"

	return res, nil
}

func DeleteModelById(id int)(Response, error){
	var res Response

	conn := database.Connect()

	sqlStatement := "UPDATE model_catalogues SET deleted_at = ? WHERE id = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	_, err = stmt.Exec(time.Now(), id)
	if err != nil {
		return res, err
	}

	res.Status = "success"
	res.Message = "Model deleted successfully"

	return res, nil
}

