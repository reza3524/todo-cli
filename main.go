package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/gocast/todo-cli/server/model"
	"practice/gocast/todo-cli/server/repository"
	"practice/gocast/todo-cli/server/service"
	"strconv"
)

const baseFileLocation = "/Users/reza/go/src/practice/gocast/todo-cli/data/"

var (
	userService     = service.NewUserService(repository.NewFileUserRepo(baseFileLocation + "users.json"))
	categoryService = service.NewCategoryService(repository.NewFileCategoryRepo(baseFileLocation + "categories.json"))
	taskService     = service.NewTaskService(repository.NewFileTaskRepo(baseFileLocation + "tasks.json"))
	loginUser       *model.User
)

func main() {
	fmt.Println("Welcome to TODO CLI App!")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nAvailable commands: register | login | create-category | list-categories | create-task | list-tasks | exit")
		fmt.Print("Enter command: ")
		scanner.Scan()
		command := scanner.Text()
		runCommand(command, scanner)
	}
}

func runCommand(command string, scanner *bufio.Scanner) {
	if command != "register" && command != "login" && loginUser == nil {
		fmt.Println("You must login or register first.")
		return
	}
	switch command {
	case "register":
		handleRegister(scanner)
	case "login":
		handleLogin(scanner)
	case "create-category":
		handleCreateCategory(scanner)
	case "list-categories":
		handleListCategories()
	case "create-task":
		handleCreateTask(scanner)
	case "list-tasks":
		handleListTasks()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Unknown command:", command)
	}
}

func handleRegister(scanner *bufio.Scanner) {
	fmt.Print("Username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()

	user, err := userService.Register(username, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("User: {} registered", user.Username)
}

func handleLogin(scanner *bufio.Scanner) {
	fmt.Print("Username: ")
	scanner.Scan()
	username := scanner.Text()

	fmt.Print("Password: ")
	scanner.Scan()
	password := scanner.Text()

	user, err := userService.Login(username, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}
	loginUser = user
	fmt.Println("Login successful.")
}

func handleCreateCategory(scanner *bufio.Scanner) {
	fmt.Print("Category name: ")
	scanner.Scan()
	name := scanner.Text()

	_, err := categoryService.AddCategory(name, loginUser.Id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Category created.")
}

func handleListCategories() {
	categories, _ := categoryService.ListCategories(loginUser.Id)
	fmt.Println("Your categories:")
	for _, c := range categories {
		fmt.Printf("- %s - [%d] \n", c.Title, c.Id)
	}
}

func handleCreateTask(scanner *bufio.Scanner) {
	fmt.Print("Task title: ")
	scanner.Scan()
	title := scanner.Text()
	handleListCategories()
	fmt.Print("Enter category [Id]: ")
	scanner.Scan()
	categoryId, _ := strconv.Atoi(scanner.Text())
	_, err := taskService.AddTask(title, loginUser.Id, categoryId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task created.")
}

func handleListTasks() {
	tasks, _ := taskService.ListTasks(loginUser.Id)
	fmt.Println("Your tasks:")
	for _, t := range tasks {
		fmt.Printf("- %s - [%v] \n", t.Title, t.Completed)
	}
}
