# Todo List CLI

A simple command-line interface (CLI) application written in Go to manage your to-do list. This application allows you to add, remove, toggle, update, and display your to-dos using various command-line flags.

## Features
- Add a new to-do item
- Remove a to-do item by index
- Toggle the completion status of a to-do item by index
- Update the title of a to-do item by index
- Show the list of to-do items

## Installation
Make sure you have Go installed on your system. Clone this repository and navigate to its directory.

```bash
git clone git@github.com:anjitpariyar/go-to-cli.git
cd go-to-cli
```

## Usage
-Run the program using . The following commands are supported:
-Add a To-Do

```bash
go run ./ --add "This is a todo"
```

-Remove a To-Do
```bash
go run ./ --remove 1
```

-Toggle a To-Do
```bash
go run ./ --toggle 1
```

-Update a To-Do
```bash
go run ./ --update 0:"Buy groceries and snacks"
```
-Show To-Do
```bash
go run ./ --show
```
