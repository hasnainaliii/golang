package main

func main() {
	todos := TodoList{}
	storage := NewStorage[TodoList]("todos.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	storage.Save(todos)

}