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
        print("%d ", )