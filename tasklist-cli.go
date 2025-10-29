package main

import (
	"log"
	"io"
	"os"
	"strconv"
	"strings"
	"fmt"
	"../tasklist-cli/tasks"
	"../tasklist-cli/utils"
)

func main() {
	log.SetOutput(io.Discard)
	path := "db/db.json"
	var task *tasks.TaskList = tasks.InitTasks(path)
	
	
	switch args := os.Args[1:]; args[0] {
	case "add":
		task.AddTask(strings.Join(args[1:], " "),path)
	case "delete":
		id_to_delete, err := strconv.Atoi(args[1])
		utils.WrapErr(err) 
		task.DeleteTask(uint16(id_to_delete), path)
	case "mark-in-progress":
		id_to_update, err := strconv.Atoi(args[1])
		utils.WrapErr(err) 
		task.UpdateStatus(uint16(id_to_update), tasks.InProgress, path)
	case "mark-done":
		id_to_update, err := strconv.Atoi(args[1])
		utils.WrapErr(err) 
		task.UpdateStatus(uint16(id_to_update), tasks.Done, path)
	case "list":
		if len(args) != 1 {
			switch args[1] {
			case "done":
				task.PrintByStatus(tasks.Done)
			case "in-progress":
				task.PrintByStatus(tasks.InProgress)
			default:
				fmt.Println("Unknown status")
			}
		} else {
			task.PrintAll()
		}

	default:
		fmt.Printf("Unknown command")
	}
	
	//t.PrintAll()
}