package error_panic_defer

import (
	"fmt"
	"go-in-action/errs"
	"strconv"
	"testing"
)

func Add(i int, j int) (int, error) {
	if i < 0 || j < 0 {
		return 0, errs.NewBizError("para can not be negative")
	} else {
		return i + j, nil
	}
}

func TestError(t *testing.T) {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	sum, err := Add(-1, 2)
	if err != nil {
		fmt.Println("can not be sum,the error msg is:", err)
	} else {
		fmt.Println("get sum:", sum)
	}
}

func Add2(i int, j int) (int) {
	if i < 0 || j < 0 {
		panic("para can not be negative")
	} else {
		return i + j
	}
}

func TestPanic1(t *testing.T) {
	fmt.Println(Add2(1, -2))
}

func TestPanic2(t *testing.T) {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("get fatal error:", p)
		}
	}()
	fmt.Println(Add2(1, -2))
}
