package main

import (
	"arraylist/app/arraylist"
	"fmt"
)

func main() {

	var list arraylist.List
	var data int
	list.ListInit()

	list.LInsert(11)
	list.LInsert(11)
	list.LInsert(22)
	list.LInsert(22)
	list.LInsert(33)

	fmt.Printf("현재 데이터의 수: %d \n", list.LCount())

	if list.LFirst(&data) {
		fmt.Printf("%d ", data)

		for list.LNext(&data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Printf("\n\n")

	if list.LFirst(&data) {
		if data == 22 {
			list.LRemove()
		}

		for list.LNext(&data) {
			if data == 22 {
				list.LRemove()
			}
		}
	}

	fmt.Printf("현재 데이터의 수: %d \n", list.LCount())

	if list.LFirst(&data) {
		fmt.Printf("%d ", data)

		for list.LNext(&data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Printf("\n\n")
}
