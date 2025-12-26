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

func printHelp() {
	fmt.Println(`
Task Manager - CLI para gerenciamento de tarefas

Comandos:
  add [title] [description] [deadline]   Adiciona uma nova tarefa
                                         deadline no formato YYYY-MM-DD
  list                                   Lista todas as tarefas
  update [id] [status]                   Atualiza o status de uma tarefa
                                         status: pendente, em_progresso, concluida
  delete [id]                            Remove uma tarefa
  help                                   Exibe esta mensagem de ajuda`)
}

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "help":
		printHelp()
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("\nUso: add [title] [description] [deadline]")
			fmt.Println("Exemplo: add \"Minha tarefa\" \"Descricao\" \"2024-12-31\"")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		deadline, err := time.Parse("2006-01-02", os.Args[4])
		if err != nil {
			fmt.Println("\nErro: data inválida. Use o formato YYYY-MM-DD (ex: 2024-12-31)")
			return
		}
		err = handlers.AddTask(title, description, deadline)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nTarefa adicionada com sucesso!")
	case "list":
		tasks, err := handlers.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("\nNenhuma tarefa encontrada.")
			return
		}
		fmt.Println()
		for _, task := range tasks {
			fmt.Printf("%d: %s - %s (Status: %s, Prazo: %s)\n", task.ID, task.Title, task.Description, task.Status, task.Deadline.Format("2006-01-02"))
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("\nUso: update [id] [status]")
			fmt.Println("Status válidos: pendente, em_progresso, concluida")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil || id <= 0 {
			fmt.Println("\nErro: ID inválido. Deve ser um número inteiro positivo.")
			return
		}
		status := os.Args[3]
		err = handlers.UpdateTaskStatus(uint(id), status)
		if err != nil {
			fmt.Printf("\nErro: %s\n", err.Error())
			return
		}
		fmt.Println("\nStatus da tarefa atualizado com sucesso!")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("\nUso: delete [id]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil || id <= 0 {
			fmt.Println("\nErro: ID inválido. Deve ser um número inteiro positivo.")
			return
		}
		err = handlers.DeleteTask(uint(id))
		if err != nil {
			fmt.Printf("\nErro: %s\n", err.Error())
			return
		}
		fmt.Println("\nTarefa removida com sucesso!")
	default:
		fmt.Println("\nComando não reconhecido.")
		printHelp()
	}
}
