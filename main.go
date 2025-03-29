package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Read(&todos)
	commandFlags := NewCommandFlags()
	commandFlags.Execuate(&todos)
	storage.Save(todos)
}
