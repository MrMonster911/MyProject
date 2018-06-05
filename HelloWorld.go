package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "time"
    "net/http"
)

type Event struct {
	Hest string
}

func main() {
	fmt.Println("Hello Friggin' World")
	ch := make(chan Event, 1000)
	//  ch <-
	go populateChan(ch)
    	
    // time.Sleep(1 * time.Second)
    go process(ch)

        go process(ch)
    go process(ch)
    go process(ch)

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	   fmt.Fprint(w, "Welcome to my website!")
        ch <- Event{Hest: "request"}
    })
    http.ListenAndServe(":8080", nil)
}

func populateChan(extdata chan Event) {
	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        time.Sleep(1 * time.Second)
        e := Event{Hest: scanner.Text()}
		extdata <- e
	}
    
}

func process(c chan Event) {
	fmt.Println("Processing...")
    i := 0
    for e := range(c){
        i = i + 1
        fmt.Printf("Godt event (%+v): %+v\n", i, &e)
        time.Sleep(1 * time.Second)
	}
}
