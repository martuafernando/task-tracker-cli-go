package main

import (
	"fmt"
	"os"
	"task-tracker-cli/repository"
	"task-tracker-cli/service"
)

func main() {
	args := os.Args[1:]
	repo := repository.TaskRepository{
		Filename: "./data/data.json",
	}
	service := service.TaskService{
		Repository: &repo,
	}

	if len(args) == 0 {
		fmt.Println("No command provided. Available commands: add")
		os.Exit(1)
	}

	switch command := args[0]; command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Task name is not provided")
			os.Exit(1)
		}

		taskName := args[1]
		if err := service.AddTask(taskName); err != nil {
			fmt.Printf("Failed to add task: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Task '%s' added successfully!\n", taskName)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
