package main //github.com/hyun6/learn_go

import (
	"fmt"

	"github.com/hyun6/learn_go/accounts"
	"github.com/hyun6/learn_go/dictionary"
)

func main() {
	fmt.Println("test")
	moonAccount := accounts.NewAccount("moon")
	fmt.Println(moonAccount)

	err := moonAccount.Withdraw(10)
	if err != nil {
		fmt.Println(err)
	}

	moonAccount.Deposit(20)
	err = moonAccount.Withdraw(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(moonAccount.Balance())
	}

	myDictionary := dictionary.Dictionary{}
	myDictionary.Add("hello", "world")
	found, err := myDictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(found)
	}
	myDictionary.Delete("hello")
	found, err = myDictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(found)
	}
}
