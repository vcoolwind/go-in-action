package concurrent

import (
	"go-in-action/concurrent"
	"testing"
)

func TestProsumer1(t *testing.T){
	concurrent.DoRunWithProsumer1()
}

func TestProsumer2(t *testing.T){
	concurrent.DoRunWithProsumerSelect()
}

func TestProsumer3(t *testing.T){
	concurrent.DoRunWithProsumerSelectWithDefault()
}
