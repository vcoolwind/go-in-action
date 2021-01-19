package concurrent

import (
	"go-in-action/utils"
	"log"
	"strconv"
	"sync"
)

func running(wg *sync.WaitGroup, cond *sync.Cond) {
	for i := 0; i < 10; i++ {
		go func(pos int) {
			defer wg.Done()

			log.Println("runner[" + strconv.Itoa(pos) + "] is ready.")
			cond.L.Lock()
			cond.Wait()
			log.Println("runner[" + strconv.Itoa(pos) + "] start running...")
			cond.L.Unlock()

			// do biz
			for k := 0; k < 3; k++ {
				utils.SleepRand(1000)
			}
			//do biz over

			log.Println("runner[" + strconv.Itoa(pos) + "] Get to the end.")

		}(i)
	}
}

func DoRunWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)

	running(&wg, cond)
	utils.SleepSeconds(2)
	go func() {
		defer wg.Done()
		cond.Broadcast()
	}()

	wg.Wait()
}
