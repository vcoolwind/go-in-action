package utils

import (
	"go-in-action/utils"
	"log"
	"testing"
)

func TestSleep(t *testing.T) {
	log.Println("start...")
	log.Println("will sellp 1500 Milliseconds")
	utils.Sleep(1500)

	log.Println("will sleep 3 seconds")
	for i := 0; i < 3; i++ {
		utils.SleepSeconds(1)
		log.Printf("has been sleep seconds:%v", i)
	}
	log.Println("over .")

}
