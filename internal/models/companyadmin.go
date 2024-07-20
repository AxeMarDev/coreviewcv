package models

type CompanyAdmin struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Hashpassword string `json:"hash_password"`
}
