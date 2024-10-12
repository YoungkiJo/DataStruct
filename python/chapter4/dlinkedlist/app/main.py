from typing import Callable
from dlinkedlist.dlinkedlist import DLinkedList

def who_is_precede(d1: int, d2: int) -> int:
    if d1 < d2:
        return 0
    else:
        return 1
    

if __name__ == "__main__":


    dlist = DLinkedList()

    dlist.SetSortRule(who_is_precede)

    dlist.LInsert(11)
    dlist.LInsert(11)
    dlist.LInsert(22)
    dlist.LInsert(22)
    dlist.LInsert(33)

    print(f"Current data count: {dlist.LCount()}")
    