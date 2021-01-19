package struct_and_interface

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	//结构体类型和普通的字符串、整型一样，也可以使用同样的方式声明和初始化。
	// 在下面的例子中，我声明了一个 person 类型的变量 p，因为没有对变量 p 初始化，所以默认会使用结构体里字段的零值。
	var p person
	fmt.Println("default zero value", p)
	p.name = "jack"
	p.age = 30
	fmt.Println("set value", p)
}

type teacher interface {
	teach(subject string)
	working()
}

type student interface {
	study(subject string)
}

type personT struct {
	name string
	age  int
}

func (s personT) teach(subject string) {
	fmt.Println("I am a teacher,I can teach ", subject)
}

type personS struct {
	name string
	age  int
}

func (s personS) study(subject string) {
	fmt.Println("I am a student,I can study ", subject)
}

type personTS struct {
	name string
	age  int
}

func (s personTS) teach(subject string) {
	fmt.Println("I am a teacher,I can teach ", subject)
}
func (s personTS) study(subject string) {
	fmt.Println("I am a student,I can study ", subject)
}

func showObj(o interface{}){
	s,oks:=o.(student)
	t,okt:=o.(teacher)

	if oks && okt {
		fmt.Println("o is student and teacher")
		s.study("english")
		t.teach("chinese")

	}else{
		if oks {
			fmt.Println("o is student")
			s.study("english")
		}else if okt {
			fmt.Println("o is teacher")
			t.teach("chinese")
		}else {
			fmt.Println("o is not student and is not teacher")
		}
	}

}

func TestInterface(test *testing.T) {
	t := personT{
		name: "teacher-zhang",
		age:  40,
	}

	s := personS{
		name: "student-ming",
		age:  10,
	}

	ts:=personTS{
		name: "teacher-student-zm",
		age:  28,
	}

	showObj(t)
	showObj(s)
	showObj(ts)
}

