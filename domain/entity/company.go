package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/pinned-company/utils"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Company struct {
	Base  `json:",inline" valid:"-"`
	Name  *string `json:"name" gorm:"column:name;not null" valid:"required"`
	Token *string `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
}

func NewCompany(name *string) (*Company, error) {
	token := primitive.NewObjectID().Hex()
	e := Company{
		Name:  name,
		Token: &token,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Company) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type SearchCompanies struct {
	Pagination `json:",inline" valid:"-"`
}

func NewSearchCompanies(pagination *Pagination) (*SearchCompanies, error) {
	e := SearchCompanies{}
	e.PageToken = pagination.PageToken
	e.PageSize = pagination.PageSize

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}
