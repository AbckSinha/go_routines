package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	//fmt.Println(<-c) hangs the program entirely

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	//_____________ implementing infinite loop ___________________________

	// for {
	// 	go checkLink(<-c, c)
	// }

	//____________________ alternate for  loop _________________________

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) // sleep by 5 seconds
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second) // sleep by 5 seconds
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, "might be down")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up !")
	// c <- "Up and running"
	c <- link
}
