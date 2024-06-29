package models

type Company struct {
	ID               string `json:"id"`
	Companyname      string `json:"company_name"`
	Companycode      string `json:"company_code"`
	MasteremployeeId int    `json:"masteremployee_id"`
}
