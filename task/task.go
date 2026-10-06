package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Nepsteren/CLI_TaskTracker.git/model"
	"github.com/Nepsteren/CLI_TaskTracker.git/output"
	"github.com/Nepsteren/CLI_TaskTracker.git/storage"
)

func New(path string) *App {
	return &App{
		store: storage.New(path),
		tasks: []model.Task{},
	}
}

func (a *App) Run() {
	a.load()
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
			a.addTask(param)
		case "update":
			a.updateTask(param)
		case "delete":
			a.deleteTask(param)
		case "list":
			status := model.Status(param)
			if status.Valid() {
				a.listSomeTask(status)
			} else {
				a.listAllTask()
			}
		case "mark":
			a.markTask(param)
		case "help":
			output.HelpOutput()
		case "exit":
			return
		default:
			output.DefaultOutput(cmd)
		}

	}
}
