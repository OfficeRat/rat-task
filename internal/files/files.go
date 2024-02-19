package files

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"officerat/ratTask/internal/helpers"
	"os"
	"path/filepath"
	"strings"
)

var tasksLocation string

func getTasksLocation() {
	homeDir := os.Getenv("HOME")
	if homeDir != "" {
		tasksLocation = filepath.Join(homeDir, ".ratTasks", "tasks.yaml")
	} else {
		tasksLocation = filepath.Join(os.Getenv("USERPROFILE"), ".ratTasks", "tasks.yaml")
	}
}

type Task struct {
	ID          int    `yaml:"id"`
	Completed   bool   `yaml:"completed"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Tasks struct {
	Tasks []Task `yaml:"tasks"`
}

var tasks Tasks

func readTasksFromFile() error {
	getTasksLocation()
	data, err := os.ReadFile(tasksLocation)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = yaml.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	return nil
}

func writeTasksToFile() error {
	data, err := yaml.Marshal(&tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(tasksLocation, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ensureDirectoryAndFile() error {
	getTasksLocation()
	dir := filepath.Dir(tasksLocation)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if _, err := os.Stat(tasksLocation); os.IsNotExist(err) {
		if _, err := os.Create(tasksLocation); err != nil {
			return err
		}
	}

	return nil
}

func NewTask(name, description string) {
	if err := ensureDirectoryAndFile(); err != nil {
		fmt.Println("Error creating directory or file:", err)
		return
	}

	if err := readTasksFromFile(); err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	newTask := Task{
		ID:          tasks.Tasks[len(tasks.Tasks)-1].ID + 1,
		Completed:   false,
		Name:        name,
		Description: description,
	}

	tasks.Tasks = append(tasks.Tasks, newTask)

	if err := writeTasksToFile(); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Task name:", newTask.Name)
	if len(newTask.Description) == 0 {
		return
	}
	fmt.Println("Description:", newTask.Description)
}

func ListTasks(ID int) {
	if err := readTasksFromFile(); err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	if ID != 0 {
		for _, task := range tasks.Tasks {
			if task.ID == ID {
				fmt.Println("------------------------------------------")
				fmt.Printf("Completed: %s\n", func() string {
					if task.Completed {
						return "yes"
					}
					return "no"
				}())
				fmt.Printf("Name: %s\n", task.Name)
				fmt.Printf("Description: %s\n", task.Description)
				fmt.Println("------------------------------------------")
				return
			}

		}
		fmt.Println("Task not found.")
	} else {
		fmt.Printf("All Tasks: \n\n")
		for _, task := range tasks.Tasks {
			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Completed: %s\n", func() string {
				if task.Completed {
					return "yes"
				}
				return "no"
			}())
			fmt.Printf("Name: %s\n", task.Name)
			fmt.Println("------------------------------------------")
		}
	}
}

func DeleteTask(ID int) {
	if err := readTasksFromFile(); err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	if ID != 0 {
		for i, task := range tasks.Tasks {
			if task.ID == ID {
				tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)

				if err := writeTasksToFile(); err != nil {
					fmt.Println("Error writing tasks:", err)
					return
				}
				return
			}

		}
		fmt.Println("Task not found.")
	}
}

func CompleteTask(ID int) {
	if err := readTasksFromFile(); err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	var taskIndex int
	for i, task := range tasks.Tasks {
		if task.ID == ID {
			taskIndex = i
			break
		}
	}

	task := tasks.Tasks[taskIndex]

	if task.Completed {
		fmt.Println("Task with given ID does not exist")
	}

	tasks.Tasks[taskIndex].Completed = true

	fmt.Println("------------------------------------------")
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Completed: %s\n", func() string {
		if tasks.Tasks[taskIndex].Completed {
			return "yes"
		}
		return "no"
	}())
	fmt.Printf("Name: %s\n", task.Name)
	fmt.Println("------------------------------------------")

	// Write the updated task list to file
	if err := writeTasksToFile(); err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

}

func UpdateTask(ID int, updateType string) {
	if err := readTasksFromFile(); err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	var taskIndex int
	for i, task := range tasks.Tasks {
		if task.ID == ID {
			taskIndex = i
			break
		}
	}
	task := tasks.Tasks[taskIndex]

	switch updateType {
	case "name":

		fmt.Print("New task name: ")
		newName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		newName = strings.TrimSpace(newName)
		tasks.Tasks[taskIndex].Name = newName
		fmt.Println("------------------------------------------")
		fmt.Printf("Completed: %s\n", func() string {
			if task.Completed {
				return "yes"
			}
			return "no"
		}())
		fmt.Printf("Name: %s\n", tasks.Tasks[taskIndex].Name)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Println("------------------------------------------")

		if err := writeTasksToFile(); err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}
		return
	case "description":
		
		fmt.Print("New task description: ")
		newDesc, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		newDesc = strings.TrimSpace(newDesc)
		tasks.Tasks[taskIndex].Description = newDesc
		fmt.Println("------------------------------------------")
		fmt.Printf("Completed: %s\n", func() string {
			if task.Completed {
				return "yes"
			}
			return "no"
		}())
		fmt.Printf("Name: %s\n", task.Name)
		fmt.Printf("Description: %s\n", tasks.Tasks[taskIndex].Description)
		fmt.Println("------------------------------------------")

		if err := writeTasksToFile(); err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}
		return
	case "complete":

		tasks.Tasks[taskIndex].Completed = !task.Completed
		fmt.Println(tasks.Tasks[taskIndex].Completed)
		fmt.Println("------------------------------------------")
		fmt.Printf("Completed: %s\n", func() string {
			if tasks.Tasks[taskIndex].Completed {
				return "yes"
			}
			return "no"
		}())
		fmt.Printf("Name: %s\n", task.Name)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Println("------------------------------------------")
		
		if err := writeTasksToFile(); err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}
		return
	default:
		helpers.CliHelp()
	}

}
