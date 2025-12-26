package handlers

import (
	"taskmanager/db"
	"taskmanager/models"
	"time"
)

func AddTask(title, description string, deadline time.Time) error {
	task := models.Task{
		Title:       title,
		Description: description,
		Status:      "pendente",
		Deadline:    deadline,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return db.DB.Create(&task).Error
}

func ListTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Find(&tasks).Error
	return tasks, err
}

func UpdateTaskStatus(id uint, status string) error {
	return db.DB.Model(&models.Task{}).Where("id = ?", id).Update("status", status).Error
}

func DeleteTask(id uint) error {
	return db.DB.Delete(&models.Task{}, id).Error
}
