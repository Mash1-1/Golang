package usecases

import (
	"task_manager_ca/Domain"
)

type TaskUseCase struct {
	TaskRepo Domain.TaskRepository
}

type TaskUseCaseI interface {
	GetElementByID(id string) (Domain.Task, error)
	GetAllElements() ([]Domain.Task, error)
	CreateTask(new_task Domain.Task) error
	UpdateTask(id string, new_task Domain.Task) error
	DeleteTask(id string) error
}

func NewTaskUseCase(tr Domain.TaskRepository) TaskUseCase{
	return TaskUseCase{
		TaskRepo: tr,
	}
}

func (tc *TaskUseCase) GetElementByID(id string) (Domain.Task, error){
	return tc.TaskRepo.GetTaskByID(id)
}

func (tc *TaskUseCase) GetAllElements() ([]Domain.Task, error){
	return tc.TaskRepo.GetAllElements()
}

func (tc *TaskUseCase) CreateTask(new_task Domain.Task) error {
	return tc.TaskRepo.CreateTask(new_task)
}

func (tc *TaskUseCase) UpdateTask(id string, new_task Domain.Task) error {
	return tc.TaskRepo.UpdateTaskByID(id, new_task)
}

func (tc *TaskUseCase) DeleteTask(id string) error {
	return tc.TaskRepo.DeleteTask(id)
}