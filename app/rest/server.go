package rest

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/patricksferraz/pinned-company/app/rest/docs"
	"github.com/patricksferraz/pinned-company/domain/service"
	"github.com/patricksferraz/pinned-company/infra/client/kafka"
	"github.com/patricksferraz/pinned-company/infra/db"
	"github.com/patricksferraz/pinned-company/infra/repo"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Company Swagger API
// @version 1.0
// @description Swagger API for Company Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email contato@coding4u.com.br

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(orm *db.DbOrm, kp *kafka.KafkaProducer, port int) {
	r := fiber.New()
	r.Use(cors.New())

	repository := repo.NewRepository(orm, kp)
	service := service.NewService(repository)
	restService := NewRestService(service)

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/swagger/*", fiberSwagger.WrapHandler)
	{
		company := v1.Group("/companies")
		company.Get("", restService.SearchCompanies)
		company.Post("", restService.CreateCompany)
		company.Get("/:company_id", restService.FindCompany)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
