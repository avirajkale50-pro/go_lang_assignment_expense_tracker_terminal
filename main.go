package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("DB Error:", err)
		return
	}
	
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== Personal Expense Tracker =====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. List Expenses")
		fmt.Println("3. Summary")
		fmt.Println("4. Filter Expenses")
		fmt.Println("5. Delete Expense")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addExpense(reader)
		case 2:
			listExpenses()
		case 3:
			showSummary()
		case 4:
			filterByCategory(reader)
		case 5:
			deleteExpense()
		case 6:
			fmt.Println("Goodbye 👋")
			return
		default:
			fmt.Println("Invalid choice, try again")
		}

		// pause before showing menu again
		fmt.Print("\nPress Enter to continue...")
		reader.ReadString('\n')
	}
}


func addExpense(reader *bufio.Reader) {
	var amount float64
	var category, note, date string

	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)

	fmt.Print("Enter category: ")
	fmt.Scanln(&category)

	fmt.Print("Enter note: ")
	reader.ReadString('\n')
	note, _ = reader.ReadString('\n')

	fmt.Print("Enter date (YYYY-MM-DD): ")
	fmt.Scanln(&date)

	query := `INSERT INTO expenses(amount, category, note, date) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, amount, category, note[:len(note)-1], date)

	if err != nil {
		fmt.Println("Error saving expense:", err)
		return
	}

	fmt.Println("✅ Expense saved to database!")
}


func listExpenses() {
	rows, err := db.Query("SELECT id, amount, category, note, date FROM expenses")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\nID | Amount | Category | Date | Note")
	fmt.Println("-------------------------------------------")

	found := false

	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.ID, &e.Amount, &e.Category, &e.Note, &e.Date); err != nil {
			fmt.Println("Error reading row:", err)
			return
		}

		found = true

		fmt.Printf("%d | %.2f | %s | %s | %s\n",
			e.ID, e.Amount, e.Category, e.Date, e.Note)
	}

	if !found {
		fmt.Println("No expenses found.")
	}
}


func deleteExpense() {
	var id int
	fmt.Print("Enter Expense ID to delete: ")
	fmt.Scanln(&id)

	result, err := db.Exec("DELETE FROM expenses WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error deleting expense:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error checking delete result:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("❌ Expense ID not found.")
	} else {
		fmt.Println("✅ Expense deleted successfully!")
	}
}

func filterByCategory(reader *bufio.Reader) {
	var category string
	fmt.Print("Enter category to filter: ")
	fmt.Scanln(&category)

	found := false

	fmt.Println("\nID | Amount | Category | Date | Note")
	fmt.Println("-------------------------------------------")

	rows, err := db.Query("SELECT id, amount, category, note, date FROM expenses WHERE category = ?", category)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.ID, &e.Amount, &e.Category, &e.Note, &e.Date); err != nil {
			fmt.Println("Error reading row:", err)
			return
		}

		fmt.Printf("%d | %.2f | %s | %s | %s\n",
			e.ID, e.Amount, e.Category, e.Date, e.Note)
		found = true
	}

	if !found {
		fmt.Println("No expenses found for this category.")
	}
}

func showSummary() {
	var totalCount int
	var totalAmount float64

	err := db.QueryRow("SELECT COUNT(*), COALESCE(SUM(amount), 0) FROM expenses").
		Scan(&totalCount, &totalAmount)
	if err != nil {
		fmt.Println("Error fetching summary:", err)
		return
	}

	fmt.Println("\n===== Expense Summary =====")
	fmt.Printf("Total entries: %d\n", totalCount)
	fmt.Printf("Total amount: %.2f\n\n", totalAmount)

	rows, err := db.Query("SELECT category, COALESCE(SUM(amount), 0) FROM expenses GROUP BY category")
	if err != nil {
		fmt.Println("Error fetching category summary:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Amount by category:")
	for rows.Next() {
		var category string
		var catAmount float64
		if err := rows.Scan(&category, &catAmount); err != nil {
			fmt.Println("Error reading row:", err)
			return
		}
		fmt.Printf("- %s: %.2f\n", category, catAmount)
	}
}
