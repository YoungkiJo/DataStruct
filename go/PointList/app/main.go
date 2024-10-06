package main

import (
	"fmt"
	"pointlist/app/pointlist"
)

func main() {
	var list pointlist.List
	var compPos pointlist.Point
	var ppos *pointlist.Point

	list.ListInit()

	ppos = &pointlist.Point{}
	ppos.SetPointPos(2, 1)
	list.LInsert(*ppos)

	ppos = &pointlist.Point{}
	ppos.SetPointPos(2, 2)
	list.LInsert(*ppos)

	ppos = &pointlist.Point{}
	ppos.SetPointPos(3, 1)
	list.LInsert(*ppos)

	ppos = &pointlist.Point{}
	ppos.SetPointPos(3, 2)
	list.LInsert(*ppos)

	fmt.Printf("현재 데이터의 수: %d \n", list.LCount())

	if list.LFirst(ppos) {
		ppos.ShowPointPos()

		for list.LNext(ppos) {
			ppos.ShowPointPos()
		}
	}
	fmt.Printf("\n")

	compPos.SetPointPos(2, 0)

	if list.LFirst(ppos) {
		if pointlist.PointComp(ppos, &compPos) == 1 {
			fmt.Printf("\n")
			list.LRemove()
		}

		for list.LNext(ppos) {
			if pointlist.PointComp(ppos, &compPos) == 1 {
				list.LRemove()
			}
		}
	}

	fmt.Printf("현재 데이터의 수: %d \n", list.LCount())

	if list.LFirst(ppos) {
		ppos.ShowPointPos()

		for list.LNext(ppos) {
			ppos.ShowPointPos()
		}
	}
	fmt.Printf("\n")
}
