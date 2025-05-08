package service

import (
	"context"

	"github.com/patricksferraz/pinned-company/domain/entity"
	"github.com/patricksferraz/pinned-company/domain/repo"
	"github.com/patricksferraz/pinned-company/infra/client/kafka/topic"
	"github.com/patricksferraz/pinned-company/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateCompany(ctx context.Context, name *string) (*string, error) {
	company, err := entity.NewCompany(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateCompany(ctx, company); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(company)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_COMPANY), utils.PString(string(eMsg)), company.ID)
	if err != nil {
		return nil, err
	}

	return company.ID, nil
}

func (s *Service) FindCompany(ctx context.Context, companyID *string) (*entity.Company, error) {
	company, err := s.Repo.FindCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *Service) SearchCompanies(ctx context.Context, pageToken *string, pageSize *int) ([]*entity.Company, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchCompanies, err := entity.NewSearchCompanies(pagination)
	if err != nil {
		return nil, nil, err
	}

	companies, nextPageToken, err := s.Repo.SearchCompanies(ctx, searchCompanies)
	if err != nil {
		return nil, nil, err
	}

	return companies, nextPageToken, nil
}
