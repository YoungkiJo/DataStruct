package dlinkedlist

type Node struct {
	data int
	next *Node
}

type List struct {
	head      *Node
	before    *Node
	cur       *Node
	numOfData int
	comp      func(d1, d2 int) int
}

func (plist *List) ListInit() {
	plist.head = &Node{next: nil}
	plist.numOfData = 0
	plist.comp = nil
}

func (plist *List) FInsert(data int) {
	newNode := &Node{data: data}

	newNode.next = plist.head.next
	plist.head.next = newNode
	plist.numOfData++
}

func (plist *List) SInsert(data int) {
	pred := plist.head
	newNode := &Node{data: data}

	for pred.next != nil && plist.comp(data, pred.next.data) != 0 {
		pred = pred.next
	}

	newNode.next = pred.next
	pred.next = newNode
	plist.numOfData++
}

func (plist *List) LInsert(data int) {
	if plist.comp == nil {
		plist.FInsert(data)
	} else {
		plist.SInsert(data)
	}
}

func (plist *List) LFirst(pdata *int) bool {
	if plist.head.next == nil {
		return false
	}

	plist.before = plist.head
	plist.cur = plist.head.next
	*pdata = plist.cur.data

	return true
}

func (plist *List) LNext(pdata *int) bool {
	if plist.cur.next == nil {
		return false
	}

	plist.before = plist.cur
	plist.cur = plist.cur.next
	*pdata = plist.cur.data

	return true
}

func (plist *List) LRemove() int {
	rdata := plist.cur.data

	plist.before.next = plist.cur.next
	plist.cur = plist.before
	plist.numOfData--

	return rdata
}

func (plist *List) LCount() int {
	return plist.numOfData
}

func SetSortRule(plist *List, comp func(d1, d2 int) int) {
	plist.comp = comp
}
