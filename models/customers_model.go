package models

import "time"

//Customer struct present databse customer table
type Customer struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Email      string    `json:"email"`
	CreateDate time.Time `json:"createdate"`
}
