package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteFile(content, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error for create file:", err)
		return
	}

	defer file.Close()
	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Error for write:")
		return
	}
	fmt.Println("\nwrite ok")
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

func ReadFile(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error for read:", err)
		return
	}
	fmt.Println(string(data))
}

func ReadFileLineByLine(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Прочитано:", line)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}

func RemoveLine(contentDel string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), contentDel) {
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	output := strings.Join(lines, "\n")
	return os.WriteFile(filename, []byte(output+"\n"), 0644)
}
