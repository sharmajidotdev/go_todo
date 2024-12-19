package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/sharmajidotdev/go_todo/internal/todo_manager"
)

func clearScreen() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func main() {

	manager := todo_manager.New()
	choice := 0
	msg := `Welcome to TODO app.

	1. Show all items
	2. Add na item
	3. Update an item
	4. Delete an item
	5. Delete all items
	6. Exit

	Choose an option: 
	`
	for choice != 6 {
		clearScreen()
		fmt.Println(msg)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			manager.DisplayTodos()
		case 2:
			fmt.Printf("Enter id for todo:")
			var id int
			fmt.Scan(&id)
			fmt.Printf("Enter todo:")
			var description string
			fmt.Scan(&description)
			manager.AddTodo(id, description)
		case 3:
			fmt.Printf("Enter id for todo:")
			var id int
			fmt.Scan(&id)
			fmt.Printf("Enter new todo:")
			var description string
			fmt.Scan(&description)
			manager.UpdateTodo(id, description)
		case 4:
			fmt.Printf("Enter id for todo:")
			var id int
			fmt.Scan(&id)
			manager.DeleteTodo(id)
		case 5:
			manager.DeleteAllTodo()
		case 6:
			fmt.Println("OK Bye!")
		default:
			fmt.Println("Invalid Choice")
		}
		fmt.Println("Press enter to continue...")
		fmt.Scanln()
	}

}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
