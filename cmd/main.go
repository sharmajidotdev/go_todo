package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
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
