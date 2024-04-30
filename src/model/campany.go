package model

import "time"

type Company struct {
    ID          int    `json:"id"`
    CompanyName        string `json:"company_name"`
    RepresentativeName string `json:"representative_name"`
    PhoneNumber        string `json:"phone_number"`
    PostalCode         string `json:"postal_code"`
    Address            string `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}
