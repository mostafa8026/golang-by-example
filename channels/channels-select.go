package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "Warning"
	logError   = "Error"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var intCh = make(chan int, 50)
var doneCh = make(chan struct{})

func main() {
	go logger()
	defer func() {
		doneCh <- struct{}{}
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting."}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting Down."}
	intCh <- 5
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case i := <-intCh:
			fmt.Println("intCh has been received: ", i)
		case <-doneCh:
			break
		}
	}
}
