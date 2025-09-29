package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func main() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	for actualLog := range logChan {
		fmt.Println(actualLog)
	}
}
