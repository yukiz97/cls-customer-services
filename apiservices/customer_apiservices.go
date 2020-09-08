package apiservices

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yukiz97/cls-customer-services/lcservices"
	"github.com/yukiz97/cls-customer-services/models"
	"github.com/yukiz97/utils/restapi"
	"log"
	"net/http"
	"strconv"
)

func home(response http.ResponseWriter, _ *http.Request) {
	restapi.RespondWithJSON(response, http.StatusOK, "Welcome to restful API of cls - customer services")
}

func insertCustomer(response http.ResponseWriter, request *http.Request) {
	modelInput := models.Customer{}

	err := json.NewDecoder(request.Body).Decode(&modelInput)

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	if modelInput.Name == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`name` must not be empty")
		return
	} else if modelInput.Address == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`address` must not be empty")
		return
	} else if modelInput.Email == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`email` must not be empty")
		return
	}

	idInserted := lcservices.InsertCustomer(modelInput)

	if idInserted != 0 {
		restapi.RespondWithJSON(response, http.StatusOK, restapi.IDItemAndMessage{ID: idInserted, Message: "Insert new customer successfully"})
	} else {
		restapi.RespondWithError(response, http.StatusBadRequest, "Insert new customer fail, please try again!")
	}
}

func updateCustomer(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idCustomer, err := strconv.Atoi(vars["id"])

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, "ID customer must be a integer")
		return
	}

	modelInput := models.Customer{}
	err = json.NewDecoder(request.Body).Decode(&modelInput)

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	if modelInput.Name == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`name` must not be empty")
		return
	} else if modelInput.Address == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`address` must not be empty")
		return
	} else if modelInput.Email == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`email` must not be empty")
		return
	}
	modelInput.ID = idCustomer
	isUpdated := lcservices.UpdateCustomer(modelInput)

	if isUpdated {
		restapi.RespondWithJSON(response, http.StatusOK, restapi.IDItemAndMessage{ID: idCustomer, Message: "Updated customer successfully"})
	} else {
		restapi.RespondWithError(response, http.StatusBadRequest, "Customer with ID "+strconv.Itoa(idCustomer)+" doesn't exists or value doesn't change!")
	}
}

func deleteCustomer(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idCustomer, err := strconv.Atoi(vars["id"])

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, "ID customer must be a integer")
		return
	}

	isDeleted := lcservices.DeleteCustomer(idCustomer)

	if isDeleted {
		restapi.RespondWithJSON(response, http.StatusOK, restapi.IDItemAndMessage{ID: idCustomer, Message: "Deleted customer successfully"})
	} else {
		restapi.RespondWithError(response, http.StatusBadRequest, "Customer with ID "+strconv.Itoa(idCustomer)+" doesn't exists!")
	}
}

func getCustomerList(response http.ResponseWriter, _ *http.Request) {
	listCustomer, arrId := lcservices.GetCustomerList("")
	listCustomerReturn := make([]models.CustomerWithProductInfo,0)

	if len(listCustomer) > 0 {
		mapDeviceInfo := lcservices.GetCustomerDeviceSimplifyInfo(arrId)
		mapLicenseInfo := lcservices.GetCustomerLicenseSimplifyInfo(arrId)
		for _, model := range listCustomer {
			devices := make([]models.CustomerDeviceSimplifyInfo,0)
			licenses := make([]models.CustomerLicenseSimplifyInfo,0)

			if _, ok := mapDeviceInfo[model.ID]; ok {
				devices = mapDeviceInfo[model.ID]
			}

			if _, ok := mapLicenseInfo[model.ID]; ok {
				licenses = mapLicenseInfo[model.ID]
			}

			modelReturn := models.CustomerWithProductInfo{Customer: model, Devices: devices,Licenses: licenses}
			listCustomerReturn = append(listCustomerReturn, modelReturn)
		}
	}

	restapi.RespondWithJSON(response, http.StatusOK, listCustomerReturn)
}

func getCustomer(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idCustomer, err := strconv.Atoi(vars["id"])

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, "ID customer must be a integer")
		return
	}

	modelCustomer := lcservices.GetCustomerByID(idCustomer)
	if modelCustomer.ID == 0 {
		restapi.RespondWithError(response, http.StatusBadRequest, "Customer with ID "+strconv.Itoa(idCustomer)+" doest not exist!")
		return
	}

	restapi.RespondWithJSON(response, http.StatusOK, modelCustomer)
}

func searchCustomerList(response http.ResponseWriter, request *http.Request) {
	modelInput := models.SearchCustomer{}

	err := json.NewDecoder(request.Body).Decode(&modelInput)

	if err != nil {
		restapi.RespondWithError(response, http.StatusBadRequest, err.Error())
		return
	}

	if modelInput.Keyword == "" {
		restapi.RespondWithError(response, http.StatusBadRequest, "`keyword` must not be empty")
		return
	}

	listCustomer,_ := lcservices.GetCustomerList(modelInput.Keyword)

	restapi.RespondWithJSON(response, http.StatusOK, listCustomer)
}

//InitRestfulAPIServices init customer restfull api
func InitRestfulAPIServices(listenPort int) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home)
	router.HandleFunc("/getCustomerList/", getCustomerList).Methods("GET")
	router.HandleFunc("/getCustomer/id/{id}", getCustomer).Methods("GET")

	router.HandleFunc("/insertCustomer/", insertCustomer).Methods("POST")
	router.HandleFunc("/searchCustomerList/", searchCustomerList).Methods("POST")

	router.HandleFunc("/updateCustomer/id/{id}", updateCustomer).Methods("PUT")

	router.HandleFunc("/deleteCustomer/id/{id}", deleteCustomer).Methods("DELETE")

	println("Running CLS - customer services.... - Listen to port :" + strconv.Itoa(listenPort))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(listenPort), router))
}
