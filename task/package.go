package task

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Nepsteren/CLI_TaskTracker.git/model"
	"github.com/Nepsteren/CLI_TaskTracker.git/storage"
)

type App struct {
	store *storage.Storage
	tasks []model.Task
}

func (a *App) load() {
	tasks, err := a.store.Load()
	if err != nil {
		log.Fatal(err)
	}
	a.tasks = tasks
}

func (a *App) save() {
	if err := a.store.Save(a.tasks); err != nil {
		log.Fatal(err)
	}
}

func (a *App) nextID() int {
	max := 0
	for _, t := range a.tasks {
		if t.Id > max {
			max = t.Id
		}
	}
	return max + 1
}

func (a *App) findIndex(id int) int {
	for i, t := range a.tasks {
		if t.Id == id {
			return i
		}
	}
	return -1
}

func (a *App) addTask(param string) {
	if param == "" {
		fmt.Println("укажи описание задачи")
		return
	}
	now := time.Now()
	a.tasks = append(a.tasks, model.Task{
		Id:          a.nextID(),
		Description: param,
		Status:      model.StatusUndone,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	a.save()
	fmt.Println("задача добавлена")
}

func (a *App) listAllTask() {
	if len(a.tasks) == 0 {
		fmt.Println("задач нет")
		return
	}
	for _, t := range a.tasks {
		fmt.Println(t.Id, t.Description, t.Status)
	}

}

func (a *App) listSomeTask(status model.Status) {
	found := false
	for _, t := range a.tasks {
		if t.Status == status {
			fmt.Println(t.Id, t.Description, t.Status)
			found = true
		}
	}
	if !found {
		fmt.Println("задач с таким статусом нет")
	}
}

func (a *App) deleteTask(param string) {
	id, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("id должно быть числом!")
		return
	}

	idx := a.findIndex(id)
	if idx == -1 {
		fmt.Println("задача не найдена")
		return
	}
	a.tasks = append(a.tasks[:idx], a.tasks[idx+1:]...)

	a.save()

	fmt.Println("задача удалена")
}

func (a *App) updateTask(param string) {
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

	idx := a.findIndex(id)
	if idx == -1 {
		fmt.Println("задача не найдена")
		return
	}

	a.tasks[idx].Description = description
	a.tasks[idx].UpdatedAt = time.Now()

	a.save()
	fmt.Println("задача обновлена")
}

func (a *App) markTask(param string) {
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

	idx := a.findIndex(id)
	if idx == -1 {
		fmt.Println("задача не найдена")
		return
	}

	a.tasks[idx].Status = status
	a.tasks[idx].UpdatedAt = time.Now()

	a.save()
	fmt.Println("статус обновлен")
}
