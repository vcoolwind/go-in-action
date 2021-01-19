package concurrent

import (
	"go-in-action/utils"
	"log"
	"strconv"
	"sync"
	"time"
)

func produce1(queue chan string, p string) {
	ret := p + "-goods-" + strconv.FormatInt(time.Now().UnixNano(), 15)
	queue <- ret
	log.Println("produce set -->:" + ret)

}

func consume1(queue chan string) string {
	return <-queue
}

func DoRunWithProsumer1() {
	log.Println("start...")

	wg := sync.WaitGroup{}
	wg.Add(3)
	queue := make(chan string, 20)
	go func() {
		for i := 0; i < 15; i++ {
			produce1(queue, "p1")
			utils.SleepRand(1000)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 15; i++ {
			produce1(queue, "p2")
			utils.SleepRand(1000)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 30; i++ {
			ret := consume1(queue)
			log.Println("consume get <--:" + ret)
			utils.SleepRand(1000)
		}
		wg.Done()
	}()

	wg.Wait()
	log.Println("over.")
}
