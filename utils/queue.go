package utils

import (
	"log"
	"sync"
)

type queue struct {
	cap    int
	qchan  chan interface{}
	locker sync.Mutex
}

// add o to queue tail, false  when queue is full
func (q *queue) offer(o interface{}) bool {
	q.locker.Lock()
	defer q.locker.Unlock()
	if len(q.qchan) == cap(q.qchan) {
		log.Println("queue is full",len(q.qchan))
		return false
	} else {
		q.qchan <- o
		return true
	}
}

// add o to queue tail, block  when queue is full
func (q *queue) offerBlock(o interface{}) bool {
	q.qchan <- o
	return true
}

//get o from queue head, return nil when queue is empty
func (q *queue) poll() interface{} {
	select {
	case o := <-q.qchan:
		return o
	default:
		return nil
	}
}

//get o from queue head, block  when queue is empty
func (q *queue) pollWithBlock() interface{} {
	return <-q.qchan
}

//clear all object
func (q *queue) clear() {
	for {
		select {
		case o := <-q.qchan:
			log.Println("remove obj from queue", o)
		default:
			log.Println("clean over")
			return
		}
	}
}

func NewQueue(cap int) queue {
	return queue{
		cap:    cap,
		qchan:  make(chan interface{}, cap),
		locker: sync.Mutex{},
	}
}
