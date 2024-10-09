package pointlist

import "fmt"

const LIST_LEN int = 100

type List struct {
	arr         [LIST_LEN]Point
	numOfData   int
	curPosition int
}

func (plist *List) ListInit() {
	plist.numOfData = 0
	plist.curPosition = -1
}

func (plist *List) LInsert(data Point) {
	if plist.numOfData >= LIST_LEN {
		fmt.Printf("저장이 불가능 합니다.")
		return
	}

	plist.arr[plist.numOfData] = data
	plist.numOfData += 1
}

func (plist *List) LFirst(pdata *Point) bool {
	if plist.numOfData == 0 {
		return false
	}

	plist.curPosition = 0
	*pdata = plist.arr[0]
	return true
}

func (plist *List) LNext(pdata *Point) bool {
	if plist.curPosition >= plist.numOfData-1 {
		return false
	}

	plist.curPosition += 1
	*pdata = plist.arr[plist.curPosition]
	return true
}

func (plist *List) LRemove() {
	var rpos int = plist.curPosition
	var num int = plist.numOfData

	for i := rpos; i < num-1; i++ {
		plist.arr[i] = plist.arr[i+1]
	}

	plist.numOfData -= 1
	plist.curPosition -= 1
}

func (plist *List) LCount() int {
	return plist.numOfData
}
