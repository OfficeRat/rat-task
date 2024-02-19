package helpers

import "fmt"

func CliHelp() {
	fmt.Println("Usage:")
	fmt.Println("-n '<task name>'   | New task")
	fmt.Println("-d '<description>' | Used with -n to create a description")
	fmt.Println("-l                 | Lists all tasks without descriptions")
	fmt.Println("-l <id>            | Lists task with corresponding id with description")
	fmt.Println("-D <id>            | Delete task")
	fmt.Println("-c <id>            | Complete task")
	fmt.Println("-u <id> -t <type>  | updates task by id and type. Available types are: name, description, complete")
}