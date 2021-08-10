package lecture

import (
	"errors"
	"fmt"
	"net/http"
)

var errReqFailed = errors.New("request failed")

func hitURL(url string) error {
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
func Lec400() {
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
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, " --> ", result)
	}
}

func Lec402() {
}
