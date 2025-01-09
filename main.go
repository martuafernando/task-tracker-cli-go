package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker-cli/models"
	"task-tracker-cli/repository"
	"task-tracker-cli/service"
	"task-tracker-cli/storage"
)

func main() {
	args := os.Args[1:]
	fileStorage := storage.FileStorageImpl{
		Filename: "./data/data.json",
	}
	repo := repository.TaskRepositoryImpl{
		Filestorage: &fileStorage,
	}
	taskService := service.TaskService{
		Repository: &repo,
	}

	if len(args) == 0 {
		fmt.Println("No command provided. Available commands: add")
		os.Exit(1)
	}

	switch command := args[0]; command {
	case "add":
		addTask(taskService, args)
	case "update":
		updateTask(taskService, args)
	case "delete":
		deleteTask(taskService, args)
	case "mark-in-progress":
		markInProgress(taskService, args)
	case "mark-done":
		markDone(taskService, args)
	case "list":
		list(taskService, args)
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

func list(service service.TaskService, args []string) {
	if len(args) < 2 {
		tasks := service.GetAllTask(nil)
		for i := range tasks {
			fmt.Printf("%d %s (%s)\n", tasks[i].Id, tasks[i].Name, models.StatusNames[tasks[i].Status])
		}
		return
	}

	taskStatus := 0
	switch args[1] {
	case "todo":
		taskStatus = models.Todo
	case "in-progress":
		taskStatus = models.Inprogress
	case "done":
		taskStatus = models.Done
	}

	tasks := service.GetAllTask(&taskStatus)
	for i := range tasks {
		fmt.Printf("%d %s (%s)\n", tasks[i].Id, tasks[i].Name, models.StatusNames[tasks[i].Status])
	}
}
