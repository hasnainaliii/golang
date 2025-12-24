package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aquasecurity/table"
)

type TodoItem struct {
	Title       string     `json:"title"`
	IsCompleted bool       `json:"isCompleted"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

type TodoList []TodoItem



func (list *TodoList) AddTodo(title string) {
	newTodo := TodoItem{
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*list = append(*list, newTodo)
}

func (list *TodoList) ValidateIndex(index int) error {
	if index < 0 || index >= len(*list) {
		return errors.New("invalid index")
	}
	return nil
}

func (list *TodoList) Delete(index int) error {
	if err := list.ValidateIndex(index); err != nil {
		return err
	}

	*list = append((*list)[:index], (*list)[index+1:]...)
	return nil
}

func (list *TodoList) Toggle(index int) error {
	if err := list.ValidateIndex(index); err != nil {
		return err
	}

	todo := &(*list)[index]

	if !todo.IsCompleted {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	todo.IsCompleted = !todo.IsCompleted
	return nil
}

func (list *TodoList) Edit(index int, newTitle string) error {
	// validate index
	if index < 0 || index >= len(*list) {
		return errors.New("invalid index")
	}

	// update title
	(*list)[index].Title = newTitle
	return nil
}


func (list *TodoList) Print() {
	tbl := table.New(os.Stdout)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, todo := range *list {
		completed := "❌"
		if todo.IsCompleted {
			completed = "✅"
		}

		completedAt := ""
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
		}

		createdAt := todo.CreatedAt.Format("2006-01-02 15:04:05")

		tbl.AddRow(
			fmt.Sprintf("%d", index+1),
			todo.Title,
			completed,
			createdAt,
			completedAt,
		)
	}

	tbl.Render()
}
