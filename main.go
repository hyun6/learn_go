package main //github.com/hyun6/learn_go

import (
	"fmt"

	"github.com/hyun6/learn_go/accounts"
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
}
