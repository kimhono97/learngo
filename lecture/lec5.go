package lecture

// echo (Web Server) download : go get github.com/labstack/echo/v4

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)

var baseURL2 string = "https://kr.indeed.com/jobs?"

func getPages2(term string) int {
	pages := 0

	prams := url.Values{}
	prams.Add("q", term)
	req_url := baseURL2 + prams.Encode()
	res, err := http.Get(req_url)
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
func getPage3(page int, term string, c chan<- []jobInfo) {
	var jobs []jobInfo

	start := page * pageUnit
	prams := url.Values{}
	prams.Add("start", strconv.Itoa(start))
	prams.Add("q", term)
	req_url := baseURL2 + prams.Encode()
	res, err := http.Get(req_url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // because res.Body is opened I/O
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	cc := make(chan jobInfo)
	searchCards := doc.Find("a")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, cc)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-cc
		if job.id != "none" {
			jobs = append(jobs, job)
		}
	}

	fmt.Println("\t+" + strconv.Itoa(len(jobs)) + " jobs from Page #" + strconv.Itoa(page) + " (" + req_url + ")")
	c <- jobs
}
func writeJobs2(term string, jobs []jobInfo) {
	f, err := os.Create("jobs/" + term + ".csv")
	checkErr(err)

	fmt.Println("** Writing Results ...")
	w := csv.NewWriter(f)
	defer w.Flush()

	headers := []string{"ID", "Title", "Company", "Location", "Salary", "Summary", "URL"}
	err = w.Write(headers)
	checkErr(err)
	for _, job := range jobs {
		link_url := "https://kr.indeed.com/viewjob?jk=" + job.id
		err = w.Write([]string{job.id, job.title, job.company, job.location, job.salary, job.summary, link_url})
		checkErr(err)
	}
}
func scrap(term string) {
	var jobs []jobInfo
	c := make(chan []jobInfo)

	totalPages := getPages2(term)
	fmt.Println("** Term : " + term)
	fmt.Println("** # of Pages : ", totalPages)
	for i := 0; i < totalPages; i++ {
		go getPage3(i, term, c)
	}
	for i := 0; i < totalPages; i++ {
		extracted := <-c
		jobs = append(jobs, extracted...) // Slice + Slice
	}
	writeJobs2(term, jobs)
	fmt.Println("Done! (" + strconv.Itoa(len(jobs)) + " jobs)")
}
func Lec500() {
	scrap("java")
}

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, world!")	// Response by Text
	return c.File("res/main.html") // Response by File (.html)
}
func handleScrape(c echo.Context) error {
	term := strings.ToLower(c.FormValue("term"))
	scrap(term)
	return c.Attachment("jobs/"+term+".csv", term+".csv")
}
func Lec501() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(":1323"))
}
