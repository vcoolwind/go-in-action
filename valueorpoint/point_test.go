package valueorpoint

import (
	"bytes"
	"fmt"
	"testing"
)

type Student struct {
	name string
	age  int
}

func (s Student) String() string {
	return fmt.Sprintf("{name=%s,age=%d}", s.name, s.age)
}

func (s *Student) increase() {
	s.age++
}

func FakeChange(stu Student) {
	stu.name = stu.name + "_fake_changed"
	fmt.Println("show in FakeChange", stu)
}

func Change(stu *Student) {
	stu.name = stu.name + "_changed"
	fmt.Println("show in Change", stu)
}

func TestValuePoint(t *testing.T) {
	stu := Student{
		name: "jack",
		age:  20,
	}
	fmt.Println("-->origin", stu)
	FakeChange(stu)
	fmt.Println("-->fake changed", stu)
	Change(&stu)
	fmt.Println("-->true changed", stu)
	stu.increase()
	fmt.Println("-->stu increase", stu)

}

func DoFunc(f func(ps []string), paras ...string) {
	fmt.Println("start in DoFunc...")
	f(paras)
	fmt.Println("end in DoFunc.")
}

func MyPrint(paras []string) {
	len := 0
	var buf bytes.Buffer
	for _, value := range paras {
		buf.WriteString(",")
		buf.WriteString(value)
		len++
	}

	fmt.Println(len, buf.String())
}

func TestDoFunc(t *testing.T) {
	DoFunc(MyPrint, "a1", "b2", "c3")
}
