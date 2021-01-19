package concurrent

import (
	"go-in-action/utils"
	"log"
	"strconv"
	"sync"
	"time"
)

/*
	以下描述了 select 语句的语法：
	每个 case 都必须是一个通信
	所有 channel 表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行，它就执行，其他被忽略。
	如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
	否则：
		如果有 default 子句，则执行该语句。
		如果[没有default] 子句，select 将[阻塞]，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
*/

func produce2(queue chan string, p string) {
	ret := p + "-goods-" + strconv.FormatInt(time.Now().UnixNano(), 15)
	queue <- ret
	log.Println("produce set -->:" + ret)

}

func DoRunWithProsumerSelect() {
	log.Println("start...")

	wg := sync.WaitGroup{}
	wg.Add(30)

	queue1 := make(chan string, 5)
	go func() {
		for i := 0; i < 15; i++ {
			produce2(queue1, "p1")
			utils.SleepRand(3000)
		}
	}()

	queue2 := make(chan string, 5)
	go func() {
		for i := 0; i < 15; i++ {
			produce2(queue2, "p2")
			utils.SleepRand(3000)
		}
	}()

	go func() {
		i := 0
		for {
			i++
			log.Println("start new select check", i)
			select {
			case p1v := <-queue1:
				log.Println("consume get <--:" + p1v)
				wg.Done()
			case p2v := <-queue2:
				log.Println("consume get <--:" + p2v)
				wg.Done()
			}
		}
	}()

	wg.Wait()
	log.Println("over.")
}

func DoRunWithProsumerSelectWithDefault() {
	log.Println("start...")

	wg := sync.WaitGroup{}
	wg.Add(30)

	queue1 := make(chan string, 5)
	go func() {
		for i := 0; i < 15; i++ {
			produce2(queue1, "p1")
			utils.SleepRand(3000)
		}
	}()

	queue2 := make(chan string, 5)
	go func() {
		for i := 0; i < 15; i++ {
			produce2(queue2, "p2")
			utils.SleepRand(3000)
		}
	}()

	go func() {
		i := 0
		for {
			i++
			log.Println("start new select check", i)
			select {
			case p1v := <-queue1:
				log.Println("consume get <--:" + p1v)
				wg.Done()
			case p2v := <-queue2:
				log.Println("consume get <--:" + p2v)
				wg.Done()
			default:
				log.Println("no consume,sleep 50 milliseconds and loop")
				utils.Sleep(50)
			}
		}
	}()

	wg.Wait()
	log.Println("over.")
}
