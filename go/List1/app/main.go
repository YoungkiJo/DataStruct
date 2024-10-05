package main

import (
	"fmt"
	arraylist "list1/app/ArrayList"
)

func main() {
	var list arraylist.List
	var data int

	arraylist.ListInit(&list)

	arraylist.LInsert(&list, 11)
	arraylist.LInsert(&list, 11)
	arraylist.LInsert(&list, 22)
	arraylist.LInsert(&list, 22)
	arraylist.LInsert(&list, 33)

	fmt.Printf("현재 데이터의 수: %d \n", arraylist.LCount(&list))

	if arraylist.LFirst(&list, &data) {
		fmt.Printf("%d ", data)

		for arraylist.LNext(&list, &data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Println()

	if arraylist.LFirst(&list, &data) {
		if data == 22 {
			arraylist.LRemove(&list)
		}

		for arraylist.LNext(&list, &data) {
			if data == 22 {
				arraylist.LRemove(&list)
			}

		}
	}
	fmt.Println()

	fmt.Printf("현재 데이터의 수: %d \n", arraylist.LCount(&list))

	if arraylist.LFirst(&list, &data) {
		fmt.Printf("%d ", data)

		for arraylist.LNext(&list, &data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Println()
}
