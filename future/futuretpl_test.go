package future

import (
	"fmt"
	"go-in-action/utils"
	"reflect"
	"testing"
)

type futuretest struct {
	f       interface{}
	paras   []interface{}
	retChan chan []reflect.Value
}

func (s *futuretest) Do() {
	s.retChan = make(chan []reflect.Value)
	go func() {
		v := apply(s.f, s.paras)
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

func (s futuretest) Get() []reflect.Value {
	if s.retChan != nil {
		return <-s.retChan
	} else {
		return nil
	}
}

func boilWater(litre float64) string {
	var ret string
	ret = fmt.Sprint("boil water   litres ", litre)
	utils.Sleep(3000)
	return ret
}

func TestFuture(t *testing.T) {
	psc := make([]interface{}, 1)
	psc[0] = 3.02

	fu := futuretest{
		f:     boilWater,
		paras: psc,
	}
	fu.Do()
	fmt.Println("do async")

	originRet:=fu.Get()

	fmt.Println("goal result:",originRet[0])
}
