from pydantic import BaseModel, Field
from typing import Optional, Callable

class Node(BaseModel):
    data: int
    next: Optional['Node'] = None


class DLinkedList(BaseModel):
    head: Optional[Node] = Field(default_factory=lambda: Node(data=None, next=None))
    before: Optional[Node] = None
    cur: Optional[Node] = None
    num_of_data: int = 0
    comp: Optional[Callable[[int, int], int]] = None

    def FInsert(self, data: int):
        new_node = Node(data=data)
        new_node.next = self.head.next
        self.head.next = new_node
        self.num_of_data += 1

    def SInsert(self, data: int):
        pred = self.head
        new_node = Node(data=data)

        while pred.next is not None and self.comp(data, pred.next.data) != 0:
            pred = pred.next

        new_node.next = pred.next
        pred.next = new_node
        self.num_of_data+=1


    def LInsert(self, data):
        if self.comp is None:
            self.FInsert(data)
        else:
            self.SInsert(data)

    def LFirst(self) -> bool:
        if self.head.next is None:
            return False
        
        self.before = self.head
        self.cur = self.head.next
        return True

    def LNext(self) -> bool:
        if self.cur.next is None:
            return False
        
        self.before = self.cur
        self.cur = self.cur.next
        return True

    def LRemove(self):
        self.before.next = self.cur.next
        self.cur = self.before
        self.num_of_data -= 1


    def LCount(self):
        return self.num_of_data

    def SetSortRule(self, comp):
        self.comp = comp
    