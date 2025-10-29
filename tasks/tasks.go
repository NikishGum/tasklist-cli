package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"../utils"
)

type TaskStatus string
const (
	ToDo TaskStatus = "To do"
	InProgress TaskStatus = "In progress"
	Done TaskStatus = "Done"
)

type Task struct {
	Description string
	Status TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskList struct {
	TaskList map[uint16]Task
	LatestId uint16
}

func (t *TaskList) AddTask (desc string, path string) uint16 {
	id := t.LatestId + 1

	task_to_save := Task {
		desc,
		ToDo,
		time.Now(),
		time.Time{},
	}

	t.TaskList[id] = task_to_save
	
	t.LatestId = id
	t.rewriteJSON(path)
	return id
}

func (t *TaskList) DeleteTask (id uint16, path string) {
	delete(t.TaskList, id)
	log.Println("Now size: ", len(t.TaskList))
	t.rewriteJSON(path)
}

func (t *TaskList) UpdateStatus (id uint16, status TaskStatus, path string) {
	// Change status and field updatedAt
	temp := t.TaskList[id] 
	temp.Status = status 
	temp.UpdatedAt = time.Now()
	t.TaskList[id] = temp
	t.rewriteJSON(path)
}	

func (t *TaskList) rewriteJSON(path string) {
	to_save := *t
	
	buf, err := json.Marshal(to_save)
	utils.WrapErr(err)
	
	err = os.WriteFile(path, buf, 0655)
	utils.WrapErr(err)
	
	log.Printf("Rewritten successfully")
}


func InitTasks(filePath string) *TaskList {
	// Read from file to TaskList
	// Open file 
	fp, err := os.OpenFile(filePath, os.O_RDONLY | os.O_CREATE, 0666)
	fi, err1 := os.Stat(filePath)

	defer func () {
		if err := fp.Close(); err != nil {
			panic (err)
		}
	}()	
	utils.WrapErr(err)
	utils.WrapErr(err1)
	
	var new_tl = TaskList{}
	
	buf := make([]byte, fi.Size())

	n, err := fp.Read(buf)
	utils.WrapErr(err)

	log.Printf("Read %d bytes, got %s", n, string(buf))

	if fi.Size() == 0 {
		return &TaskList{make(map[uint16]Task, 0), 0}
	}

	err = json.Unmarshal(buf, &new_tl)
	utils.WrapErr(err)

    log.Printf("Unmarshalled to %d objects", len(new_tl.TaskList))

	return &new_tl
}

// Useless
func (t *TaskList) getListTasks() map[uint16]Task {
	return t.TaskList
}

func (t *TaskList) PrintAll() {
	var allTasks = t.getListTasks()
	for i:=1; i <= int(t.LatestId); i+= 1 {
		if len(allTasks[uint16(i)].Description) == 0 { continue }
		fmt.Println("Task #", i)
		fmt.Println("Status: ", allTasks[uint16(i)].Status)
		fmt.Println("Description: ", allTasks[uint16(i)].Description)
		fmt.Println("Created at: ", allTasks[uint16(i)].CreatedAt.Format(time.RFC822))
		if !allTasks[uint16(i)].UpdatedAt.IsZero() { 
			fmt.Println("Last updated at: ", allTasks[uint16(i)].UpdatedAt.Format(time.RFC822))
		} else {
			fmt.Println("Just added")
		}
	} 
}

func (t *TaskList) PrintByStatus(status TaskStatus) {
	var allTasks = t.getListTasks()
	for i:=1; i <= int(t.LatestId); i+= 1 {
		if len(allTasks[uint16(i)].Description) == 0 || allTasks[uint16(i)].Status != status { continue }
		fmt.Println("Task #", i)
		fmt.Println("Status: ", allTasks[uint16(i)].Status)
		fmt.Println("Description: ", allTasks[uint16(i)].Description)
		fmt.Println("Created at: ", allTasks[uint16(i)].CreatedAt.Format(time.RFC822))
		if !allTasks[uint16(i)].UpdatedAt.IsZero() { 
			fmt.Println("Last updated at: ", allTasks[uint16(i)].UpdatedAt.Format(time.RFC822))
		} else {
			fmt.Println("Just added")
		}
	} 
}