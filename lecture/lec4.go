package lecture

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// goquery download : go get github.com/PuerkitoBio/goquery

var (
	baseURL  string = "https://kr.indeed.com/jobs?q=python"
	pageUnit int    = 10
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("<!> Request Error with Status Code :", res.StatusCode)
	}
}
func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // because res.Body is opened I/O
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length() - 1
	})

	return pages
}
func getPage(page int) {
	start := page * pageUnit
	//url := fmt.Sprint(baseURL, "&start=", start)
	url := baseURL + "&start=" + strconv.Itoa(start)
	fmt.Println(url)
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // because res.Body is opened I/O
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		id, isExist := s.Attr("data-jk")
		if isExist {
			fmt.Println(" ---> ", id)
		}
	})
}
func Lec400() {
	totalPages := getPages()
	fmt.Println("# of Pages : ", totalPages)
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}
