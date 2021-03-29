package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	/*GetHtml()
	bdtxt := ReadBodyTxt()
	fmt.Print(bdtxt)
	exampletxt := ReadExampleHtml()
	fmt.Print(exampletxt)*/
	sub_result := GetSubResult()
	fmt.Println(sub_result)
	add_result := GetAddResult()
	fmt.Println(add_result)
}

func GetHtml() {
	resp, err := http.Get("http://localhost:3000/example.html/")
	body, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile("body.txt", body, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadBodyTxt() string {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("body.txt")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}
	// if it was successful in reading the file then
	// print out the contents as a string
	//fmt.Print(string(data))

	return string(data)
}

func ReadExampleHtml() string {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("HTMLFolder/static/example.html")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}
	// if it was successful in reading the file then
	// print out the contents as a string
	//fmt.Print(string(data))

	return string(data)
}

func GetSubResult() string {
	var sub_title string = ""
	// Request the HTML page.
	res, err := http.Get("http://localhost:3000/example.html/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".subclass").Each(func(i int, s *goquery.Selection) {
		title := s.Find("p").Text()

		//fmt.Println(title)
		sub_title = string(title)
	})
	return sub_title
}

func GetAddResult() string {
	var add_title string = ""
	// Request the HTML page.
	res, err := http.Get("http://localhost:3000/example.html/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".addclass").Each(func(i int, s *goquery.Selection) {
		title := s.Find("p").Text()

		//fmt.Println(title)
		add_title = string(title)
	})
	return add_title
}
