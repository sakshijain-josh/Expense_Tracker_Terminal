package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var scanner *bufio.Scanner

// InitScanner initializes the scanner for user input
func InitScanner() {
	scanner = bufio.NewScanner(os.Stdin)
}

// DisplayMenu displays the main menu
func DisplayMenu() {
	fmt.Println("\n=== Expense Tracker ===")
	fmt.Println("1. Add Expense")
	fmt.Println("2. View All Expenses")
	fmt.Println("3. View Expenses by Category")
	fmt.Println("4. View Expenses by Date Range")
	fmt.Println("5. Search Expenses")
	fmt.Println("6. Delete Expense")
	fmt.Println("7. View Statistics")
	fmt.Println("8. Exit")
	fmt.Print("\nSelect an option: ")
}

// GetUserInput reads a line of input from the user
func GetUserInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// PromptMonthlyBudget prompts the user to enter monthly budget
func PromptMonthlyBudget() (float64, error) {
	fmt.Println("\n=== Welcome to Expense Tracker ===")
	fmt.Print("Enter your expected spending per month: ")
	input := GetUserInput()
	budget, err := strconv.ParseFloat(input, 64)
	if err != nil || budget <= 0 {
		return 0, fmt.Errorf("invalid budget amount. Please enter a positive number")
	}
	return budget, nil
}

// AddExpensePrompt prompts user for expense details and returns an Expense
func AddExpensePrompt() (Expense, error) {
	var expense Expense

	fmt.Println("\n=== Add New Expense ===")

	// Amount
	fmt.Print("Enter amount: ")
	amountStr := GetUserInput()
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		return expense, fmt.Errorf("invalid amount. Please enter a positive number")
	}
	expense.Amount = amount

	// Description
	fmt.Print("Enter description: ")
	expense.Description = GetUserInput()
	if expense.Description == "" {
		return expense, fmt.Errorf("description cannot be empty")
	}

	// Category
	fmt.Print("Enter category (e.g., Food, Transport, Entertainment): ")
	expense.Category = GetUserInput()
	if expense.Category == "" {
		return expense, fmt.Errorf("category cannot be empty")
	}

	// Date
	fmt.Print("Enter date (YYYY-MM-DD) or press Enter for today: ")
	dateStr := GetUserInput()
	if dateStr == "" {
		expense.Date = time.Now()
	} else {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return expense, fmt.Errorf("invalid date format. Please use YYYY-MM-DD")
		}
		expense.Date = date
	}

	// Payment Method
	fmt.Print("Enter payment method (Cash/UPI): ")
	paymentMethod := GetUserInput()
	paymentMethod = strings.ToUpper(strings.TrimSpace(paymentMethod))
	if paymentMethod != "CASH" && paymentMethod != "UPI" {
		return expense, fmt.Errorf("invalid payment method. Please enter 'Cash' or 'UPI'")
	}
	expense.PaymentMethod = paymentMethod

	// Notes
	fmt.Print("Enter notes (optional): ")
	expense.Notes = GetUserInput()

	return expense, nil
}

// DisplayExpenses displays a list of expenses in a formatted table
func DisplayExpenses(expenses []Expense) {
	if len(expenses) == 0 {
		fmt.Println("\nNo expenses found.")
		return
	}

	fmt.Println("\n" + strings.Repeat("=", 120))
	fmt.Printf("%-5s %-12s %-25s %-20s %-12s %-10s %-20s\n", "ID", "Amount", "Description", "Category", "Date", "Payment", "Notes")
	fmt.Println(strings.Repeat("=", 120))

	for _, exp := range expenses {
		fmt.Printf("%-5d %-12.2f %-25s %-20s %-12s %-10s %-20s\n",
			exp.ID, exp.Amount, truncateString(exp.Description, 25), truncateString(exp.Category, 20),
			exp.Date.Format("2006-01-02"), exp.PaymentMethod, truncateString(exp.Notes, 20))
	}
	fmt.Println(strings.Repeat("=", 120))
	fmt.Printf("Total: %d expense(s)\n", len(expenses))
}

