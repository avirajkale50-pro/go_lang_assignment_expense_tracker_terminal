## Expense Tracker 

This repository contains a simple **terminal-based expense tracker** written in Go, backed by a SQLite database.

It is structured as a small learning project to practice:
- Basic Go syntax
- Working with `database/sql` and SQLite
- Building an interactive CLI application

---

## Project structure

- **`main.go`**: Entry point for the CLI app (menu, user input handling).
- **`db.go`**: Database initialization and table creation (`expenses` table in `expenses.db`).
- **`model.go`**: `Expense` struct definition.
- **`go.mod` / `go.sum`**: Go module definition and dependency lock files.
- **`expenses.db`**: SQLite database file created at runtime (if not present).


```bash
go mod tidy
```

---

## How to run

From the project root:

```bash
go run .
```
## Features (CLI)

The terminal app provides a simple menu:

- **1. Add Expense**
  - Prompts for: amount, category, note, and date (`YYYY-MM-DD`)
  - Saves the entry to the SQLite `expenses` table.

- **2. List Expenses**
  - Shows all expenses stored in the database with:
    - ID, amount, category, date, note

- **3. Summary**
  - Displays:
    - Total number of expense records
    - Total amount spent
    - Amount spent per category

- **4. Filter Expenses**
  - Asks for a category
  - Shows expenses only in that category.

- **5. Delete Expense**
  - Asks for an expense ID
  - Deletes the row from the database if it exists.

- **6. Exit**
  - Quits the application.



