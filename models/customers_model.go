package models

//SearchCustomer struct for search customer
type SearchCustomer struct {
	Keyword string `json:"keyword"`
}

//Customer struct present database customer table
type Customer struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	CreateDate string `json:"createdate"`
}
