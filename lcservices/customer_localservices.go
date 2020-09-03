package lcservices

import (
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
func GetCustomerList(keyWord string) []models.Customer {
	listCustomer := make([]models.Customer, 0)
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
	}

	return listCustomer
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
