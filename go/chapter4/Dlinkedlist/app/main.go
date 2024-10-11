package main

import (
	"dlinkedlist/app/dlinkedlist"
	"fmt"
)

func WhoIsPrecede(d1, d2 int) int {
	if d1 < d2 {
		return 0
	} else {
		return 1
	}
}

func main() {
	var list dlinkedlist.List
	var data int

	list.ListInit()

	dlinkedlist.SetSortRule(&list, WhoIsPrecede)

	list.LInsert(11)
	list.LInsert(11)
	list.LInsert(22)
	list.LInsert(22)
	list.LInsert(33)

	fmt.Printf("Current data count: %d\n", list.LCount())

	if list.LFirst(&data) {
		fmt.Printf("%d ", data)

		for list.LNext(&data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Println()

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

	fmt.Printf("Current data count after removal: %d\n", list.LCount())

	if list.LFirst(&data) {
		fmt.Printf("%d ", data)

		for list.LNext(&data) {
			fmt.Printf("%d ", data)
		}
	}
	fmt.Println()

}
