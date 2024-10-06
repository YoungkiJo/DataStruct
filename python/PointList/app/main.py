from pointlist.list import *
from pointlist.point import *


if __name__ == "__main__":
    list = List()
    compPos = Point()

    ppos = Point()
    ppos.SetPointPos(2, 1)
    list.LInsert(ppos)

    ppos = Point()
    ppos.SetPointPos(2, 2)
    list.LInsert(ppos)

    ppos = Point()
    ppos.SetPointPos(3, 1)
    list.LInsert(ppos)

    ppos = Point()
    ppos.SetPointPos(3, 2)
    list.LInsert(ppos)

    print(f"현재 데이터의 수: {list.LCount()}")

    if list.LFirst():
        list.arr[list.curPosition].ShowPointPos()

        while list.LNext():
            list.arr[list.curPosition].ShowPointPos()
    print("\n")

    compPos.xpos = 2
    compPos.ypos = 0

    if list.LFirst():
        if PointComp(list.arr[list.curPosition], compPos) == 1:
            list.LRemove()
        
        while list.LNext():
            if PointComp(list.arr[list.curPosition], compPos) == 1:
                list.LRemove()

    print(f"현재 데이터의 수: {list.LCount()}")

    if list.LFirst():
        list.arr[list.curPosition].ShowPointPos()

        while list.LNext():
            list.arr[list.curPosition].ShowPointPos()
    print("\n")