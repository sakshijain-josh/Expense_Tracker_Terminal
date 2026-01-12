# Expense Tracker Terminal

A simple and efficient command-line expense tracking application written in Go. Track your daily expenses, manage your monthly budget, and analyze your spending patterns all from your terminal.

## Features

- **Add Expenses**: Record expenses with amount, description, category, date, payment method (Cash/UPI), and optional notes
- **View All Expenses**: Display all your expenses in a formatted table
- **Filter by Category**: View expenses filtered by specific categories (e.g., Food, Transport, Entertainment)
- **Filter by Date Range**: View expenses within a specific date range
- **Search Expenses**: Search for expenses by description or notes
- **Delete Expenses**: Remove expenses by ID
- **View Statistics**: Get comprehensive statistics including:
  - Monthly budget tracking
  - Current month total expenses
  - Remaining budget or over-budget warnings
  - Category-wise expense totals
  - Overall expense summary
- **Monthly Budget Management**: Set and track monthly budgets with automatic warnings when exceeded

## Prerequisites

- **Go 1.24.0 or later** - [Download Go](https://golang.org/dl/)
- SQLite database support (automatically handled by the `modernc.org/sqlite` package)

## Installation

1. **Clone or download the repository**
   ```bash
   git clone https://github.com/sakshijain-josh/Expense_Tracker_Terminal.git
   cd Expense_Tracker_Terminal
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   go build
   ```
   This will create an executable file named `expense-tracker` (or `expense-tracker.exe` on Windows).

4. **Run the application**
   ```bash
   ./expense-tracker
   ```
   Or on Windows:
   ```bash
   expense-tracker.exe
   ```

## Usage

### Starting the Application

When you first run the application, you'll be prompted to set your monthly budget:

```
=== Welcome to Expense Tracker ===
Enter your expected spending per month: 5000
✓ Monthly budget set to ₹5000.00
```

### Main Menu

The application provides an interactive menu with the following options:

```
=== Expense Tracker ===
1. Add Expense
2. View All Expenses
3. View Expenses by Category
4. View Expenses by Date Range
5. Search Expenses
6. Delete Expense
7. View Statistics
8. Exit

Select an option:
```

### Adding an Expense

Select option `1` to add a new expense. You'll be prompted for:
- **Amount**: A positive number (e.g., 150.50)
- **Description**: Brief description of the expense
- **Category**: Category name (e.g., Food, Transport, Entertainment)
- **Date**: Date in YYYY-MM-DD format (or press Enter for today's date)
- **Payment Method**: Either "Cash" or "UPI"
- **Notes**: Optional additional notes

The application will warn you if adding the expense would exceed your monthly budget.

### Viewing Expenses

- **Option 2**: View all expenses sorted by date (newest first)
- **Option 3**: Filter expenses by category
- **Option 4**: View expenses within a date range (start date and end date)
- **Option 5**: Search expenses by keyword (searches in description and notes)

### Deleting an Expense

Select option `6` and enter the expense ID to delete it. The expense ID is shown in the expense list.

### Viewing Statistics

Select option `7` to view comprehensive statistics including:
- Monthly budget and current month spending
- Remaining budget or over-budget amount
- Total expenses across all time
- Category-wise expense breakdown

### Example Workflow

```
=== Expense Tracker ===
Select an option: 1

=== Add New Expense ===
Enter amount: 250.00
Enter description: Lunch at restaurant
Enter category (e.g., Food, Transport, Entertainment): Food
Enter date (YYYY-MM-DD) or press Enter for today: 2024-01-15
Enter payment method (Cash/UPI): UPI
Enter notes (optional): Business lunch

✓ Expense added successfully!
```

## Project Structure

```
Expense_Tracker_Terminal/
├── main.go          # Entry point and main menu loop
├── cli.go           # CLI interface, user prompts, and display functions
├── expense.go       # Expense data structure definition
├── database.go      # Database operations (SQLite queries and connections)
├── go.mod           # Go module dependencies
├── go.sum           # Go module checksums
├── expenses.db      # SQLite database file (created automatically)
└── README.md        # This file
```

### File Descriptions

- **`main.go`**: Contains the main function that initializes the database, handles the monthly budget setup, and runs the main menu loop with switch-case logic for each menu option.

- **`cli.go`**: Handles all user interaction including:
  - Menu display
  - User input prompts
  - Expense display formatting
  - Success, error, and warning messages
  - Statistics display

- **`expense.go`**: Defines the `Expense` struct with fields for ID, Amount, Description, Category, Date, PaymentMethod, and Notes.

- **`database.go`**: Manages all database operations:
  - Database initialization and table creation
  - CRUD operations for expenses
  - Monthly budget storage and retrieval
  - Query functions (by category, date range, search)
  - Statistics calculations

## Technologies Used

- **Go (Golang)**: The programming language used for the application
- **SQLite**: Embedded database for data persistence (via `modernc.org/sqlite` package)

## Database

The application uses a SQLite database file (`expenses.db`) that is automatically created in the project directory on first run.

### Database Schema

**expenses table:**
- `id` (INTEGER PRIMARY KEY AUTOINCREMENT)
- `amount` (REAL NOT NULL)
- `description` (TEXT NOT NULL)
- `category` (TEXT NOT NULL)
- `date` (TEXT NOT NULL) - Stored as YYYY-MM-DD format
- `payment_method` (TEXT NOT NULL)
- `notes` (TEXT)

**settings table:**
- `id` (INTEGER PRIMARY KEY AUTOINCREMENT)
- `key` (TEXT UNIQUE NOT NULL)
- `value` (TEXT NOT NULL)

The settings table stores the monthly budget with key `monthly_budget`.

## Example Usage

### Setting Up Monthly Budget
```
=== Welcome to Expense Tracker ===
Enter your expected spending per month: 10000
✓ Monthly budget set to ₹10000.00
```

### Viewing Statistics
```
=== EXPENSE STATISTICS ===
Monthly Budget:        ₹10000.00
Current Month Total:   ₹7500.00
Remaining Budget:      ₹2500.00
Total Expenses (All):  ₹15000.00
Total Records:         25

Category-wise Totals:
Food                   ₹5000.00
Transport              ₹3000.00
Entertainment          ₹2000.00
```

### Budget Warning
If you exceed your monthly budget, you'll see a warning:



## License

This project is open source. Feel free to use, modify, and distribute as needed.
