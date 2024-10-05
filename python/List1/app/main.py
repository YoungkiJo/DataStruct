from ArrayList.ArrayList import ArrayList

from typing import Optional

if __name__ == "__main__":
    data: Optional[int]

    list = ArrayList()

    list.LInsert(11)
    list.LInsert(11)
    list.LInsert(22)
    list.LInsert(22)
    list.LInsert(33)

    print(f"현재 데이터의 수: {list.LCount()}")

    if list.LFirst():
        print(f"{list.arr[list.curPosition]}", end=" ")

        while(list.LNext()):
            print(f"{list.arr[list.curPosition]}", end=" ")
    print("\n")

    if list.LFirst():
        if list.arr[list.curPosition] == 22:
            list.LRemove()

        while list.LNext():
            if list.arr[list.curPosition] == 22:
                list.LRemove()
    print("\n")

    print(f"현재 데이터의 수: {list.LCount()}")

    if list.LFirst():
        print(f"{list.arr[list.curPosition]}", end=" ")

        while list.LNext():
            print(f"{list.arr[list.curPosition]}", end=" ")
    print("\n")