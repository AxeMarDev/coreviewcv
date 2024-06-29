package models

type Client struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Hashpassword string `json:"hash_password"`
	Phone        string `json:"phone"`
	Companyid    int    `json:"company_id"`
}
