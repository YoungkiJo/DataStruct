package read

import "fmt"

type Node struct {
	data int
	next *Node
}

func Linkedread() {
	var head *Node = nil
	var tail *Node = nil
	var cur *Node = nil

	var newNode *Node = nil
	var readData int

	for {
		fmt.Printf("자연수 입력: ")
		fmt.Scanf("%d", &readData)
		if readData < 1 {
			break
		}

		newNode = &Node{data: readData, next: nil}

		if tail == nil {
			tail = newNode
		} else {
			head.next = newNode
		}
		head = newNode
	}
	fmt.Printf("\n")

	fmt.Printf("입력 받은 데이터의 전체출력! \n")
	if tail == nil {
		fmt.Printf("저장된 자연수가 존재하지 않습니다. \n")
	} else {
		cur = tail
		fmt.Printf("%d ", cur.data)

		for cur.next != nil {
			cur = cur.next
			fmt.Printf("%d ", cur.data)
		}
	}
	fmt.Printf("\n\n")

	if tail == nil {
		return
	} else {
		cur := tail
		for cur != nil {
			fmt.Printf("%d을(를) 삭제합니다. \n", cur.data)
			cur = cur.next
		}
	}
}
