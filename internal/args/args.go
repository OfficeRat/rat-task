package Args

import (
	"officerat/ratTask/internal/files"
	"officerat/ratTask/internal/helpers"
	"reflect"
	"strconv"
)

func IsInt(v interface{}) bool {
    return reflect.TypeOf(v).Kind() == reflect.Int
}

func StringToInt(s string) (int) {
    intValue, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return intValue
}

func ArgHandler(args []string) {
	
	var newTask bool;
	var taskName string;
	var description bool;
	var descriptionText string;
	var listTasks bool;
	var delete bool;
	var complete bool;
	var update bool;
	var updateType string;
	var taskID int;



	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-n":
			newTask = true
			if i+1 < len(args) {
				taskName = args[i+1]
			} else {
				helpers.CliHelp()
			}
        case "-d":
			description = true
			if i+1 < len(args) {
				descriptionText = args[i+1]
			} else {
				helpers.CliHelp()
				return
			}
		case "-l":
			listTasks = true
			if i+1 < len(args) {
				taskID = StringToInt(args[i+1])
			} else {
				taskID = 0
			}
		case "-D":
			delete = true
			if i+1 < len(args) {
				taskID = StringToInt(args[i+1])
			} else {
				taskID = 0
			}
		case "-c":
			complete = true
			if i+1 < len(args) {
				taskID = StringToInt(args[i+1])
			} else {
				taskID = 0
			}
		case "-u":
			update = true;
			if i+1 < len(args) {
				taskID = StringToInt(args[i+1])
			} else {
				helpers.CliHelp()
				return
			}
			if i+3 < len(args) {
				if args[2] == "-t" {
					updateType = args[3]
				}
				
			} else {
				helpers.CliHelp()
				return
			}
		}
		
	}

	if len(args) == 0 {
		helpers.CliHelp()
		return
	}

	if description && !newTask  {
		helpers.CliHelp()
		return
	}

	if newTask {
		if description{
			files.NewTask(taskName, descriptionText)
		} else {
			files.NewTask(taskName, descriptionText)
		}
	}

	if listTasks {
		files.ListTasks(taskID)
	}

	if delete {
		files.DeleteTask(taskID)
	}

	if complete {
		files.CompleteTask(taskID)
	}
	
	if update {
		files.UpdateTask(taskID, updateType)
	}

}

