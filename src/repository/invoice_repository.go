package repository

import (
	"siharai-api/src/model"
	"time"

	"gorm.io/gorm"
)

type IInvoiceRepository interface {
	CreateInvoice(invoice *model.Invoice) error
	GetInvoicesWithinPeriod(invoices *[]model.Invoice, userId uint, start time.Time, end time.Time ) error
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) IInvoiceRepository {
    return &invoiceRepository{db}
}

func (ir *invoiceRepository) GetInvoicesWithinPeriod(invoices *[]model.Invoice, userId uint, start time.Time, end time.Time) error {
    sql := "SELECT i.id, i.company_id, i.client_id, i.issue_date, i.payment_amount, i.fee, i.tax, i.billed_amount, i.payment_due_date, i.status  FROM invoices i INNER JOIN users u ON i.company_id = u.company_id WHERE u.id = ? AND issue_date BETWEEN ? AND ?"
    if err := ir.db.Raw(sql, userId, start, end).Scan(invoices).Error; err != nil {
		return err
	}
	return nil
}

func (ir *invoiceRepository) CreateInvoice(invoice *model.Invoice) error {
	if err := ir.db.Create(invoice).Error; err != nil {
		return err
	}
	return nil
}
