// model/client_bank_account.go
package model

import "time"

type ClientBankAccount struct {
    ID          int    `json:"id"`
    ClientID           int    `json:"client_id"`
    BankName           string `json:"bank_name"`
    BranchName         string `json:"branch_name"`
    AccountNumber      string `json:"account_number"`
    AccountHolderName  string `json:"account_holder_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}
