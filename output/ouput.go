package output

import "fmt"

func GreetingOutput() {
	fmt.Println("Привет, это твой трекер задач!!!")
	fmt.Println("Если ты здесь впервые команда 'help' тебе поможет!")
}

func DefaultOutput(cmd string) {
	fmt.Println(cmd, "- unknown command")
	fmt.Println("Run 'help' for usage")
}

func HelpOutput() {
	fmt.Println("'add' - add new task (string)")
	fmt.Println("'update' - update task (id, string)")
	fmt.Println("'delete' - delete task by id (int)")
	fmt.Println("'mark' - mark task (in progress | done)")
	fmt.Println("'list' - list all task")
	fmt.Println("'list done' - list all done task")
	fmt.Println("'list undone' - list all undone task")
	fmt.Println("'list in progress' - list all task in progress")
}
