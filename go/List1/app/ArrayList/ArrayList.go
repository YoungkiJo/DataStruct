package arraylist

// List 구조체 정의
type List struct {
	arr         []int
	numOfData   int
	curPosition int
}

// 리스트 초기화 함수
func ListInit(list *List) {
	list.arr = []int{}
	list.numOfData = 0
	list.curPosition = -1
}

// 리스트에 데이터 삽입 함수
func LInsert(list *List, data int) {
	list.arr = append(list.arr, data)
	list.numOfData++
}

// 첫 번째 데이터를 가져오는 함수
func LFirst(list *List, data *int) bool {
	if list.numOfData == 0 {
		return false
	}
	list.curPosition = 0
	*data = list.arr[0]
	return true
}

// 다음 데이터를 가져오는 함수
func LNext(list *List, data *int) bool {
	if list.curPosition >= list.numOfData-1 {
		return false
	}
	list.curPosition++
	*data = list.arr[list.curPosition]
	return true
}

// 데이터를 삭제하는 함수
func LRemove(list *List) int {
	rdata := list.arr[list.curPosition]
	list.arr = append(list.arr[:list.curPosition], list.arr[list.curPosition+1:]...)
	list.numOfData--
	list.curPosition--
	return rdata
}

// 리스트에 저장된 데이터 수를 반환하는 함수
func LCount(list *List) int {
	return list.numOfData
}
