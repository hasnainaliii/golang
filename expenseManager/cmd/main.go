package main

import (
	"ExpenseManager/internal/expense"
	"fmt"
)

func main() {
	fmt.Println("hi this is my new go program")

    var expenses expense.ExpenseList

		expenses.AddExpense("lunch",50)
		expenses.AddExpense("Car fuel",100)
		expenses.AddExpense("New Shoes",200)

		expenses.PrintExpenses()

	// 	fmt.Println("\n---- UPDATE EXPENSE (ID = 2) ----")
	// err := expenses.UpdateExpense(2, "Internet", 2800, "Billing")
	// if err != nil {
	// 	fmt.Println("Update error:", err)
	// }
	// expenses.PrintExpenses()


	// fmt.Println("\n---- DELETE EXPENSE (ID = 1) ----")
	// err = expenses.DeleteExpense(1, "")
	// if err != nil {
	// 	fmt.Println("Delete error:", err)
	// }

	// expenses.PrintExpenses()
}