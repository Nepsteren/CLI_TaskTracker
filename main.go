package main

import "os"

func main(){

	fileName := "tasks.json"
	_, err := os.Stat(fileName)

	if os.IsNotExist(err){
		os.Create(fileName)
	}
	// task1 := task.Task{
	// 	Id: 1,
	// 	Description: "Description",
	// 	Status: "NOT",
	// }

	// bytes, err := json.Marshal(task1)
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bytes))

	// file, err := os.Create("task.json")
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// file.Write(bytes)

	// task2 := task.Task{}

	// data, err := os.ReadFile("task.json")
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// if err = json.Unmarshal(data, &task2); err != nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(data))
}