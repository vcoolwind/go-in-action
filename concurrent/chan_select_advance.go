package concurrent

import (
	"go-in-action/utils"
	"log"
	"sync"
)

func RUN_chan_select_advance() {
	signal_chan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		watch_dog(signal_chan)
	}()

	go func() {
		utils.Sleep(5000)
		log.Println("exit on 5 seconds sleep")
		signal_chan <- "exit"
	}()
	wg.Wait()

}

func watch_dog(signal chan string) {
	for {
		select {
		case <-signal:
			log.Println("get exit signal,exit loop")
			return
		default:
			log.Println("watch_dog is working...")
		}
		utils.SleepRand(1000)
	}
}
