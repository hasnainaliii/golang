package expense

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

//*************
// Add function
//*************
func (expenses *ExpenseList) AddExpense(title string , amount float64) {
	if amount < 0 {
		panic("Please Enter a Valid amount")
	}

	if title == "" {
		panic("Expense Title Cannot be empty ")
	}
	NewItemID := len(*expenses) + 1

	newItem:= Expense{
		ID: NewItemID,
		Title: title,
		Category: "",
		Amount: amount,
		Date: time.Now(),


	}
	*expenses = append(*expenses, newItem)
}

//****************
// Delete fucntion
//****************
func (expenses *ExpenseList) DeleteExpense(id int, title string) error {

	if id <= 0 && title == "" {
		return errors.New("Either Title or Id is required ")
	}
  
	for i,e:= range *expenses{


		if(id !=0 && e.ID== id )||(title!="" && e.Title == title) {
			

        *expenses = append((*expenses)[:i], (*expenses)[i+1:]...)
				return  nil

		}
	}
  return  errors.New("Unexpected Error is caused")

}


//*********************
// Edit/Update fucntion
//*********************

func (expenses *ExpenseList) UpdateExpense(id int,title string, amount float64 , category string) error {
     
	if id <= 0 {
		return  errors.New("Provide an item id to delelte ")
	}
    

	 if(title=="")&& (amount <=0) && (category == ""){
            return errors.New("You need to provide either title, amount or category")
	 }
     
	 for i,e := range *expenses{
		  if e.ID == id{

				if(title!=""){
					(*expenses)[i].Title = title
				}
				if amount > 0 {
					(*expenses)[i].Amount=amount
				}
				if category!="" {
					(*expenses)[i].Category = category
				}

				return  nil
			}
	 }

 	 return errors.New("Unexpected Error happened ")
}

//************************************
// List all expenses in table fucntion
//************************************


func (expenses *ExpenseList) PrintExpenses() {
	newTable:=table.NewWriter()
	newTable.SetOutputMirror(os.Stdout)
	
	newTable.AppendHeader(table.Row{"#","Title","Amount","Description","Date"})
  newTable.SetStyle(table.StyleRounded)
	for _, e := range *expenses {
		newTable.AppendRow(table.Row{
			e.ID,
			e.Title,
			fmt.Sprintf("%.2f", e.Amount),
			e.Category,
			e.Date.Format("2006-01-02"),
		})
		newTable.AppendSeparator()
	}
newTable.Render()
	
	
}
