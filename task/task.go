package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
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
		case "delete":
		case "list":
			if param == "done" || param == "undone" || param == "in progress" {
				ListSomeTask(param)
			} else {
				ListAllTask()
			}

		//case "list done":
		//case "list undone":
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
