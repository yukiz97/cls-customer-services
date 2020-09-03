package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/yukiz97/cls-customer-services/lcservices"
	"github.com/yukiz97/utils/config"
)

type configuration struct {
	MYSQLHost     string
	MYSQLUsername string
	MYSQLPassword string
	MYSQLDB       string
}

func main() {
	configuration := configuration{}
	mapConfig := config.ParseJSONConfigToMap("D:\\DevApps\\_Workspace\\Golang\\.mydata\\cls-services\\config.json")
	err := mapstructure.Decode(mapConfig, &configuration)

	if err != nil {
		log.Fatal(err)
	}

	lcservices.InitLocalServices(configuration.MYSQLHost, configuration.MYSQLUsername, configuration.MYSQLPassword, configuration.MYSQLDB)

	fmt.Printf("%+v\n",lcservices.GetCustomerByID(5))
}
