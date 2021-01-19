package utils

import (
	"math/rand"
	"time"
)

func Sleep(qty int64) {
	time.Sleep(time.Millisecond * time.Duration(qty))
}

func SleepRand(r int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(r)))
}


func SleepSeconds(qty int64) {
	time.Sleep(time.Second * time.Duration(qty))
}