// truncateString truncates a string to the specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// PromptCategory prompts user for category
func PromptCategory() string {
	fmt.Print("Enter category to filter: ")
	return GetUserInput()
}

// PromptDateRange prompts user for date range
func PromptDateRange() (time.Time, time.Time, error) {
	fmt.Print("Enter start date (YYYY-MM-DD): ")
	startStr := GetUserInput()
	startDate, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start date format. Please use YYYY-MM-DD")
	}

	fmt.Print("Enter end date (YYYY-MM-DD): ")
	endStr := GetUserInput()
	endDate, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end date format. Please use YYYY-MM-DD")
	}

	if startDate.After(endDate) {
		return time.Time{}, time.Time{}, fmt.Errorf("start date cannot be after end date")
	}

	return startDate, endDate, nil
}

// PromptSearchTerm prompts user for search term
func PromptSearchTerm() string {
	fmt.Print("Enter search term (searches in description and notes): ")
	return GetUserInput()
}

// PromptExpenseID prompts user for expense ID to delete
func PromptExpenseID() (int, error) {
	fmt.Print("Enter expense ID to delete: ")
	idStr := GetUserInput()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID. Please enter a number")
	}
	return id, nil
}

// DisplayStatistics displays expense statistics
func DisplayStatistics() error {
	// Get monthly budget
	budget, err := GetMonthlyBudget()
	if err != nil {
		return fmt.Errorf("failed to get monthly budget: %v", err)
	}

	// Get current month totals
	now := time.Now()
	currentMonthTotal, err := GetTotalExpensesForMonth(now.Year(), now.Month())
	if err != nil {
		return fmt.Errorf("failed to get current month total: %v", err)
	}

	// Get category totals
	categoryTotals, err := GetCategoryTotals()
	if err != nil {
		return fmt.Errorf("failed to get category totals: %v", err)
	}

	// Get all expenses for total
	allExpenses, err := GetAllExpenses()
	if err != nil {
		return fmt.Errorf("failed to get all expenses: %v", err)
	}

	var totalExpenses float64
	for _, exp := range allExpenses {
		totalExpenses += exp.Amount
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("=== EXPENSE STATISTICS ===")
	fmt.Println(strings.Repeat("=", 60))

	if budget > 0 {
		fmt.Printf("Monthly Budget:        ₹%.2f\n", budget)
		fmt.Printf("Current Month Total:   ₹%.2f\n", currentMonthTotal)
		remaining := budget - currentMonthTotal
		if remaining >= 0 {
			fmt.Printf("Remaining Budget:      ₹%.2f\n", remaining)
		} else {
			fmt.Printf("Over Budget:           ₹%.2f\n", -remaining)
			fmt.Println(strings.Repeat("-", 60))
			DisplayWarning(fmt.Sprintf("You have exceeded your monthly budget by ₹%.2f!", -remaining))
			fmt.Println(strings.Repeat("-", 60))
		}
		if remaining >= 0 {
			fmt.Println(strings.Repeat("-", 60))
		}
	}

	fmt.Printf("Total Expenses (All):  ₹%.2f\n", totalExpenses)
	fmt.Printf("Total Records:         %d\n", len(allExpenses))
	fmt.Println(strings.Repeat("-", 60))

	if len(categoryTotals) > 0 {
		fmt.Println("\nCategory-wise Totals:")
		fmt.Println(strings.Repeat("-", 60))
		for category, total := range categoryTotals {
			fmt.Printf("%-25s ₹%.2f\n", category, total)
		}
	}

	fmt.Println(strings.Repeat("=", 60))
	return nil
}

// DisplayError displays an error message
func DisplayError(err error) {
	fmt.Printf("\nError: %v\n", err)
}

// DisplaySuccess displays a success message
func DisplaySuccess(message string) {
	fmt.Printf("\n✓ %s\n", message)
}

// DisplayWarning displays a warning message
func DisplayWarning(message string) {
	fmt.Printf("\n⚠ WARNING: %s\n", message)
}
