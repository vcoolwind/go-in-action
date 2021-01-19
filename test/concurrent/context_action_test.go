package concurrent

import (
	"go-in-action/concurrent"
	"testing"
)

func Test_RUN_ctx_1(t *testing.T){
	concurrent.RUN_ctx_cancel_1()
}

func Test_RUN_ctx_2(t *testing.T){
	concurrent.RUN_ctx_cancel_2()
}
