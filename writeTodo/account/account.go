package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type Account struct {
	login    string
	password string
	url      string
}

func NewAccount(login, password, urlUser string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID USER")
	}
	_, err := url.ParseRequestURI(urlUser)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}
	res := &Account{
		login:    login,
		password: password,
		url:      urlUser,
	}
	if password == "" {
		res.genPassword(5)
	}
	return res, nil
}

func (acc *Account) genPassword(n int) {
	letter := []rune("1234567890-=asdfghjkl;'zxcvbnm,.")
	res := make([]rune, n)
	for item := range res {
		res[item] = letter[rand.IntN(len(letter))]
	}
	acc.password = string(res)
}

func (acc *Account) Str() string { // Стандартный интерфейс fmt.Stringer
	return fmt.Sprintf("Login: %s | Password: %s | URL: %s",
		acc.login, acc.password, acc.url)
}

func (acc *Account) PrintInfo() {
	fmt.Print(acc.Str())
}

func StrInfo(acc *Account) string {
	return acc.Str()
}
