package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add         string
	Remove      int
	Toggle      int
	UpdateTitle string
	Show        bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo, e.g. --add 'complete this'")
	flag.IntVar(&cf.Remove, "remove", -1, "Remove a todo by index, e.g. --remove 1")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the completion of a todo by index, e.g. --toggle 1")
	flag.StringVar(&cf.UpdateTitle, "update", "", "Update the title of a todo by index, e.g. --update 1:new title")
	flag.BoolVar(&cf.Show, "show", false, "Show the todos")
	flag.Parse()
	return &cf
}

func (cf *CommandFlags) Execuate(todos *Todos) {
	switch {
	case cf.Show:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Remove != -1:
		todos.remove(cf.Remove)
	case cf.Toggle != -1:
		todos.toggleCompletion(cf.Toggle)
	case cf.UpdateTitle != "":
		parts := strings.Split(cf.UpdateTitle, ":")
		if len(parts) != 2 {
			panic("Invalid update format. Please use --update index:new title")
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Invalid index")
			os.Exit(1)
		}
		todos.updateTitle(index, parts[1])
	default:
		panic("Invalid command")
	}
}
