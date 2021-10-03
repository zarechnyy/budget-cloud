package controllers

import (
	"budget-cloud/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
)

type BudgetController struct {
	DataStore models.DataStorer
}

func (c *BudgetController) CreateBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "NO POST", r.Method)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		requestBody := struct {
			Name      	   string
			MaxExpense     int
		}{}
		if err := json.Unmarshal(body, &requestBody); err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		budget, err := c.DataStore.CreateBudget(requestBody.Name, requestBody.MaxExpense)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		json.NewEncoder(w).Encode(budget)
	}
}

func (c *BudgetController) UpdateBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "NO POST", r.Method)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		var budget = models.Budget{}
		if err := json.Unmarshal(body, &budget); err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		updatedBudget, err := c.DataStore.UpdateBudget(budget)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		response := struct {
			Budget models.Budget `json:"budget"`
		}{}
		response.Budget = updatedBudget

		json.NewEncoder(w).Encode(response)
	}
}

func (c *BudgetController) FetchAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "NO GET", r.Method)
			return
		}

		budgets, err := c.DataStore.FetchAllBudget()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		response := struct {
			Budgets []models.Budget `json:"budgets"`
		}{}
		response.Budgets = budgets

		json.NewEncoder(w).Encode(response)
	}
}

func (c *BudgetController) FetchBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "NO GET", r.Method)
			return
		}

		urlString := r.RequestURI

		searchIDPath := path.Base(urlString)
		searchID, err := strconv.Atoi(searchIDPath)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fetchedBudget, err := c.DataStore.FetchBudget(searchID)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		response := struct {
			Budget models.Budget `json:"budget"`
		}{}
		response.Budget = fetchedBudget

		json.NewEncoder(w).Encode(response)
	}
}

func (c *BudgetController) DeleteBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "NO DELETE", r.Method)
			return
		}

		urlString := r.RequestURI

		searchIDPath := path.Base(urlString)
		searchID, err := strconv.Atoi(searchIDPath)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		err = c.DataStore.DeleteBudget(searchID)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		response := struct {
			Result string `json:"result"`
		}{}
		response.Result = "deleted"

		json.NewEncoder(w).Encode(response)
	}
}