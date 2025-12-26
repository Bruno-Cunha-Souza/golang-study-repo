package handlers

import (
	"errors"
	"taskmanager/db"
	"taskmanager/models"
	"time"
)

var (
	ErrTaskNotFound  = errors.New("tarefa não encontrada")
	ErrInvalidStatus = errors.New("status inválido. Use: pendente, em_progresso, concluida")
)

var validStatuses = map[string]bool{
	"pendente":     true,
	"em_progresso": true,
	"concluida":    true,
}

func AddTask(title, description string, deadline time.Time) error {
	task := models.Task{
		Title:       title,
		Description: description,
		Status:      "pendente",
		Deadline:    deadline,
	}
	return db.DB.Create(&task).Error
}

func ListTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Find(&tasks).Error
	return tasks, err
}

func UpdateTaskStatus(id uint, status string) error {
	if !validStatuses[status] {
		return ErrInvalidStatus
	}

	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		return ErrTaskNotFound
	}

	return db.DB.Model(&task).Update("status", status).Error
}

func DeleteTask(id uint) error {
	var task models.Task
	result := db.DB.First(&task, id)
	if result.Error != nil {
		return ErrTaskNotFound
	}

	return db.DB.Delete(&task).Error
}
