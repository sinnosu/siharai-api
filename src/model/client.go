package model

import "time"

type Client struct {
    ID           int    `json:"id"`
    CompanyID          int    `json:"company_id"`
    ClientName         string `json:"client_name"`
    RepresentativeName string `json:"representative_name"`
    PhoneNumber        string `json:"phone_number"`
    PostalCode         string `json:"postal_code"`
    Address            string `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}
