from pydantic import BaseModel, Field
from typing import Optional, Callable

class Node(BaseModel):
    data: int
    next: Optional['Node'] = None

class DLinkedlist:
    # head: Optional[Node] = Field(default_factory=lambda: Node(data=None, next=None))
    head: Optional[Node] = None
    before: Optional[Node] = None
    cur: Optional[Node] = None
    num_of_data: int = 0
    comp: Optional[Callable[[int, int], int]] = None
    data: int = 0

    def FInsert(self, data: int):
        newNode = Node(next=None, data=data)
        newNode.next = self.head.next
        self.head.next = newNode

        self.num_of_data+=1

    def SInsert(self, data: int):
        newNode = Node(next=None, data=data)
        pred = self.head

        while pred.next != None and self.comp(data, pred.next.data) != 0:
            pred = pred.next
        
        newNode.next = pred.next
        pred.next = newNode

        self.num_of_data+=1

    def LInsert(self, data):
        if self.comp == None:
            self.FInsert(data)
        else:
            self.SInsert(data)

    def LFirst(self) -> bool:
        if self.head.next == None:
            return False
        
        self.before = self.head
        self.cur = self.head.next

        self.data = self.cur.data
        return True
    
    def LNext(self) -> bool:
        if self.cur.next == None:
            return False
        
        self.before = self.cur
        self.cur = self.cur.next

        self.data = self.cur.data
        return True
    
    def LRemove(self):
        self.before.next = self.cur.next
        self.cur = self.before

        self.num_of_data-=1

    def LCount(self) -> int:
        return self.num_of_data

    def SetSortRule(self, comp):
        self.comp = comp