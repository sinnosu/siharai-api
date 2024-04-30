package validator

import (
	"siharai-api/src/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IInvoiceValidator interface {
	InvoiceValidate(invoice model.InvoiceInput) error
}

type invoiceValidator struct{}

func NewInvoiceValidator() IInvoiceValidator {
	return &invoiceValidator{}
}

func (iv *invoiceValidator) InvoiceValidate(invoice model.InvoiceInput) error {
	return validation.ValidateStruct(&invoice,
		validation.Field(
			&invoice.CompanyID,
			validation.Required.Error("CompanyID is required"),
		),
	)
}
