package main

import "github.com/Nepsteren/CLI_TaskTracker.git/task"

func main() {
	app := task.New("tasks.json")
	app.Run()
}
