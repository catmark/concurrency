package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var BOOKS int = 10

func serial() {
	start := time.Now()
	total := 0
	for i := 0; i < BOOKS; i++ {
		total += countWords(i)
	}
	fmt.Println("Serial", total, time.Since(start))
}

func parallel() {
	start := time.Now()
	resp := make(chan int)
	total := 0

	for i := 0; i < BOOKS; i++ {
		go countWordsParallel(i, resp)
	}
	for i := 0; i < BOOKS; i++ {
		total += <- resp
	}
	fmt.Println("Parallel", total, time.Since(start))
}

func countWordsParallel(bookNo int, resp chan int) {
	resp <- countWords(bookNo)
}

func countWords(bookNo int) int {
	resp, err := http.Get(fmt.Sprintf("https://www.gutenberg.org/ebooks/%d.txt.utf-8", bookNo))
	if err != nil {return 0}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {return 0}
	words := strings.Split(string(body), " ")
	fmt.Printf("Counted book %d: %d words\n", bookNo, len(words))
	return len(words)
}

func main() {
	serial()
	parallel()
}