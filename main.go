package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker-cli/models"
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
		addTask(service, args)
	case "update":
		updateTask(service, args)
	case "delete":
		deleteTask(service, args)
	case "mark-in-progress":
		markInProgress(service, args)
	case "mark-done":
		markDone(service, args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func addTask(service service.TaskService, args []string) {
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
}

func updateTask(service service.TaskService, args []string) {
	if len(args) < 3 {
		fmt.Println("Id and Task name is not provided")
		os.Exit(1)
	}

	taskId := args[1]
	taskName := args[2]

	num, err := strconv.Atoi(taskId)

	if err != nil {
		fmt.Println("Id must be integer")
		os.Exit(1)
	}

	updatedTask := models.Task{
		Name: taskName,
	}

	if err := service.UpdateTask(num, updatedTask); err != nil {
		fmt.Printf("Failed to add task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task '%s' updated successfully!\n", taskName)
}

func deleteTask(service service.TaskService, args []string) {
	if len(args) < 2 {
		fmt.Println("Id task is not provided")
		os.Exit(1)
	}

	taskId := args[1]

	num, err := strconv.Atoi(taskId)

	if err != nil {
		fmt.Println("Id must be integer")
		os.Exit(1)
	}

	if err := service.DeleteTask(num); err != nil {
		fmt.Printf("Failed to add task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task with Id %s deleted successfully!\n", taskId)
}

func markInProgress(service service.TaskService, args []string) {
	if len(args) < 2 {
		fmt.Println("Id Task is not provided")
		os.Exit(1)
	}

	taskId := args[1]

	num, err := strconv.Atoi(taskId)

	if err != nil {
		fmt.Println("Id must be integer")
		os.Exit(1)
	}

	updatedTask := models.Task{
		Status: models.Inprogress,
	}

	if err := service.UpdateTask(num, updatedTask); err != nil {
		fmt.Printf("Failed to add task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task '%s' status updated successfully!\n", taskId)
}

func markDone(service service.TaskService, args []string) {
	if len(args) < 2 {
		fmt.Println("Id Task is not provided")
		os.Exit(1)
	}

	taskId := args[1]

	num, err := strconv.Atoi(taskId)

	if err != nil {
		fmt.Println("Id must be integer")
		os.Exit(1)
	}

	updatedTask := models.Task{
		Status: models.Done,
	}

	if err := service.UpdateTask(num, updatedTask); err != nil {
		fmt.Printf("Failed to add task: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Task '%s' status updated successfully!\n", taskId)
}
