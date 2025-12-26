package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"taskmanager/db"
	"taskmanager/handlers"
	"time"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("\nUse: [comand] [arguments]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("\nUse: add [title] [description] [term]")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		deadline, _ := time.Parse("2006-01-02", os.Args[4])
		err := handlers.AddTask(title, description, deadline)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\ntask added successfully!")
	case "list":
		tasks, err := handlers.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
		for _, task := range tasks {
			fmt.Printf("\n%d: %s - %s (Status: %s, term: %s)\n", task.ID, task.Title, task.Description, task.Status, task.Deadline.Format("2006-01-02"))
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("\nUse: update [id] [status]")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		status := os.Args[3]
		err := handlers.UpdateTaskStatus(uint(id), status)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\ntask status updated successfully!")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("\nUse: delete [id]")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := handlers.DeleteTask(uint(id))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\ntask deleted successfully!")
	default:
		fmt.Println("\nCommand not recognized.")
	}
}
