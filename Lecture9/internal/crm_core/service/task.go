package service

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/entity"
)

func (s *Service) GetTasks(ctx *gin.Context) (*[]entity.Task, error) {
	tasks, err := s.Repo.GetTasks(ctx)

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Service) GetTask(ctx *gin.Context, id string) (*entity.Task, error) {
	task, err := s.Repo.GetTask(ctx, id)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Service) CreateTask(ctx *gin.Context, task entity.Task) error {
	if err := s.Repo.CreateTask(ctx, &task); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateTask(ctx *gin.Context, newTask entity.Task, id string) error {
	task, err := s.Repo.GetTask(ctx, id)
	if err != nil {
		return err
	}

	task.Description = newTask.Description
	task.DueDate = newTask.DueDate
	task.AssignedTo = newTask.AssignedTo
	task.AssociatedDealID = newTask.AssociatedDealID

	if err = s.Repo.SaveTask(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTask(ctx *gin.Context, id string) error {
	company, err := s.Repo.GetTask(ctx, id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteTask(ctx, id, company); err != nil {
		return err
	}

	return nil
}
