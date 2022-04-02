package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-company/domain/service"
	"github.com/gofiber/fiber/v2"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateCompany godoc
// @Summary create a new company
// @ID createCompany
// @Tags Company
// @Description Router for create a new company
// @Accept json
// @Produce json
// @Param body body CreateCompanyRequest true "JSON body for create a new company"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /companies [post]
func (t *RestService) CreateCompany(c *fiber.Ctx) error {
	var req CreateCompanyRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	companyID, err := t.Service.CreateCompany(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *companyID})
}

// FindCompany godoc
// @Summary find a company
// @ID findCompany
// @Tags Company
// @Description Router for find a company
// @Accept json
// @Produce json
// @Param company_id path string true "Company ID"
// @Success 200 {object} Company
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /companies/{company_id} [get]
func (t *RestService) FindCompany(c *fiber.Ctx) error {
	companyID := c.Params("company_id")
	if !govalidator.IsUUIDv4(companyID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "company_id is not a valid uuid",
		})
	}

	company, err := t.Service.FindCompany(c.Context(), &companyID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}

// SearchCompanies godoc
// @Summary search companies
// @ID searchCompanies
// @Tags Company
// @Description Router for search companies
// @Accept json
// @Produce json
// @Param page_size query int false "page size"
// @Param page_token query string false "page token"
// @Success 200 {object} SearchCompaniesResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /companies [get]
func (t *RestService) SearchCompanies(c *fiber.Ctx) error {
	var req SearchCompaniesRequest

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	companies, nextPageToken, err := t.Service.SearchCompanies(c.Context(), &req.PageToken, &req.PageSize)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"companies":       companies,
		"next_page_token": nextPageToken,
	})
}
