package futureadvance

import (
	"fmt"
	"go-in-action/utils"
	"reflect"
	"testing"
)

type futureadvance struct {
	retChan chan []reflect.Value
	Result  []reflect.Value
}

func (s *futureadvance) Do(f interface{}, paras ...interface{}) {
	s.retChan = make(chan []reflect.Value)
	go func() {
		v := apply(f, paras)
		s.retChan <- v
	}()
}

func apply(f interface{}, args []interface{}) []reflect.Value {
	fun := reflect.ValueOf(f)
	inargs := make([]reflect.Value, len(args))
	for k, param := range args {
		inargs[k] = reflect.ValueOf(param)
	}
	r := fun.Call(inargs)
	return r
}

func (s *futureadvance) Get() []reflect.Value {
	if s.Result == nil && s.retChan != nil {
		s.Result = <-s.retChan
	}
	return s.Result
}

func (s *futureadvance) GetNow() []reflect.Value {
	if s.Result == nil && s.retChan != nil {
		select {
		case s.Result = <-s.retChan:
			return s.Result
		default:
			return nil
		}
	}
	return s.Result
}

func boilWater(litre float64) string {
	var ret string
	ret = fmt.Sprint("boil water   litres ", litre)
	utils.SleepRand(5000)
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
	utils.SleepRand(5000)
	return sum, str
}

func TestFuture1(t *testing.T) {
	boilFuture := futureadvance{}
	boilFuture.Do(boilWater, 3.02)
	addFuture := futureadvance{}
	addFuture.Do(add, 1, 2)
	fmt.Println("do async")

	boilRet1 := boilFuture.Get()
	fmt.Println("boil goal result:", boilRet1[0])
	addRet := addFuture.Get()
	fmt.Println("add goal result:", addRet[0], addRet[1])

}

func TestFuture2(t *testing.T) {
	boilFuture := futureadvance{}
	boilFuture.Do(boilWater, 3.02)
	addFuture := futureadvance{}
	addFuture.Do(add, 1, 2)
	fmt.Println("do async")
	for {
		boilRet1 := boilFuture.GetNow()
		if boilRet1 != nil {
			fmt.Println("boil goal result:", boilRet1[0])
		} else {
			fmt.Println("boil goal result nil")
		}
		addRet := addFuture.GetNow()
		if addRet != nil {
			fmt.Println("add goal result:", addRet[0], addRet[1])
		} else {
			fmt.Println("add goal result nil")
		}
		if boilRet1 != nil && addRet != nil {
			break
		} else {
			utils.Sleep(500)
		}
	}

}
