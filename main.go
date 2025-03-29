package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID                 int
	Title, Description string
	IsCompleted        bool
}

type TaskList struct {
	Tasks  []Task
	nextId int
}

func (t *TaskList) AddTask(title, description string) {
	newTask := Task{
		ID:          t.nextId + 1,
		Title:       title,
		Description: description,
		IsCompleted: false,
	}
	t.nextId++
	t.Tasks = append(t.Tasks, newTask)
}

func (t *TaskList) RemoveTask(id int) bool {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (t *TaskList) UpdateTask(id int, title, description string) bool {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Title = title
			t.Tasks[i].Description = description
			return true
		}
	}
	return false
}

func (t *TaskList) MarkAsDone(id int) bool {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].IsCompleted = true
			return true
		}
	}
	return false
}

func (t *TaskList) ListTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}
	for _, task := range t.Tasks {
		fmt.Println("Id of task:", task.ID)
		fmt.Println("Title of task:", task.Title)
		fmt.Println("Description of task:", task.Description)
		if task.IsCompleted {
			fmt.Println("Completed!")
		} else {
			fmt.Println("Not completed")
		}
		fmt.Println()
	}
}

func help() {
	fmt.Println("add - adds a new task")
	fmt.Println("remove - removes a task by its id")
	fmt.Println("update - updates a tasks title and description by its id")
	fmt.Println("mark - marks a task done by its id")
	fmt.Println("list - lists all tasks")
	fmt.Println("help - lists all commands")
	fmt.Println("quit - closes the application")
}

func main() {
	var tasksList TaskList
	input := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to task manager 1.0!\nTo create a new task write 'add'")
	for {
		inputString, _ := input.ReadString('\n')
		inputString = strings.TrimSpace(inputString)
		switch inputString {
		case "add":
			fmt.Print("Enter the title of new task: ")
			title, _ := input.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter the description of new task: ")
			description, _ := input.ReadString('\n')
			description = strings.TrimSpace(description)
			tasksList.AddTask(title, description)
			fmt.Println("Task added successfully!")
		case "remove":
			var id int
			fmt.Print("Enter the id of task you want to remove: ")
			fmt.Scanln(&id)
			flag := tasksList.RemoveTask(id)
			if flag {
				fmt.Println("Task removed successfully!")
			} else {
				fmt.Println("Error! Task not found!")
			}
		case "update":
			var id int
			fmt.Print("Enter the id of task you want to update: ")
			fmt.Scanln(&id)
			fmt.Print("Enter new title of this task: ")
			title, _ := input.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter new desciption of this task: ")
			description, _ := input.ReadString('\n')
			description = strings.TrimSpace(description)
			flag := tasksList.UpdateTask(id, title, description)
			if flag {
				fmt.Println("Task updated successfully!")
			} else {
				fmt.Println("Error! Task not found!")
			}
		case "mark":
			var id int
			fmt.Print("Enter the id of task you want to mark as done: ")
			fmt.Scanln(&id)
			flag := tasksList.MarkAsDone(id)
			if flag {
				fmt.Println("Task marked as done successfully!")
			} else {
				fmt.Println("Error! Task not found!")
			}
		case "list":
			tasksList.ListTasks()
		case "help":
			help()
		case "quit":
			os.Exit(1)
		default:
			fmt.Println("No command found!\nFor help write command 'help'")
		}
	}
}
