package main

import (
	"fmt"
	"writeTodo/files"
)

func main() {
	//login := dataUser("Login")
	//password := dataUser("Password")
	//urlUser := dataUser("URL")
	//
	//acc, err := account.NewAccount(login, password, urlUser)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//acc.PrintInfo()

	//files.WriteFile(account.StrInfo(acc), "files/new.txt")
	//files.AppendToFile(account.StrInfo(acc), "files/new.txt")

	//files.ReadFileLineByLine("files/new.txt")
	if err := files.RemoveLine("Login: 1", "files/new.txt"); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Строка удалена")
	}
}

func dataUser(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}
