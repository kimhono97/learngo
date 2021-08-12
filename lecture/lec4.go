package lecture

// goquery download : go get github.com/PuerkitoBio/goquery

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	baseURL1 string = "https://kr.indeed.com/jobs?q=python"
	pageUnit int    = 10
)

type jobInfo struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

func cleanString(s string) string {
	//return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
	return strings.ReplaceAll(s, "\n", "")
}
func (j *jobInfo) clean() {
	j.id = cleanString(j.id)
	j.title = cleanString(j.title)
	j.company = cleanString(j.company)
	j.location = cleanString(j.location)
	j.salary = cleanString(j.salary)
	j.summary = cleanString(j.summary)
}

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
func getPages1() int {
	pages := 0

	res, err := http.Get(baseURL1)
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
func getPage1(page int) (jobs []jobInfo) {
	start := page * pageUnit
	//url := fmt.Sprint(baseURL, "&start=", start)
	url := baseURL1 + "&start=" + strconv.Itoa(start)
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
			job := jobInfo{
				id:       id,
				title:    s.Find(".jobTitle>span").Text(),
				company:  s.Find(".companyName>a").Text(),
				location: s.Find(".companyLocation").Text(),
				salary:   s.Find(".salary-snippet").Text(),
				summary:  s.Find(".job-snippet").Text(),
			}
			job.clean()
			jobs = append(jobs, job)
		}
	})
	return
}
func writeJobs1(jobs []jobInfo) {
	f, err := os.Create("jobs.csv")
	checkErr(err)

	fmt.Println("Writing Results ...")
	w := csv.NewWriter(f)
	defer w.Flush()

	headers := []string{"ID", "Title", "Company", "Location", "Salary", "Summary", "URL"}
	err = w.Write(headers)
	checkErr(err)
	for _, job := range jobs {
		url := "https://kr.indeed.com/viewjob?jk=" + job.id
		err = w.Write([]string{job.id, job.title, job.company, job.location, job.salary, job.summary, url})
		checkErr(err)
	}
}
func Lec400() {
	var jobs []jobInfo

	totalPages := getPages1()
	fmt.Println("# of Pages : ", totalPages)
	for i := 0; i < totalPages; i++ {
		extracted := getPage1(i)
		jobs = append(jobs, extracted...) // Slice + Slice
	}
	writeJobs1(jobs)
	fmt.Println("Done! (" + strconv.Itoa(len(jobs)) + " jobs)")
}

func extractJob(card *goquery.Selection, cc chan<- jobInfo) {
	job := jobInfo{id: "none"}
	id, isExist := card.Attr("data-jk")
	if isExist {
		job = jobInfo{
			id:       id,
			title:    card.Find(".jobTitle>span").Text(),
			company:  card.Find(".companyName>a").Text(),
			location: card.Find(".companyLocation").Text(),
			salary:   card.Find(".salary-snippet").Text(),
			summary:  card.Find(".job-snippet").Text(),
		}
		job.clean()
	}
	cc <- job
}
func getPage2(page int, c chan<- []jobInfo) {
	var jobs []jobInfo

	start := page * pageUnit
	//url := fmt.Sprint(baseURL, "&start=", start)
	url := baseURL1 + "&start=" + strconv.Itoa(start)
	res, err := http.Get(url)
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

	fmt.Println(" +" + strconv.Itoa(len(jobs)) + " jobs from Page #" + strconv.Itoa(page) + " (" + url + ")")
	c <- jobs
}
func Lec405() {
	var jobs []jobInfo
	c := make(chan []jobInfo)

	totalPages := getPages1()
	fmt.Println("# of Pages : ", totalPages)
	for i := 0; i < totalPages; i++ {
		go getPage2(i, c)
	}
	for i := 0; i < totalPages; i++ {
		extracted := <-c
		jobs = append(jobs, extracted...) // Slice + Slice
	}
	writeJobs1(jobs)
	fmt.Println("Done! (" + strconv.Itoa(len(jobs)) + " jobs)")
}
