package main

import (
	"fmt"
	"sort"
)

//pointers refrences in GO
func _pointer() {
	var p *int //pointer

	if p != nil {
		fmt.Println("the value of pointer is :", *p)
	} else {
		fmt.Println("the value of pointer is null")
	}
}

//simple arrays examples
func _arraysExamples(index int) {
	var _days [7]string

	_days[0] = "Monday"
	_days[1] = "Tuesday"
	_days[2] = "wenesday"
	_days[3] = "Thursday"
	_days[4] = "Friday"
	_days[5] = "Saturday"
	_days[6] = "Sunday"

	//trying to get array index
	if index != 0 {
		println(_days[index])
	} else {
		fmt.Println(_days)
	}
}

//Slices are resizable and can be sorted
func _slicesExamples() {
	var _premiumPhones = []string{"Ztel", "Samsung", "Iphone", "Nokia, zaki, Alphine"}

	fmt.Println(_premiumPhones)
	//Adding more phones to slice. => slice is similar to list
	_premiumPhones = append(_premiumPhones, "Itel Phones")
	fmt.Println(_premiumPhones)

	//removing an item from slice
	_premiumPhones = append(_premiumPhones[:len(_premiumPhones)-1])
	fmt.Println(_premiumPhones)

	fmt.Println(sort.StringSlice(_premiumPhones))

}

func main() {
	_slicesExamples()
	_arraysExamples(5)
	_pointer()
}

//in my oppinion, slice is similar to  list, enumerables in c#
