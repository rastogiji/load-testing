package main

import (
	"fmt"
	"net/http"
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

	start := time.Now()
	for i := 0; i < vu; i++ {
		go sendRequest(link, c, i)
	}
	for i := 0; i < vu; i++ {
		s[<-c]++
	}

	fmt.Printf("Time taken: %v", time.Since(start))
	fmt.Println(s)
}

func sendRequest(l string, c chan int, i int) {
	resp, _ := http.Get(l)
	fmt.Printf("Request #%v", i)
	c <- resp.StatusCode
}

