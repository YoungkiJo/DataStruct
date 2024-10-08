package read

import "fmt"

func Arrayread() {
	var arr [10]int
	var readCount int = 0
	var readData int

	for {
		fmt.Printf("자연수 입력: ")
		fmt.Scanf("%d", &readData)
		if readData < 1 {
			break
		}
		arr[readCount] = readData
		readCount += 1
	}

	for i := 0; i < readCount; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}
