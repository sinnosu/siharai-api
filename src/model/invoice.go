package model

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
    ID      int             `json:"id"`
    CompanyID      int             `json:"company_id"`
    ClientID       int             `json:"client_id"`
    IssueDate      time.Time       `json:"issue_date"`
    PaymentAmount  decimal.Decimal `json:"payment_amount"`
    Fee            decimal.Decimal `json:"fee"`
    Tax            decimal.Decimal `json:"tax"`
    BilledAmount   decimal.Decimal `json:"billed_amount"`
    PaymentDueDate time.Time       `json:"payment_due_date"`
    Status         string          `json:"status"`
    CreatedAt      time.Time       `json:"created_at"`
    UpdatedAt      time.Time       `json:"updated_at"`
}

type InvoicesRequests struct {
	Invoices []InvoiceInput `json:"invoices"`
}

type InvoiceInput struct {
    CompanyID      int             `json:"company_id"`
    ClientID       int             `json:"client_id"`
    IssueDate      time.Time       `json:"issue_date"`
    PaymentAmount  decimal.Decimal `json:"payment_amount"`
    PaymentDueDate time.Time       `json:"payment_due_date"`
}

type InvoiceResponse struct {
    ID      int             `json:"id"`
    CompanyID      int             `json:"company_id"`
    ClientID       int             `json:"client_id"`
    IssueDate      time.Time       `json:"issue_date"`
    PaymentAmount  decimal.Decimal `json:"payment_amount"`
    Fee            decimal.Decimal `json:"fee"`
    Tax            decimal.Decimal `json:"tax"`
    BilledAmount   decimal.Decimal `json:"billed_amount"`
    PaymentDueDate time.Time       `json:"payment_due_date"`
    Status         string          `json:"status"`
}


// JSON 文字列から time.Time 型に変換するカスタム関数
func (i *InvoiceInput) UnmarshalJSON(data []byte) error {
    type Alias InvoiceInput
    aux := &struct {
        IssueDate      string `json:"issue_date"`
        PaymentDueDate string `json:"payment_due_date"`
        *Alias
    }{
        Alias: (*Alias)(i),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    var err error
    i.IssueDate, err = time.Parse("2006-01-02", aux.IssueDate)
    if err != nil {
        return err
    }
    i.PaymentDueDate, err = time.Parse("2006-01-02", aux.PaymentDueDate)
    if err != nil {
        return err
    }

    return nil
}

const (
    StatusPending      = "Pending"
    StatusProcessing   = "Processing"
    StatusPaid         = "Paid"
    StatusError        = "Error"
)