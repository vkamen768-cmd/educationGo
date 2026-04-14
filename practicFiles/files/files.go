package files

import (
	"fmt"
	"os"
)

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")

}

func WriteListToFile(list []string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for _, item := range list {
		_, err = file.WriteString(item + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

func AppendToFile(content string, name string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}
