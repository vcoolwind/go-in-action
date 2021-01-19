package concurrent

import (
	"go-in-action/concurrent"
	"testing"
)

func Test_mutex_r(t *testing.T){
	concurrent.RunProsumerWithMutex()
}
