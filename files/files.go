package files

import (
	"os"

)

func CreateOnStart() {
	fileName := "tasks.json"
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		os.Create(fileName)
	}

}

