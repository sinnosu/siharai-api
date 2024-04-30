package usecase

import (
	"siharai-api/src/model"
	"siharai-api/src/repository"
	"siharai-api/src/validator"
	"time"

	"github.com/shopspring/decimal"
)

type IInvoiceUsecase interface {
	GetInvoicesWithinPeriod(userId uint, fromDate string, toDate string) ([]model.InvoiceResponse, error)
	CreateAndSaveInvoices(invoices []model.InvoiceInput) ([]model.Invoice, error)
}

type invoiceUsecase struct {
	ir repository.IInvoiceRepository
	iv validator.IInvoiceValidator
}

func NewInvoiceUsecase(ir repository.IInvoiceRepository, iv validator.IInvoiceValidator) IInvoiceUsecase {
	return &invoiceUsecase{ir, iv}
}

func (iu *invoiceUsecase) GetInvoicesWithinPeriod(userId uint, fromDate string, toDate string) ([]model.InvoiceResponse, error) {
	invoices := []model.Invoice{}
	start, _ := time.Parse("2006-01-02", fromDate)
	end, _ := time.Parse("2006-01-02", toDate)
	if err := iu.ir.GetInvoicesWithinPeriod(&invoices, userId, start, end); err != nil {
		return nil, err
	}
	resInvoices := []model.InvoiceResponse{}
	for _, v := range invoices {
		i := model.InvoiceResponse{
			ID:             v.ID,
			CompanyID:      v.CompanyID,
			ClientID:       v.ClientID,
			IssueDate:      v.IssueDate,
			PaymentAmount:  v.PaymentAmount,
			Fee:            v.Fee,
			Tax:            v.Tax,
			BilledAmount:   v.BilledAmount,
			PaymentDueDate: v.PaymentDueDate,
			Status:         v.Status,
		}
		resInvoices = append(resInvoices, i)
	}
	return resInvoices, nil
}

func (iu *invoiceUsecase) CreateAndSaveInvoices(invoices []model.InvoiceInput) ([]model.Invoice, error) {
	var savedInvoices []model.Invoice
	for _, invoice := range invoices {
		if err := iu.iv.InvoiceValidate(invoice); err != nil {
			return nil, err
		}

		// 請求金額の計算
		fee := invoice.PaymentAmount.Mul(decimal.NewFromFloat(0.04))
		tax := fee.Mul(decimal.NewFromFloat(0.10))
		billedAmount := invoice.PaymentAmount.Add(fee).Add(tax)

		invoiceData := model.Invoice{
			CompanyID:      invoice.CompanyID,
			ClientID:       invoice.ClientID,
			IssueDate:      invoice.IssueDate,
			PaymentAmount:  invoice.PaymentAmount,
			Fee:            fee,
			Tax:            tax,
			BilledAmount:   billedAmount,
			PaymentDueDate: invoice.PaymentDueDate,
			Status:         model.StatusPending,
		}

		err := iu.ir.CreateInvoice(&invoiceData)
		if err != nil {
			return nil, err
		}

		savedInvoices = append(savedInvoices, invoiceData)
	}

	return savedInvoices, nil
}
