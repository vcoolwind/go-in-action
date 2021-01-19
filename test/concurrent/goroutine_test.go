package concurrent

import (
	"go-in-action/utils"
	"log"
	"testing"
)

func TestRoutine_1(t *testing.T) {
	go func() {
		log.Println("i am running in routine thread")
	}()
	log.Println("i am running in main routine thread")

	log.Println("main thread over.")
	// TODO 这里面有个bug，主线程已经退出，但routine还没有执行完毕。

}

func TestRoutine_2(t *testing.T) {
	//所以这段代码里有两个 goroutine，
	// 一个是 main 函数启动的 main goroutine，
	// 一个是通过 go 关键字启动的 goroutine。
	go func() {
		log.Println("i am running in routine thread")
	}()
	log.Println("i am running in main routine thread")

	//通过sleep使得routine执行完成，但并不靠谱
	utils.SleepSeconds(1)
	log.Println("main thread over.")
}

func TestRoutine_3(t *testing.T) {
	signal:=make(chan string)
	go func() {
		log.Println("i am running in routine thread")
		signal<-"routine exec ok"
	}()

	log.Println("i am running in main routine thread")
	ret:=<-signal
	log.Println("get result from chan:"+ret)
	log.Println("main thread over.")
}
