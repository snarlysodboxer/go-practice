/*
Practice working with Golang Channels
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendMessages(strings []string, messageChannel chan string) {
	defer wg.Done()
	for _, str := range strings {
		messageChannel <- str
	}
}

func printMessages(messageChannel chan string) {
	for message := range messageChannel {
		fmt.Println(message)
	}
}

func main() {
	messageChannel := make(chan string)

	go printMessages(messageChannel)

	strings := [][]string{[]string{"asdf", "fdsa", "abcd", "efgh"}, []string{"ASDF", "FDSA", "ABCD", "EFGH"}}

	for _, strs := range strings {
		wg.Add(1)
		go sendMessages(strs, messageChannel)
	}

	wg.Wait()

	close(messageChannel)
}
