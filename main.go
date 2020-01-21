package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func readUrlsFromFile(filename string) []string {
	var urls []string
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer file.Close()
	fileReader := bufio.NewReader(file)

	for line, _, err := fileReader.ReadLine(); err != io.EOF; line, _, err = fileReader.ReadLine() {
		if err != nil {
			log.Fatalf("%v", err)
		}

		urls = append(urls, string(line))
	}

	return urls
}

func checkUrl(url string) (int, error) {
	response, err := http.Get(url)

	if err != nil {
		return -1, err
	}

	return response.StatusCode, nil
}

func printResult(status int, url string) {
	fmt.Printf("%d | %s\n", status, url)
}

func printErroredUrls(erroredUrls []string) {
	for _, url := range erroredUrls {
		fmt.Println(url)
	}
}

func runCommand(fileName string, url string) {
	var urls []string
	var erroredUrls []string

	if fileName != "" {
		fileUrls := readUrlsFromFile(fileName)
		urls = append(urls, fileUrls...)
	}

	if url != "" {
		urls = append(urls, url)
	}

	for _, url := range urls {
		status, err := checkUrl(url)

		if err != nil {
			erroredUrls = append(erroredUrls, url)
		} else {
			printResult(status, url)
		}
	}

	if len(erroredUrls) != 0 {
		fmt.Println("Error occurred during execution of the urls below")
		printErroredUrls(erroredUrls)
	}
}

func main() {
	fileFlag := flag.String("f", "", "-f <filename>")
	urlFlag := flag.String("u", "", "-u <url>")

	flag.Parse()
	runCommand(*fileFlag, *urlFlag)
}
