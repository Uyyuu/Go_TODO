package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewtaskRepository (db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

// TODO: 元の回答とチェック
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// TODO: 元の回答とチェック
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Where("user_id=? AND Id=?", userId, taskId).First(task).Error; err != nil {
		return err
	}
	return nil
}

// TODO: 元の回答とチェック
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// TODO: 元の回答と結構違うぽい．チェックする
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Where("user_id=? AND id=?", userId, taskId).Save(task).Error; err != nil {
		return err
	}
	return nil
}

// TODO: もし消せなかった場合は？
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	if err := tr.db.Where("user_id=? AND id=?", userId, taskId).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}