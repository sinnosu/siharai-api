package controller

import (
	"net/http"
	"siharai-api/src/model"
	"siharai-api/src/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IInvoiceController interface {
	CreateInvoice(c echo.Context) error
	GetInvoicesWithinPeriod(c echo.Context) error
}

type invoiceController struct {
    iu usecase.IInvoiceUsecase
}

func NewInvoiceController(iu usecase.IInvoiceUsecase) IInvoiceController {
    return &invoiceController{iu}
}

func (ic *invoiceController) GetInvoicesWithinPeriod(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	fromDate := c.QueryParam("fromDate")
	toDate := c.QueryParam("toDate")
	invoicesRes, err := ic.iu.GetInvoicesWithinPeriod(uint(userId.(float64)), fromDate, toDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoicesRes)
}

func (ic *invoiceController) CreateInvoice(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	companyId := claims["company_id"]

	var request model.InvoicesRequests
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// check if the company_id in the request is the same as the user's company_id
	for _, invoice := range request.Invoices {
		if invoice.CompanyID != int(companyId.(float64)) {
			return c.JSON(http.StatusBadRequest, "CompanyID in the request is invalid")
		}
	}

	invoices, err := ic.iu.CreateAndSaveInvoices(request.Invoices)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

    return c.JSON(http.StatusOK, invoices)
}
