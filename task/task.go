package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Nepsteren/CLI_TaskTracker.git/files"
	"github.com/Nepsteren/CLI_TaskTracker.git/output"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var path = "tasks.json"
var Tasks []Task

func addTask(param string) {
	now := time.Now()
	Tasks = append(Tasks, Task{
		Id:          len(Tasks) + 1,
		Description: param,
		Status:      "undone",
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	data, err := json.Marshal(Tasks)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		log.Fatal(err)
	}

}

func GetJson() {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func ListAllTask() {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)

	for _, t := range tasks {
		fmt.Println(t.Id, t.Description, t.Status)
	}

}

func ListSomeTask(param string) {

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	json.Unmarshal(data, &tasks)

	for _, t := range tasks {
		if t.Status == param {
			fmt.Println(t.Id, t.Description, t.Status)
		}
	}

}

func DeleteTask(param string) {
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("id должно быть числом!")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		log.Fatal(err)
	}

	found := -1
	for i, t := range tasks {
		if t.Id == id {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("задача не найдена")
		return
	}

	tasks = append(tasks[:found], tasks[found+1:]...)

	out, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, out, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("задача удалена")
}

func UpdateTask(param string) {
	str := strings.Fields(param)

	if len(str) < 2 {
		fmt.Println("not enough param")
		return
	}

	ind, err := strconv.Atoi(str[0])
	if err != nil {
		log.Fatal(err)
	}
	
	description := strings.Join(str[1:], " ")

	file, err := os.ReadFile(path)
	if err != nil{
		log.Fatal(err)
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == ind {
			tasks[i].Description = description
			fmt.Println("задача обновлена")
		}
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile(path, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func Start() {
	files.CreateOnStart()
	GetJson()
	output.GreetingOutput()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("task cli > ")
		if err := scanner.Scan(); !err {
			log.Fatal(err)
		}

		input := scanner.Text()

		fields := strings.Fields(input)
		if len(fields) == 0 {
			continue
		}

		cmd := fields[0]

		param := strings.Join(fields[1:], " ")

		switch cmd {
		case "add":
			addTask(param)
		case "update":
			UpdateTask(param)
		case "delete":
			DeleteTask(param)
		case "list":
			if param == "done" || param == "undone" || param == "in progress" {
				ListSomeTask(param)
			} else {
				ListAllTask()
			}
		case "list in progress":
		case "mark":
		case "help":
			output.HelpOutput()
		case "exit":
			return
		default:
			output.DefaultOutput(cmd)
		}

	}
}
