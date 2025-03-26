package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) remove(index int) error {
	t := *todos

	err := t.validateIndex(index)

	if err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil

}

// toggle the completion
func (todos *Todos) toggleCompletion(index int) error {
	t := *todos
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	// assign index value to variable
	isCompleted := t[index].IsCompleted
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].IsCompleted = !isCompleted
	return nil
}

// update the title
func (todos *Todos) updateTitle(index int, title string) error {
	t := *todos
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

// print all todos in a table
func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for i, todo := range *todos {
		completed := "No"
		if todo.IsCompleted {
			completed = "Yes"
		}
		completedAt := ""
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
		}
		table.AddRow(strconv.Itoa(i+1), todo.Title, completed, todo.CreatedAt.Format("2006-01-02 15:04:05"), completedAt)
	}
	table.Render()
}

// helper function
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}
