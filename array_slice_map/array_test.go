package array_slice_map

import (
	"fmt"
	"testing"
)

func TestArray1(t *testing.T) {
	arr1 := [3]string{"1a1", "1b2", "1c3"}
	fmt.Println("the length of arr1 is", len(arr1), "cap is", cap(arr1))
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	for pos, v := range arr1 {
		fmt.Println("[", pos, "] = ", v)
	}

	arr2 := []string{"2a1", "2b2", "2c3", "2d4"}
	fmt.Println("the length of arr2 is", len(arr2), "cap is", cap(arr2))
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[3])
	for pos, v := range arr2 {
		fmt.Println("[", pos, "] = ", v)
	}
}

func TestSlice(t *testing.T) {
	arr1 := [3]string{"1a1", "1b2", "1c3"}
	fmt.Println(arr1)
	slice1 := arr1[1:]
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice1[0], slice1[1])
	// slice is point ref
	slice1[0] = slice1[0] + "_changed"
	slice1[1] = slice1[1] + "_changed"

	fmt.Println(slice1)
	fmt.Println(arr1)
}

func TestMap(t *testing.T) {
	mymap := make(map[string]string)
	mymap["k1"] = "v1"
	mymap["k2"] = "v2"
	mymap["k3"] = "v3"
	mymap["k4"] = "v4"
	fmt.Println(len(mymap), mymap)
	for k, v := range mymap {
		fmt.Println(k, " = ", v)
	}
	getFromMap(mymap, "k1")
	getFromMap(mymap, "k5")
	fmt.Println("delete key","k1","k5")
	delete(mymap, "k1")
	delete(mymap, "k5")
	getFromMap(mymap, "k1")
	getFromMap(mymap, "k5")

}

func getFromMap(m map[string]string, k string) {
	v, ok := m[k]
	if ok {
		fmt.Println("get value from key ", k, v)
	} else {
		fmt.Println("can not get value from key  ", k, v)
	}
}
