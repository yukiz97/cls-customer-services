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

type CustomerWithProductInfo struct {
	Customer
	Devices []CustomerDeviceSimplifyInfo `json:"devices"`
	Licenses []CustomerLicenseSimplifyInfo `json:"licenses"`
}

//CustomerDeviceSimplifyInfo use for store device info of customer
type CustomerDeviceSimplifyInfo struct {
	ID int `json:"id"`
	Product string `json:"product"`
	Serial string `json:"serial"`
	ExpireDate string `json:"expiredate"`
}

//CustomerLicenseSimplifyInfo use for store license info of customer
type CustomerLicenseSimplifyInfo struct {
	ID int `json:"id"`
	Product string `json:"license"`
	Code string `json:"code"`
	ExpireDate string `json:"expiredate"`
}
