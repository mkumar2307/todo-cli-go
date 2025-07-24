package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mkumar2307/todo-cli-go/storage"
)

// printUsage prints how to use the CLI
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo add \"task description\"")
	fmt.Println("  todo list")
	fmt.Println("  todo done <task number>")
	fmt.Println("  todo delete <task number>")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// First CLI argument is the command
	command := os.Args[1]
	tasks := storage.LoadTasks() // Load current task list from file

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task description required.")
			return
		}
		desc := os.Args[2]
		tasks = storage.AddTask(tasks, desc)
		storage.SaveTasks(tasks)
		fmt.Println("Task added.")

	case "list":
		storage.PrintTasks(tasks)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task number required.")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number.")
			return
		}
		tasks = storage.MarkDone(tasks, index-1)
		storage.SaveTasks(tasks)
		fmt.Println("Task marked as done.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task number required.")
			return
		}
		index, err := strconv.Atoi(os.Args[2])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number.")
			return
		}
		tasks = storage.DeleteTask(tasks, index-1)
		storage.SaveTasks(tasks)
		fmt.Println("Task deleted.")

	default:
		printUsage()
	}
}
