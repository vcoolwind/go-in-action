package utils

import (
	"log"
	"testing"
)

func TestQueue(t *testing.T) {
	queue:=NewQueue(3)
	queue.offer(1)
	queue.offer(2)
	queue.offer(3)
	if !queue.offer(4) {
		log.Println("add fail ,match test")
	}
	log.Println(queue.poll(),queue.poll(),queue.poll())
	log.Println(queue.poll()==nil)
	go func() {
		Sleep(5000)
		queue.offer(4)
	}()
	log.Println("queue is empty,will block 5 second")
	log.Println(queue.pollWithBlock())

}
