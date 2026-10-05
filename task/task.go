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

	WriteFile()

}

func GetFile() {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteFile() {
	out, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, out, 0644); err != nil {
		log.Fatal(err)
	}
}

func ListAllTask() {
	GetFile()
	for _, t := range Tasks {
		fmt.Println(t.Id, t.Description, t.Status)
	}

}

func ListSomeTask(param string) {
	GetFile()
	for _, t := range Tasks {
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

	GetFile()

	found := -1
	for i, t := range Tasks {
		if t.Id == id {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("задача не найдена")
		return
	}

	Tasks = append(Tasks[:found], Tasks[found+1:]...)

	WriteFile()

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

	GetFile()

	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].Id == ind {
			Tasks[i].Description = description
			fmt.Println("задача обновлена")
		}
	}

	WriteFile()
}

func MarkTask(param string) {
	fields := strings.Fields(param)

	if len(fields) < 2 {
		fmt.Println("not enough param")
		return
	}

	ind, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Join(fields[1:], " ")

	GetFile()

	if str == "done" || str == "undone" || str == "in progress" {
		for i := 0; i < len(Tasks); i++ {
			if Tasks[i].Id == ind {
				Tasks[i].Status = str
				fmt.Println("статус обновлен")
			}
		}
	} else {
		fmt.Println("wrong status")
	}

	WriteFile()
}

func Start() {
	files.CreateOnStart()
	GetFile()
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
		case "mark":
			MarkTask(param)
		case "help":
			output.HelpOutput()
		case "exit":
			return
		default:
			output.DefaultOutput(cmd)
		}

	}
}
