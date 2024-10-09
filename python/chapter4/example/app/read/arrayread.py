

def Arrayread():
    arr: list = list()
    readCount: int = 0
    readData: int

    while True:
        readData = int(input("자연수 입력: "))
        if readData < 1:
            break

        arr.append(readData)
        readCount += 1

    for i in range(readCount):
        print(f"{arr[i]}", end=" ")
    print("\n")