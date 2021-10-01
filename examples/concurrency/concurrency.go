package main

import (
	"fmt"
	"log"
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

func getLink(link string) (*http.Response, error) {
	if res, err := http.Get(link); err != nil {
		log.Println(err)
		return nil, err
	} else {
		fmt.Printf("[%d] - %s\n", res.StatusCode, link)
		return res, nil
	}
}

func getLinkWithWG(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	getLink(link)
}

func getLinkAndLogStatus(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, _ := getLink(link)
	if res != nil {
		mutex.Lock()
		statuses[res.StatusCode]++
		mutex.Unlock()
	}
}

func getLinkWithChannel(link string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	if res, err := getLink(link); err != nil {
		ch <- fmt.Sprintf("[channel-err]: %s", err)
	} else {
		ch <- fmt.Sprintf("[channel-val]: %s - (%d)", link, res.StatusCode)
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
	wg := sync.WaitGroup{}
	ch := make(chan string, len(links))

	for _, link := range links {
		wg.Add(1)
		go getLinkWithChannel(link, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
