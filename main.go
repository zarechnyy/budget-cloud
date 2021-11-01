package main

import (
	"budget-cloud/controllers"
	"budget-cloud/driver"
	"budget-cloud/models"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	db := driver.ConnectDB()
	dataStore := models.DataStore{DB: db}

	budgetController := controllers.BudgetController{DataStore: &dataStore}
	router := mux.NewRouter()

	router.Handle("/budget/new", budgetController.CreateBudget()).Methods("POST")
	router.Handle("/budgets", budgetController.UpdateBudget()).Methods("POST")
	router.Handle("/budgets", budgetController.FetchAll()).Methods("GET")
	router.Handle("/budget/{id}", budgetController.FetchBudget()).Methods("GET")
	router.Handle("/budget/{id}", budgetController.DeleteBudget()).Methods("DELETE")

	fmt.Println("Server is listening...")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8181", loggedRouter))
}
