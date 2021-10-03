package models

import (
	"gorm.io/gorm"
	"time"
)

type DataStorer interface {
	CreateBudget(name string, maxExpense int) (Budget, error)
	UpdateBudget(budget Budget) (Budget, error)
	FetchAllBudget() ([]Budget, error)
	FetchBudget(id int) (Budget, error)
	DeleteBudget(id int) error
}

type DataStore struct {
	DB *gorm.DB
}

type Budget struct {
	BudgetID int `gorm:"column:budget_id;primary_key;auto_increment:true" json:"id"`
	Name string `gorm:"column:budget_name;type:varchar(50)" json:"name"`
	MaximumExpense int `gorm:"column:maximum_expense" json:"maximum_expense"`
	CurrentExpense int `gorm:"column:current_expense" json:"current_expense"`
	Updated   int64 `gorm:"autoUpdateTime:nano" json:"updated"`	// Use unix nano seconds as updating time
	Created   int64 `gorm:"autoCreateTime:nano" json:"created"`	// Use unix nano seconds as creating time
}

func (dataStore *DataStore) CreateBudget(name string, maxExpense int) (Budget, error) {
	budget := Budget {
		Name: name,
		MaximumExpense: maxExpense,
		CurrentExpense: 0,
	}
	result := dataStore.DB.Create(&budget)
	err := result.Error
	if err != nil {
		return Budget{}, err
	}
	return budget, nil
}

func (dataStore *DataStore) UpdateBudget(budget Budget) (Budget, error) {
	 newBudget := Budget{
		 Name: budget.Name,
		 MaximumExpense: budget.MaximumExpense,
		 CurrentExpense: budget.CurrentExpense,
		 Updated: int64(time.Now().Nanosecond()),
	 }
	result := dataStore.DB.Model(&budget).Updates(newBudget)
	err := result.Error

	if err != nil {
		return Budget{}, err
	}

	return budget, nil
}

func (dataStore *DataStore) FetchAllBudget() ([]Budget, error) {
	var budgets []Budget
	result := dataStore.DB.Find(&budgets)
	err := result.Error
	if err != nil {
		return []Budget{}, err
	}
	return budgets, nil
}

func (dataStore *DataStore) FetchBudget(id int) (Budget, error) {
	budget := Budget{}
	result := dataStore.DB.Where("id = ?", id).First(&budget)
	err := result.Error
	if err != nil {
		return Budget{}, err
	}
	return budget, nil
}

func (dataStore *DataStore) DeleteBudget(id int) error {
	result := dataStore.DB.Delete(Budget{}, id)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}