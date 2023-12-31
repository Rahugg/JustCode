package service

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/entity"
)

func (s *Service) GetCompanies(ctx *gin.Context) (*[]entity.Company, error) {
	companies, err := s.Repo.GetCompanies(ctx)

	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *Service) GetCompany(ctx *gin.Context, id string) (*entity.Company, error) {
	company, err := s.Repo.GetCompany(ctx, id)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *Service) CreateCompany(ctx *gin.Context, company entity.Company) error {
	if err := s.Repo.CreateCompany(ctx, &company); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateCompany(ctx *gin.Context, newCompany entity.NewCompany, id string) error {
	company, err := s.Repo.GetCompany(ctx, id)
	if err != nil {
		return err
	}

	company.Name = newCompany.Name
	company.Address = newCompany.Address
	company.Phone = newCompany.Phone

	if err = s.Repo.SaveCompany(ctx, company); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteCompany(ctx *gin.Context, id string) error {
	company, err := s.Repo.GetCompany(ctx, id)
	if err != nil {
		return err
	}

	if err = s.Repo.DeleteCompany(ctx, id, company); err != nil {
		return err
	}

	return nil
}
