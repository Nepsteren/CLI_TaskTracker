package task

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Nepsteren/CLI_TaskTracker.git/model"
	"github.com/Nepsteren/CLI_TaskTracker.git/output"
)

var path = "tasks.json"
var Tasks []model.Task

func nextID() int {
	max := 0
	for _, t := range Tasks {
		if t.Id > max {
			max = t.Id
		}
	}
	return max + 1
}

func findIndex(id int) int {
	for i, t := range Tasks {
		if t.Id == id {
			return i
		}
	}
	return -1
}

func addTask(param string) {
	if param == "" {
		fmt.Println("укажи описание задачи")
		return
	}
	now := time.Now()
	Tasks = append(Tasks, model.Task{
		Id:          nextID(),
		Description: param,
		Status:      model.StatusUndone,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	WriteFile()
	fmt.Println("задача добавлена")
}

func GetFile() {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		return
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
	if len(Tasks) == 0 {
		fmt.Println("задач нет")
		return
	}
	for _, t := range Tasks {
		fmt.Println(t.Id, t.Description, t.Status)
	}

}

func ListSomeTask(status model.Status) {
	found := false
	for _, t := range Tasks {
		if t.Status == status {
			fmt.Println(t.Id, t.Description, t.Status)
			found = true
		}
	}
	if !found {
		fmt.Println("задач с таким статусом нет")
	}
}

func DeleteTask(param string) {
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("id должно быть числом!")
		return
	}

	idx := findIndex(id)
	if idx == -1 {
		fmt.Println("задача не найдена")
		return
	}
	Tasks = append(Tasks[:idx], Tasks[idx+1:]...)

	WriteFile()

	fmt.Println("задача удалена")
}

func UpdateTask(param string) {
	str := strings.Fields(param)

	if len(str) < 2 {
		fmt.Println("not enough param")
		return
	}

	id, err := strconv.Atoi(str[0])
	if err != nil {
		fmt.Println("id должно быть числом")
		return
	}

	description := strings.Join(str[1:], " ")

	idx := findIndex(id)
	if idx == -1 {
		fmt.Println("задача не найдена")
		return
	}

	Tasks[idx].Description = description
	Tasks[idx].UpdatedAt = time.Now()

	WriteFile()
	fmt.Println("задача обновлена")
}

func MarkTask(param string) {
	fields := strings.Fields(param)

	if len(fields) < 2 {
		fmt.Println("not enough param")
		return
	}

	id, err := strconv.Atoi(fields[0])
	if err != nil {
		fmt.Println("id должно быть числом")
		return
	}

	str := strings.Join(fields[1:], " ")

	status := model.Status(str)
	if !status.Valid() {
		fmt.Println("wrong status")
		return
	}

	idx := findIndex(id)
	if idx == -1 {
		fmt.Println("задача не найдена")
		return
	}

	Tasks[idx].Status = status
	Tasks[idx].UpdatedAt = time.Now()

	WriteFile()
	fmt.Println("статус обновлен")
}

func Start() {
	GetFile()
	output.GreetingOutput()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("task cli > ")
		if !scanner.Scan() {
			return
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
			status := model.Status(param)
			if status.Valid() {
				ListSomeTask(status)
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
