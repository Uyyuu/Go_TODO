package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error) 
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) 
	DeleteTask(userId uint, taskId uint) error 
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

// TODO
func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	return nil, nil
}

// TODO
func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error){
	return nil, nil
}

// TODO: なんで引数にuser_idないのか？
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	return nil, nil
}

// TODO
func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	return nil, nil
}

// TODO
func (tu* taskUsecase) DeleteTask(userId uint, taskId uint) error {
	return nil
}