package lcservices

import (
	"github.com/yukiz97/utils"
	"time"

	"github.com/yukiz97/cls-customer-services/models"
	"github.com/yukiz97/utils/date"
	"github.com/yukiz97/utils/dbcon"
)

var strDBConnect string
var mapCustomerField map[string]string

//InsertCustomer insert customer by value of struct
func InsertCustomer(modelCustomer models.Customer) int64 {
	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	insertQuery, err := db.Prepare("INSERT INTO Customer(" + mapCustomerField["name"] + ", " + mapCustomerField["address"] + ", " + mapCustomerField["email"] + ") VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	result, errInsert := insertQuery.Exec(modelCustomer.Name, modelCustomer.Address, modelCustomer.Email)

	if errInsert != nil {
		panic(errInsert)
	}

	idInserted, _ := result.LastInsertId()

	return idInserted
}

//UpdateCustomer update customer by value of struct
func UpdateCustomer(modelCustomer models.Customer) bool {
	isUpdated := true
	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	updateQuery, err := db.Prepare("UPDATE Customer SET " + mapCustomerField["name"] + " = ?," + mapCustomerField["address"] + " = ?, " + mapCustomerField["email"] + " = ? WHERE " + mapCustomerField["id"] + " = ?")
	if err != nil {
		panic(err.Error())
	}
	result, errUpdate := updateQuery.Exec(modelCustomer.Name, modelCustomer.Address, modelCustomer.Email, modelCustomer.ID)

	if errUpdate != nil {
		panic(errUpdate)
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected == 0 {
		isUpdated = false
	}

	return isUpdated
}

//DeleteCustomer delete customer by id
func DeleteCustomer(idCustomer int) bool {
	isDeleted := true
	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	deleteQuery, err := db.Prepare("DELETE FROM Customer WHERE " + mapCustomerField["id"] + " = ?")
	if err != nil {
		panic(err.Error())
	}
	result, errDelete := deleteQuery.Exec(idCustomer)

	if errDelete != nil {
		panic(errDelete)
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected == 0 {
		isDeleted = false
	}

	return isDeleted
}

//GetCustomerList get customer list by keyword
func GetCustomerList(keyWord string) ([]models.Customer,[]int) {
	listCustomer := make([]models.Customer, 0)
	listCustomerId := make([]int, 0)
	keyWord = "%" + keyWord + "%"

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()

	selectQuery, err := db.Prepare("SELECT * FROM Customer WHERE " + mapCustomerField["name"] + " LIKE ?")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query(keyWord)
	for result.Next() {
		var createDate time.Time
		modelCustomer := models.Customer{}

		result.Scan(&modelCustomer.ID, &modelCustomer.Name, &modelCustomer.Address, &modelCustomer.Email, &createDate)
		modelCustomer.CreateDate = date.FormatTimeToString(createDate, date.Format1)
		listCustomer = append(listCustomer, modelCustomer)
		listCustomerId = append(listCustomerId, modelCustomer.ID)
	}

	return listCustomer , listCustomerId
}

//GetCustomerByID get customer by customer id
func GetCustomerByID(idCustomer int) models.Customer {
	var modelCustomer models.Customer

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()

	selectQuery, err := db.Prepare("SELECT * FROM Customer WHERE " + mapCustomerField["id"] + " = ?")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query(idCustomer)
	for result.Next() {
		var createDate time.Time
		modelCustomer = models.Customer{}

		result.Scan(&modelCustomer.ID, &modelCustomer.Name, &modelCustomer.Address, &modelCustomer.Email, &createDate)
		modelCustomer.CreateDate = date.FormatTimeToString(createDate, date.Format1)
	}

	return modelCustomer
}
//GetCustomerDeviceSimplifyInfo get customer device info of arr customer id
func GetCustomerDeviceSimplifyInfo(arrId []int) map[int][]models.CustomerDeviceSimplifyInfo{
	if len(arrId) == 0 {
		return nil
	}
	mapDeviceInfo := make(map[int][]models.CustomerDeviceSimplifyInfo)

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	strIds := utils.IntArrayToString(arrId,",")
	selectQuery, err := db.Prepare("SELECT IdDevice,ProductName,DeviceSerial,GuaranteeExpireDate,IdCustomer FROM device JOIN product ON device.IdProduct = product.IdProduct WHERE IdCustomer IN ("+strIds+")")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query()

	for result.Next() {
		var expireDate time.Time
		var idCustomer int
		modelDevice := models.CustomerDeviceSimplifyInfo{}

		result.Scan(&modelDevice.ID, &modelDevice.Product, &modelDevice.Serial, &expireDate, &idCustomer)
		modelDevice.ExpireDate = date.FormatTimeToString(expireDate, date.Format2)

		if _, ok := mapDeviceInfo[idCustomer]; !ok {
			mapDeviceInfo[idCustomer] = make([]models.CustomerDeviceSimplifyInfo,0)
		}
		mapDeviceInfo[idCustomer] = append(mapDeviceInfo[idCustomer],modelDevice)
	}

	return mapDeviceInfo
}
//GetCustomerLicenseSimplifyInfo get customer license info of arr customer id
func GetCustomerLicenseSimplifyInfo(arrId []int) map[int][]models.CustomerLicenseSimplifyInfo{
	if len(arrId) == 0 {
		return nil
	}
	mapLicenseInfo := make(map[int][]models.CustomerLicenseSimplifyInfo)

	db := dbcon.InitDBMySQL(strDBConnect)
	defer db.Close()
	strIds := utils.IntArrayToString(arrId,",")
	selectQuery, err := db.Prepare("SELECT IdLicense,ProductName,LicenseCode,ExpireDate,IdCustomer FROM license JOIN product ON license.IdProduct = product.IdProduct WHERE IdCustomer IN ("+strIds+")")
	if err != nil {
		panic(err.Error())
	}
	result, _ := selectQuery.Query()

	for result.Next() {
		var expireDate time.Time
		var idCustomer int
		modelLicense := models.CustomerLicenseSimplifyInfo{}

		result.Scan(&modelLicense.ID, &modelLicense.Product, &modelLicense.Code, &expireDate, &idCustomer)
		modelLicense.ExpireDate = date.FormatTimeToString(expireDate, date.Format2)

		if _, ok := mapLicenseInfo[idCustomer]; !ok {
			mapLicenseInfo[idCustomer] = make([]models.CustomerLicenseSimplifyInfo,0)
		}
		mapLicenseInfo[idCustomer] = append(mapLicenseInfo[idCustomer], modelLicense)
	}

	return mapLicenseInfo
}

//InitLocalServices init value for database functions
func InitLocalServices(host string, userName string, password string, db string) {
	strDBConnect = dbcon.GetMySQLOpenConnectString(host, userName, password, db)

	mapCustomerField = make(map[string]string)
	mapCustomerField["id"] = "IdCustomer"
	mapCustomerField["name"] = "CustomerName"
	mapCustomerField["address"] = "CustomerAddress"
	mapCustomerField["email"] = "CustomerEmail"
	mapCustomerField["createdate"] = "CreateDate"
}
