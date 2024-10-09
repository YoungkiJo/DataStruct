package main

import (
	"fmt"
	"p3_1/app/arraylist"
)

func main() {

	var list arraylist.List
	var data int
	list.ListInit()

	list.LInsert(1)
	list.LInsert(2)
	list.LInsert(3)
	list.LInsert(4)
	list.LInsert(5)
	list.LInsert(6)
	list.LInsert(7)
	list.LInsert(8)
	list.LInsert(9)

	fmt.Printf("현재 데이터의 수: %d \n", list.LCount())

	if list.LFirst(&data) {
		fmt.Printf("%d ", data)

		for list.LNext(&data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Printf("\n\n")

	var sum int = 0
	if list.LFirst(&data) {
		sum += data

		for list.LNext(&data) {
			sum += data
		}
	}

	fmt.Printf("현재 데이터 총합: %d \n", sum)

	if list.LFirst(&data) {
		if data%2 == 0 || data%3 == 0 {
			list.LRemove()
		}

		for list.LNext(&data) {
			if data%2 == 0 || data%3 == 0 {
				list.LRemove()
			}
		}
	}
	fmt.Printf("현재 데이터 수: %d \n", list.LCount())

	if list.LFirst(&data) {
		fmt.Printf("%d ", data)

		for list.LNext(&data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Printf("\n\n")
}
