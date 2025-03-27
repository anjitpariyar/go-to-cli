package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Read(&todos)
	todos.add("Learn go")
	todos.add("why so error")
	todos.toggleCompletion(1)
	todos.updateTitle(1, "why so error agian?")
	todos.print()
	storage.Save(todos)
}
