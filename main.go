package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Initialize scanner for user input
	InitScanner()

	// Initialize database
	err := InitDB()
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	defer CloseDB()

	// Check if monthly budget is set, if not, prompt user
	budget, err := GetMonthlyBudget()
	if err != nil {
		fmt.Printf("Error checking monthly budget: %v\n", err)
		os.Exit(1)
	}

	if budget == 0 {
		// First run - prompt for monthly budget
		budget, err = PromptMonthlyBudget()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		err = SetMonthlyBudget(budget)
		if err != nil {
			fmt.Printf("Failed to save monthly budget: %v\n", err)
			os.Exit(1)
		}
		DisplaySuccess(fmt.Sprintf("Monthly budget set to ₹%.2f", budget))
	}

	// Main menu loop
	for {
		DisplayMenu()
		choice := GetUserInput()

		switch choice {
		case "1":
			// Add Expense
			expense, err := AddExpensePrompt()
			if err != nil {
				DisplayError(err)
				continue
			}
			
			// Check if adding this expense would exceed monthly budget
			budget, err := GetMonthlyBudget()
			if err == nil && budget > 0 {
				now := time.Now()
				currentMonthTotal, err := GetTotalExpensesForMonth(now.Year(), now.Month())
				if err == nil {
					// Check if expense is in current month
					if expense.Date.Year() == now.Year() && expense.Date.Month() == now.Month() {
						newTotal := currentMonthTotal + expense.Amount
						if newTotal > budget {
							overAmount := newTotal - budget
							DisplayWarning(fmt.Sprintf("Adding this expense will exceed your monthly budget by ₹%.2f!", overAmount))
						}
					}
				}
			}
			
			err = AddExpense(expense)
			if err != nil {
				DisplayError(err)
			} else {
				DisplaySuccess("Expense added successfully!")
			}

		case "2":
			// View All Expenses
			expenses, err := GetAllExpenses()
			if err != nil {
				DisplayError(err)
			} else {
				DisplayExpenses(expenses)
			}

		case "3":
			// View Expenses by Category
			category := PromptCategory()
			if category == "" {
				DisplayError(fmt.Errorf("category cannot be empty"))
				continue
			}
			expenses, err := GetExpensesByCategory(category)
			if err != nil {
				DisplayError(err)
			} else {
				DisplayExpenses(expenses)
			}

		case "4":
			// View Expenses by Date Range
			startDate, endDate, err := PromptDateRange()
			if err != nil {
				DisplayError(err)
				continue
			}
			expenses, err := GetExpensesByDateRange(startDate, endDate)
			if err != nil {
				DisplayError(err)
			} else {
				DisplayExpenses(expenses)
			}

		case "5":
			// Search Expenses
			searchTerm := PromptSearchTerm()
			if searchTerm == "" {
				DisplayError(fmt.Errorf("search term cannot be empty"))
				continue
			}
			expenses, err := SearchExpenses(searchTerm)
			if err != nil {
				DisplayError(err)
			} else {
				DisplayExpenses(expenses)
			}

		case "6":
			// Delete Expense
			id, err := PromptExpenseID()
			if err != nil {
				DisplayError(err)
				continue
			}
			err = DeleteExpense(id)
			if err != nil {
				DisplayError(err)
			} else {
				DisplaySuccess(fmt.Sprintf("Expense with ID %d deleted successfully!", id))
			}

		case "7":
			// View Statistics
			err := DisplayStatistics()
			if err != nil {
				DisplayError(err)
			}

		case "8":
			// Exit
			fmt.Println("\nThank you for using Expense Tracker!")
			os.Exit(0)

		default:
			fmt.Println("\nInvalid option. Please select a number between 1-8.")
		}
	}
}
