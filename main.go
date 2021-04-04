package main //github.com/hyun6/learn_go

import (
	"fmt"

	"github.com/hyun6/learn_go/accounts"
	"github.com/hyun6/learn_go/dictionary"
	"github.com/hyun6/learn_go/urlcheck"
)

func main() {
	testAccount()
	testDictionary()
	testUrlcheck()
}

func testAccount() {
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

func testDictionary() {
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

func testUrlcheck() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	results := map[string]string{}
	for _, url := range urls {
		err := urlcheck.Check(url)
		if err != nil {
			results[url] = "FAILED"
		} else {
			results[url] = "OK"
		}
	}

	for url, result := range results {
		fmt.Println(url, ": ", result)
	}
}
