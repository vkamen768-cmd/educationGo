package main

import "practicFiles/files"

func main() {
	//files.WriteFile("Hello, I'm file", "files/file.txt")

	//myList := []string{"Яблоки", "Молоко", "Хлеб", "Кофе"}
	//files.WriteListToFile(myList, "files/shopping.txt")

	files.AppendToFile("Чай", "files/shopping.txt")
}
