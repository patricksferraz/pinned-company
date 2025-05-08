package repo

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/pinned-company/domain/entity"
	"github.com/patricksferraz/pinned-company/infra/client/kafka"
	"github.com/patricksferraz/pinned-company/infra/db"
)

type Repository struct {
	Orm *db.DbOrm
	Kp  *kafka.KafkaProducer
}

func NewRepository(orm *db.DbOrm, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Orm: orm,
		Kp:  kp,
	}
}

func (r *Repository) CreateCompany(ctx context.Context, company *entity.Company) error {
	err := r.Orm.Db.Create(company).Error
	return err
}

func (r *Repository) FindCompany(ctx context.Context, companyID *string) (*entity.Company, error) {
	var e entity.Company

	r.Orm.Db.First(&e, "id = ?", *companyID)

	if e.ID == nil {
		return nil, fmt.Errorf("no company found")
	}

	return &e, nil
}

func (r *Repository) SaveCompany(ctx context.Context, company *entity.Company) error {
	err := r.Orm.Db.Save(company).Error
	return err
}

func (r *Repository) SearchCompanies(ctx context.Context, searchCompany *entity.SearchCompanies) ([]*entity.Company, *string, error) {
	var e []*entity.Company

	q := r.Orm.Db
	if *searchCompany.PageToken != "" {
		q = q.Where("token < ?", *searchCompany.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchCompany.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) < *searchCompany.PageSize {
		return e, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
