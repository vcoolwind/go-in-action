package utils

import (
	"reflect"
)

type future struct {
	retChan chan []reflect.Value
	Result  []interface{}
}

func RunAsync(f interface{}, paras ...interface{}) future {
	fu := future{}
	fu.execute(f, paras...)
	return fu
}

func (s *future) execute(f interface{}, paras ...interface{}) {
	s.retChan = make(chan []reflect.Value)
	go func() {
		fun := reflect.ValueOf(f)
		inargs := make([]reflect.Value, len(paras))
		for k, param := range paras {
			inargs[k] = reflect.ValueOf(param)
		}
		r := fun.Call(inargs)
		s.retChan <- r
	}()
}

func convert(values []reflect.Value) []interface{} {
	if values != nil {
		arrs := make([]interface{}, len(values))
		for i, v := range values {
			arrs[i] = v.Interface()
		}
		return arrs
	} else {
		return nil
	}
}

func (s *future) Get() []interface{} {
	if s.Result == nil && s.retChan != nil {
		s.Result = convert(<-s.retChan)
	}
	return s.Result
}

func (s *future) GetNow(valueIfAbsent ...interface{}) []interface{} {
	if s.Result != nil {
		return s.Result
	} else {
		if s.retChan != nil {
			select {
			case values := <-s.retChan:
				s.Result = convert(values)
				return s.Result
			default:
				return valueIfAbsent
			}
		} else {
			panic("future must be revoked by RunAsync")
		}
	}
}
