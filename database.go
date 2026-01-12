package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// InitDB initializes the database connection and creates tables if they don't exist
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite", "./expenses.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Create expenses table
	createExpensesTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		description TEXT NOT NULL,
		category TEXT NOT NULL,
		date TEXT NOT NULL,
		payment_method TEXT NOT NULL,
		notes TEXT
	);`

	_, err = db.Exec(createExpensesTable)
	if err != nil {
		return fmt.Errorf("failed to create expenses table: %v", err)
	}

	// Create settings table
	createSettingsTable := `
	CREATE TABLE IF NOT EXISTS settings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT UNIQUE NOT NULL,
		value TEXT NOT NULL
	);`

	_, err = db.Exec(createSettingsTable)
	if err != nil {
		return fmt.Errorf("failed to create settings table: %v", err)
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// GetMonthlyBudget retrieves the monthly budget from settings
func GetMonthlyBudget() (float64, error) {
	var budget float64
	query := "SELECT value FROM settings WHERE key = 'monthly_budget'"
	err := db.QueryRow(query).Scan(&budget)
	if err == sql.ErrNoRows {
		return 0, nil // Budget not set yet
	}
	if err != nil {
		return 0, fmt.Errorf("failed to get monthly budget: %v", err)
	}
	return budget, nil
}

// SetMonthlyBudget sets the monthly budget in settings
func SetMonthlyBudget(budget float64) error {
	query := `
	INSERT INTO settings (key, value) 
	VALUES ('monthly_budget', ?)
	ON CONFLICT(key) DO UPDATE SET value = excluded.value`

	_, err := db.Exec(query, budget)
	if err != nil {
		return fmt.Errorf("failed to set monthly budget: %v", err)
	}
	return nil
}

// AddExpense adds a new expense to the database
func AddExpense(expense Expense) error {
	query := `
	INSERT INTO expenses (amount, description, category, date, payment_method, notes)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query, expense.Amount, expense.Description, expense.Category,
		expense.Date.Format("2006-01-02"), expense.PaymentMethod, expense.Notes)
	if err != nil {
		return fmt.Errorf("failed to add expense: %v", err)
	}
	return nil
}

// GetAllExpenses retrieves all expenses from the database
func GetAllExpenses() ([]Expense, error) {
	query := "SELECT id, amount, description, category, date, payment_method, notes FROM expenses ORDER BY date DESC, id DESC"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query expenses: %v", err)
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var exp Expense
		var dateStr string
		err := rows.Scan(&exp.ID, &exp.Amount, &exp.Description, &exp.Category,
			&dateStr, &exp.PaymentMethod, &exp.Notes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %v", err)
		}
		exp.Date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %v", err)
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

// GetExpensesByCategory retrieves expenses filtered by category
func GetExpensesByCategory(category string) ([]Expense, error) {
	query := "SELECT id, amount, description, category, date, payment_method, notes FROM expenses WHERE category = ? ORDER BY date DESC, id DESC"
	rows, err := db.Query(query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to query expenses by category: %v", err)
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var exp Expense
		var dateStr string
		err := rows.Scan(&exp.ID, &exp.Amount, &exp.Description, &exp.Category,
			&dateStr, &exp.PaymentMethod, &exp.Notes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %v", err)
		}
		exp.Date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %v", err)
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

// GetExpensesByDateRange retrieves expenses within a date range
func GetExpensesByDateRange(startDate, endDate time.Time) ([]Expense, error) {
	query := "SELECT id, amount, description, category, date, payment_method, notes FROM expenses WHERE date >= ? AND date <= ? ORDER BY date DESC, id DESC"
	rows, err := db.Query(query, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("failed to query expenses by date range: %v", err)
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var exp Expense
		var dateStr string
		err := rows.Scan(&exp.ID, &exp.Amount, &exp.Description, &exp.Category,
			&dateStr, &exp.PaymentMethod, &exp.Notes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %v", err)
		}
		exp.Date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %v", err)
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

// SearchExpenses searches expenses by description or notes
func SearchExpenses(searchTerm string) ([]Expense, error) {
	query := "SELECT id, amount, description, category, date, payment_method, notes FROM expenses WHERE description LIKE ? OR notes LIKE ? ORDER BY date DESC, id DESC"
	searchPattern := "%" + searchTerm + "%"
	rows, err := db.Query(query, searchPattern, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("failed to search expenses: %v", err)
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var exp Expense
		var dateStr string
		err := rows.Scan(&exp.ID, &exp.Amount, &exp.Description, &exp.Category,
			&dateStr, &exp.PaymentMethod, &exp.Notes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %v", err)
		}
		exp.Date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %v", err)
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

// DeleteExpense deletes an expense by ID
func DeleteExpense(id int) error {
	query := "DELETE FROM expenses WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete expense: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("expense with ID %d not found", id)
	}
	return nil
}

// GetTotalExpensesForMonth calculates total expenses for a given month
func GetTotalExpensesForMonth(year int, month time.Month) (float64, error) {
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0).AddDate(0, 0, -1)

	query := "SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE date >= ? AND date <= ?"
	var total float64
	err := db.QueryRow(query, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total expenses: %v", err)
	}
	return total, nil
}

// GetCategoryTotals calculates total expenses grouped by category
func GetCategoryTotals() (map[string]float64, error) {
	query := "SELECT category, SUM(amount) FROM expenses GROUP BY category"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query category totals: %v", err)
	}
	defer rows.Close()

	totals := make(map[string]float64)
	for rows.Next() {
		var category string
		var total float64
		err := rows.Scan(&category, &total)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category total: %v", err)
		}
		totals[category] = total
	}
	return totals, nil
}
