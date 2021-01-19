package concurrent

import (
	"container/list"
	"go-in-action/utils"
	"log"
	"strconv"
	"sync"
)

var repos = list.New()
var mutex = sync.Mutex{}

func produce_list(group *sync.WaitGroup, p string) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		m := p + "-msg-" + strconv.Itoa(i)
		log.Println("[produce]set data to list:", m)
		repos.PushBack(m)
		utils.SleepRand(300)
		mutex.Unlock()
	}
	group.Done()
}

func consume_list() interface{} {
	mutex.Lock()
	defer mutex.Unlock()
	front := repos.Front()
	if front != nil {
		repos.Remove(front)
		log.Println("[consume]get data from list:", front.Value)
		return front.Value
	} else {
		return nil
	}
}

func RunProsumerWithMutex() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 3; i++ {
		go produce_list(&wg, "p"+strconv.Itoa(i))
	}
	go func() {
		k := 0
		for {
			data := consume_list()
			if data == nil {
				utils.SleepRand(200)
				k++
				if k > 10 {
					wg.Done()
				}
			}
		}
	}()
	wg.Wait()
}
