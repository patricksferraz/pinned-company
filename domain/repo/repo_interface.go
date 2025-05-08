package repo

import (
	"context"

	"github.com/patricksferraz/pinned-company/domain/entity"
)

type RepoInterface interface {
	CreateCompany(ctx context.Context, company *entity.Company) error
	FindCompany(ctx context.Context, companyID *string) (*entity.Company, error)
	SaveCompany(ctx context.Context, company *entity.Company) error
	SearchCompanies(ctx context.Context, searchCompany *entity.SearchCompanies) ([]*entity.Company, *string, error)

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
