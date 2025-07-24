package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const taskFileName = ".todo/tasks.json" // File to store tasks

// getTaskFilePath builds the path to the task file inside the user's home directory
func getTaskFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not get user home directory")
	}
	return filepath.Join(home, taskFileName)
}

// LoadTasks reads the task list from the file, or returns an empty list
func LoadTasks() []Task {
	filePath := getTaskFilePath()
	var tasks []Task

	data, err := os.ReadFile(filePath)
	if err == nil {
		_ = json.Unmarshal(data, &tasks)
	}

	return tasks
}

// SaveTasks writes the current task list to the file in JSON format
func SaveTasks(tasks []Task) {
	filePath := getTaskFilePath()
	os.MkdirAll(filepath.Dir(filePath), os.ModePerm) // Ensure ~/.todo directory exists
	data, _ := json.MarshalIndent(tasks, "", "  ")   // Pretty-print
	_ = os.WriteFile(filePath, data, 0644)           // Save with user read/write permissions
}

// AddTask creates a new task and appends it to the list
func AddTask(tasks []Task, desc string) []Task {
	newTask := Task{Description: desc, Done: false}
	return append(tasks, newTask)
}

// MarkDone marks the task at the given index as done
func MarkDone(tasks []Task, index int) []Task {
	if index >= 0 && index < len(tasks) {
		tasks[index].Done = true
	}
	return tasks
}

// DeleteTask removes the task at the given index
func DeleteTask(tasks []Task, index int) []Task {
	if index >= 0 && index < len(tasks) {
		return append(tasks[:index], tasks[index+1:]...)
	}
	return tasks
}

// PrintTasks displays all tasks with their status and index
func PrintTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i, t := range tasks {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, t.Description)
	}
}
