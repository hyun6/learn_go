package main //github.com/hyun6/learn_go

import (
	"fmt"
	"time"

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

	now := time.Now()
	results := map[string]string{}
	ch := make(chan urlcheck.Result)
	for _, url := range urls {
		go urlcheck.Check(url, ch)
	}

	//for result := range ch { // wait forever
	for i:= 0; i < len(urls); i++ {
		result := <- ch
		results[result.Url] = result.Status
		fmt.Println(result.Url, ": ", result.Status, "ms")
	}

	fmt.Println("done")
	fmt.Println("time : ", time.Since(now).Milliseconds())
}
