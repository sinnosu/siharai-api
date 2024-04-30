package usecase

import (
	"siharai-api/src/model"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// リポジトリのモック
type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) CreateInvoice(invoice *model.Invoice) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func (m *MockInvoiceRepository) GetInvoicesWithinPeriod(invoices *[]model.Invoice, userId uint, start time.Time, end time.Time) error {
	err := m.Called(invoices, userId, start, end)
	if err != nil {
		return err.Error(0)
	}
	return err.Error(1)
}

// バリデータのモック
type MockInvoiceValidator struct {
	mock.Mock
}

func (m *MockInvoiceValidator) InvoiceValidate(invoice model.InvoiceInput) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func TestCreateInvoices(t *testing.T) {
	mockRepo := new(MockInvoiceRepository)
	mockValidator := new(MockInvoiceValidator)

	// 日付のフォーマット
	layout := "2006-01-02"
	issueDate1, _ := time.Parse(layout, "2021-01-01")
	paymentDueDate1, _ := time.Parse(layout, "2021-02-01")
	issueDate2, _ := time.Parse(layout, "2021-01-02")
	paymentDueDate2, _ := time.Parse(layout, "2021-02-02")

	invoices := []model.InvoiceInput{
		{CompanyID: 1, ClientID: 1, IssueDate: issueDate1, PaymentAmount: decimal.NewFromFloat(1000), PaymentDueDate: paymentDueDate1},
		{CompanyID: 1, ClientID: 2, IssueDate: issueDate2, PaymentAmount: decimal.NewFromFloat(2000), PaymentDueDate: paymentDueDate2},
	}

	uc := NewInvoiceUsecase(mockRepo, mockValidator)

	// モックの設定
	for _, invoice := range invoices {
		mockValidator.On("InvoiceValidate", invoice).Return(nil)
		fee := invoice.PaymentAmount.Mul(decimal.NewFromFloat(0.04))
		tax := fee.Mul(decimal.NewFromFloat(0.10))
		billedAmount := invoice.PaymentAmount.Add(fee).Add(tax)
		expectedInvoice := model.Invoice{
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
		mockRepo.On("CreateInvoice", &expectedInvoice).Return(nil)
	}

	result, err := uc.CreateAndSaveInvoices(invoices)

	// 検証
	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestGetInvoicesWithinPeriod(t *testing.T) {
	mockRepo := new(MockInvoiceRepository)
	uc := NewInvoiceUsecase(mockRepo, nil)

	userId := uint(1)
	startDate, _ := time.Parse("2006-01-02", "2021-01-01")
	endDate, _ := time.Parse("2006-01-02", "2021-01-31")

	validInvoices := []model.Invoice{
		{ID: 1, CompanyID: 1, ClientID: 1, IssueDate: startDate},
		{ID: 2, CompanyID: 1, ClientID: 2, IssueDate: startDate},
	}
	validResponses := make([]model.InvoiceResponse, len(validInvoices))
	for i, inv := range validInvoices {
		validResponses[i] = model.InvoiceResponse{
			ID:        inv.ID,
			CompanyID: inv.CompanyID,
			ClientID:  inv.ClientID,
			IssueDate: inv.IssueDate,
		}
	}

	emptyInvoices := []model.Invoice{}

	// Scenario 1: Valid data
	mockRepo.On("GetInvoicesWithinPeriod", mock.AnythingOfType("*[]model.Invoice"), userId, startDate, endDate).Run(func(args mock.Arguments) {
		invoicesPtr := args.Get(0).(*[]model.Invoice)
		*invoicesPtr = validInvoices
	}).Return(nil).Once()

	// Scenario 2: No records match the company ID
	mockRepo.On("GetInvoicesWithinPeriod", mock.AnythingOfType("*[]model.Invoice"), userId, startDate, endDate).Run(func(args mock.Arguments) {
		invoicesPtr := args.Get(0).(*[]model.Invoice)
		*invoicesPtr = emptyInvoices
	}).Return(nil).Once()

	// Scenario 3: No records within the specified period
	mockRepo.On("GetInvoicesWithinPeriod", mock.AnythingOfType("*[]model.Invoice"), userId, startDate, endDate).Run(func(args mock.Arguments) {
		invoicesPtr := args.Get(0).(*[]model.Invoice)
		*invoicesPtr = emptyInvoices
	}).Return(nil).Once()

	// Test for valid data
	invoices, err := uc.GetInvoicesWithinPeriod(userId, "2021-01-01", "2021-01-31")
	assert.NoError(t, err)
	assert.Equal(t, validResponses, invoices, "Should find invoices when valid data exists")

	// Test for no records match the company ID
	invoices, err = uc.GetInvoicesWithinPeriod(userId, "2021-01-01", "2021-01-31")
	assert.NoError(t, err)
	assert.Empty(t, invoices, "Should find no invoices when no records match the company ID")

	// Test for no records within the specified period
	invoices, err = uc.GetInvoicesWithinPeriod(userId, "2021-01-01", "2021-01-31")
	assert.NoError(t, err)
	assert.Empty(t, invoices, "Should find no invoices within the specified period")

	mockRepo.AssertExpectations(t)
}
