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

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, task := range tasks {
		resTask := model.TaskResponse{
				ID 	      : task.ID,
				Title     : task.Title,
				CreatedAt : task.CreatedAt,
				UpdatedAT : task.UpdatedAT,
		}
		resTasks = append(resTasks, resTask)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error){
	task:= model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	reTask := model.TaskResponse{
			ID 	      : task.ID,
			Title     : task.Title,
			CreatedAt : task.CreatedAt,
			UpdatedAT : task.UpdatedAT,
	}
	return reTask, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, nil
	}
	reTask := model.TaskResponse{
			ID 	      : task.ID,
			Title     : task.Title,
			CreatedAt : task.CreatedAt,
			UpdatedAT : task.UpdatedAT,
	}
	return reTask, nil
}

// TODO: これtaskを値渡しして，UpdateTaskでは参照渡しにすることで，変更前と変更後の情報をどっちも保持できる？
func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, nil
	}
	reTask := model.TaskResponse{
			ID 	      : task.ID,
			Title     : task.Title,
			CreatedAt : task.CreatedAt,
			UpdatedAT : task.UpdatedAT,
	}
	return reTask, nil
}

func (tu* taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}