from ArrayList.ArrayList import ArrayList

from typing import Optional

if __name__ == "__main__":
    data: Optional[int]

    list = ArrayList()

    list.LInsert(1)
    list.LInsert(2)
    list.LInsert(3)
    list.LInsert(4)
    list.LInsert(5)
    list.LInsert(6)
    list.LInsert(7)
    list.LInsert(8)
    list.LInsert(9)

    print(f"현재 데이터의 수: {list.LCount()}")

    if list.LFirst():
        print(f"{list.arr[list.curPosition]}", end=" ")

        while(list.LNext()):
            print(f"{list.arr[list.curPosition]}", end=" ")
    print("\n")

    sum=0
    if list.LFirst():
        sum += list.arr[list.curPosition]

        while list.LNext():
            sum += list.arr[list.curPosition]
    print(f"현재 데이터의 총합: {sum}")


    if list.LFirst():
        if list.arr[list.curPosition]%2==0 or list.arr[list.curPosition]%3==0:
            list.LRemove()
        
        while list.LNext():
            if list.arr[list.curPosition]%2==0 or list.arr[list.curPosition]%3==0:
                list.LRemove()
    print(f"현재 데이터의 수: {list.LCount()}")

    if list.LFirst():
        print(f"{list.arr[list.curPosition]}", end=" ")

        while(list.LNext()):
            print(f"{list.arr[list.curPosition]}", end=" ")
    print("\n")