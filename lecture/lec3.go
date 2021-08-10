package lecture

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errReqFailed = errors.New("request failed")

func hitURL1(url string) error {
	fmt.Println("Cheking: ", url, " ...")
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(" <!> ", err)
		return err
	}
	if resp.StatusCode >= 400 {
		//fmt.Println(" <!> ", resp.StatusCode, " Error")
		return errReqFailed
	}
	return nil
}
func Lec300() {
	var results = make(map[string]string) // undefined map will be nil, so make is needed for an empty map

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL1(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, " --> ", result)
	}
}

func sexyCount(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "is sexy", i+1)
		time.Sleep(time.Second)
	}
}
func Lec302() {
	go sexyCount("nico") // go-Routines will be available until the main routine is done
	sexyCount("momo")
}

func isSexy1(person string, c chan bool) {
	for i := 0; i < 3; i++ {
		fmt.Println("Checking if", person, "is sexy or not:", i+1)
		time.Sleep(time.Second * 5 / 3)
	}
	fmt.Println(person, "is sexy")
	c <- true // send msg to c
}
func Lec303() {
	c := make(chan bool) // channel
	people := [4]string{"nico", "momo", "hono", "apple"}
	for _, person := range people {
		go isSexy1(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c) // receive msg from c
	}
	//fmt.Println(<-c) // Dead Lock
}

func isSexy2(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}
func Lec304() {
	c := make(chan string)
	people := []string{"nico", "momo", "hono", "apple"}
	for _, person := range people {
		go isSexy2(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

type result struct {
	status string
	url    string
}

func hitURL2(url string, c chan<- result) { // send only
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, status: "FAILED"}
		return
	}
	c <- result{url: url, status: "OK"}
}
func Lec305() {
	//results := make(map[string]string)
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL2(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
