package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// -----------------
// https://medium.com/wesionary-team/practical-example-of-concurrency-on-golang-fc4609ea8ed1
// -----------------

var statuses = make(map[int]int)
var mutex sync.Mutex

var links = []string{
	"https://mcevik.com",
	"https://www.linkedin.com/in/semustafacevik",
	"https://github.com/semustafacevik",
	"https://medium.com/@semustafacevik",
	"https://kaggle.com/semustafacevik",
}

func getLink(link string) {
	if res, err := http.Get(link); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[%d] - %s\n", res.StatusCode, link)
	}
}

func getLinkWithWG(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	getLink(link)
}

func getLinkAndLogStatus(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	if res, err := http.Get(link); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[%d] - %s\n", res.StatusCode, link)

		mutex.Lock()
		statuses[res.StatusCode]++
		mutex.Unlock()
	}
}

func getLinkWithChannel(link string, ch chan string) {
	if res, err := http.Get(link); err != nil {
		ch <- err.Error()
	} else {
		ch <- fmt.Sprintf("[%d] - %s", res.StatusCode, link)
	}
}

func main() {
	zeroConcurrency()
	// withoutWaitGroup()
	// withWaitGroup()
	// addingMutex()
	// withChannel()
}

func zeroConcurrency() {
	for _, link := range links {
		getLink(link)
	}
}

func withoutWaitGroup() {
	for _, link := range links {
		go getLink(link)
	}

	time.Sleep(3 * time.Second)
}

func withWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(len(links))

	for _, link := range links {
		go getLinkWithWG(link, &wg)
	}

	wg.Wait()
}

func addingMutex() {
	wg := sync.WaitGroup{}

	for _, link := range links {
		go getLinkAndLogStatus(link, &wg)
		wg.Add(1)
	}

	for _, link := range links {
		go getLinkAndLogStatus(link, &wg)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(statuses)
}

func withChannel() {
	ch := make(chan string, len(links))

	for _, link := range links {
		go getLinkWithChannel(link, ch)
	}

	for msg := range ch {
		fmt.Println(msg)
	}

	// TODO : ask stackoverflow
}
