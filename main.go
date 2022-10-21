package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	var link string
	var vu int
	s := make(map[int]int)
	c := make(chan int, 99999)

	fmt.Println("Enter the Link you want to Load test: ")
	fmt.Scanf("%v", &link)

	fmt.Println("Enter the Number of Concurrent Users to Simulate: ")
	fmt.Scanf("%v", &vu)

	u, err := url.ParseRequestURI(link)
	if err != nil {
		fmt.Printf("Error: %v\nEnter correct Link\n", err)
		os.Exit(1)
	}
	fmt.Println(u)
	start := time.Now()
	for i := 0; i < vu; i++ {
		go sendRequest(link, c, i)
	}
	for i := 0; i < vu; i++ {
		s[<-c]++
	}

	fmt.Printf("Time taken: %v\n", time.Since(start))
	fmt.Println(s)
}

func sendRequest(l string, c chan int, i int) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Printf("Request #%v\n", i)
	c <- resp.StatusCode
}
