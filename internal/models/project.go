package models

type Project struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Companyid int    `json:"company_id"`
}
