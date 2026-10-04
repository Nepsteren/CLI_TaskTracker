package task

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Nepsteren/CLI_TaskTracker.git/output"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	output.GreetingOutput()

	for {
		fmt.Print("task cli > ")
		if err := scanner.Scan(); !err {
			log.Fatal(err)
		}
		cmd := scanner.Text()

		switch cmd {
		case "add":
		case "update":
		case "delete":
		case "list":
		case "list done":
		case "list undone":
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
