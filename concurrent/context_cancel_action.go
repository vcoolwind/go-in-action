package concurrent

import (
	"context"
	"go-in-action/utils"
	"log"
	"strconv"
	"sync"
	"time"
)

func RUN_ctx_cancel_1() {
	ctx, stop_watch := context.WithCancel(context.Background())
	valueCtx := context.WithValue(context.Background(),"useID","zhangsan")
	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 1; i < 4; i++ {
		go func(pos int) {
			defer wg.Done()
			watch_dog_ctx(ctx, "dog"+strconv.Itoa(pos))
		}(i)
	}
	go get_userid_from_ctx(valueCtx,&wg)

	go func() {
		utils.Sleep(5000)
		log.Println("exit on 5 seconds sleep")
		stop_watch()
	}()

	wg.Wait()
}

func RUN_ctx_cancel_2() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(pos int) {
			defer wg.Done()
			watch_dog_ctx(ctx, "dog"+strconv.Itoa(pos))
		}(i)
	}
	wg.Wait()
}

func watch_dog_ctx(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name + " get exit signal,exit loop")
			return
		default:
			log.Println(name + " watch_dog is working...")
		}
		utils.SleepRand(1000)
	}
}

func get_userid_from_ctx(ctx context.Context,wg *sync.WaitGroup){
	defer wg.Done()
	value := ctx.Value("useID")
	log.Println("get userID from ctx:",value)
}