package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Nepsteren/CLI_TaskTracker.git/model"
	"github.com/Nepsteren/CLI_TaskTracker.git/output"
	"github.com/Nepsteren/CLI_TaskTracker.git/service"
)

type CLI struct {
	svc *service.Service
}

func New(svc *service.Service) *CLI {
	return &CLI{svc: svc}
}

func (c *CLI) handleAdd(param string) {
	if err := c.svc.Add(param); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("задача добавлена")
}

func (c *CLI) handleDelete(param string) {
	id, ok := parseID(param)
	if !ok {
		return
	}
	if err := c.svc.Delete(id); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("задача удалена")
}

func (c *CLI) handleUpdate(param string) {
	fields := strings.Fields(param)
	if len(fields) < 2 {
		fmt.Println("использование: update <id> <описание>")
		return
	}

	id, ok := parseID(fields[0])
	if !ok {
		return
	}

	desc := strings.Join(fields[1:], " ")
	if err := c.svc.Update(id, desc); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("задача обновлена")
}

func (c *CLI) handleMark(param string) {
	fields := strings.Fields(param)
	if len(fields) < 2 {
		fmt.Println("использование: mark <id> <статус>")
		return
	}

	id, ok := parseID(fields[0])
	if !ok {
		return
	}

	status := model.Status(strings.Join(fields[1:], " "))
	if err := c.svc.Mark(id, status); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("статус обновлён")
}

func (c *CLI) handleList(param string) {
	if param == "" {
		printTasks(c.svc.All())
		return
	}

	status := model.Status(param)
	if !status.Valid() {
		printTasks(c.svc.All())
		return
	}
	printTasks(c.svc.ByStatus(status))
}

func parseID(s string) (int, bool) {
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("id должно быть числом")
		return 0, false
	}
	return id, true
}

func printTasks(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("задач нет")
		return
	}
	for _, t := range tasks {
		fmt.Printf("%d. %s [%s]\n", t.Id, t.Description, t.Status)
	}
}

func (c *CLI) Run() {
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
			c.handleAdd(param)
		case "update":
			c.handleUpdate(param)
		case "delete":
			c.handleDelete(param)
		case "list":
			c.handleList(param)
		case "mark":
			c.handleMark(param)
		case "help":
			output.HelpOutput()
		case "exit":
			return
		default:
			output.DefaultOutput(cmd)
		}

	}
}
