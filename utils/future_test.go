package utils

import (
	"fmt"
	"goinaction/utils"
	"log"
	"testing"
)

func boilWater(litre float64) string {
	var ret string
	ret = fmt.Sprint("boil water   litres ", litre)
	SleepRand(5000)
	return ret
}
func add(i int, j int) (int, string) {
	var sum int
	var str string
	if i < 0 || j < 0 {
		sum = -1
		str = "para can not be negative"
	} else {
		sum = i + j
		str = "success"
	}
	SleepRand(5000)
	return sum, str
}

func TestFuture1(t *testing.T) {
	boilFuture := RunAsync(boilWater, 3.02)
	addFuture := RunAsync(add, 1, 2)
	log.Println("do async")
	
	boilRet := boilFuture.Get()
	log.Println("boil goal result:", boilRet[0])
	addRet := addFuture.Get()
	log.Println("add goal result:", addRet[0], addRet[1])

}

func TestFuture2(t *testing.T) {
	boilFuture := RunAsync(boilWater, 3.02)
	addFuture := RunAsync(add, 1, 2)
	log.Println("do async")
	for {
		boilRet1 := boilFuture.GetNow()
		if boilRet1 != nil {
			log.Println("boil goal result:", boilRet1[0])
		} else {
			log.Println("boil goal result nil")
		}
		addRet := addFuture.GetNow()
		if addRet != nil {
			log.Println("add goal result:", addRet[0], addRet[1])
		} else {
			log.Println("add goal result nil")
		}
		if boilRet1 != nil && addRet != nil {
			break
		} else {
			utils.Sleep(500)
		}
	}
}

func TestFuture3(t *testing.T) {
	boilFuture := RunAsync(boilWater, 3.02)
	addFuture := RunAsync(add, 1, 2)
	log.Println("do async")
	for i := 0; i < 20; i++ {
		boilRet := boilFuture.GetNow("unknown")
		log.Println("boil goal result:", boilRet[0])
		addRet := addFuture.GetNow(0, "fail")
		log.Println("add goal result:", addRet[0], addRet[1])
		utils.Sleep(500)
	}
}

type person struct {
	Name string
	Age  int
}

func clonePerson(p person) person {
	c := p
	c.Name = c.Name + "_cloned"
	return c
}

func newPerson(nm string, age int) person {
	p := person{
		Name: nm,
		Age:  age,
	}
	return p
}

func TestFuture4(t *testing.T) {
	np := newPerson("jack", 12)
	cp := clonePerson(np)
	log.Println("new person", np, &np)
	log.Println("clone person", cp, &cp)
	log.Println("compare", cp == np, &cp == &np)

	newFuture := RunAsync(newPerson, "jackma", 20)
	ret := newFuture.Get()
	log.Println(ret, ret[0])
	if retp, ok := ret[0].(person); ok {
		log.Println("future get:", retp.Name, retp.Age)
	}
}
