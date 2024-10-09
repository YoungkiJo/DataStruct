from pydantic import BaseModel
from typing import Optional

class Node(BaseModel):
    data: int
    next: Optional['Node'] = None

def Linkedread():
    head: Node = None
    tail: Node = None
    cur: Node = None

    newNode: Node = None
    readData: int

    while True:
        readData = int(input("자연수 입력: "))
        if readData < 1:
            break

        newNode = Node(
            data=readData,
            next=None
        )

        if head == None:
            head = newNode
        else:
            tail.next = newNode
        
        tail = newNode
    print("\n")

    print("입력 받은 데이터의 전체출력! \n")
    if head == None:
        print("저장된 자연수가 존재하지 않습니다. \n")
    else:
        cur = head
        print(f"{cur.data}", end=" ")

        while cur.next != None:
            cur = cur.next
            print(f"{cur.data}", end=" ")
    
    print("\n\n")

    if head == None:
        return
    else:
        cur = head
        while cur != None:
            print(f"{cur.data}을(를) 삭제합니다. \n")
            cur = cur.next


