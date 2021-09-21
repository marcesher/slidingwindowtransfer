package main

import "fmt"

func canTransfer(data []int, threshold int) bool {
	max := 0
	for _, v := range data {
		if v > max {
			max = v
		}

		if max-v > threshold {
			return false
		}
	}
	return true
}

func canTransferInWindow(data []int, threshold int, windowSize int) bool {
	left := 0
	for i := range data {
		if i >= windowSize {
			left = i - windowSize + 1
		}
		slice := data[left : i+1]
		//fmt.Printf("v is %d; Slice is %+v\n\n", data[i], slice)
		if ok := canTransfer(slice, threshold); !ok {
			return false
		}
	}
	return true
}

func main() {
	input := []int{80, 90, 88, 91, 87, 90, 60}
	ok := canTransfer(input, 5)

	if ok {
		panic("should not transfer but did")
	}

	input = []int{90, 89, 87, 85, 83, 81, 79}
	ok = canTransferInWindow(input, 10, 5)
	fmt.Println("OK", ok)
	if !ok {
		panic("should transfer in window but didn't")
	}

	max := 100000000
	largeSlice := make([]int, max, max)
	for i := max - 1; i >= 0; i-- {
		largeSlice[i] = max - i
	}

	ok = canTransferInWindow(largeSlice, 5, 10)
	fmt.Println("OK", ok)

	if ok {
		panic("should not be OK and should reach threshold quickly")
	}

	ok = canTransferInWindow(largeSlice, 15, 10)
	fmt.Println("OK", ok)

	if !ok {
		panic("should be OK because threshold is greater than window size and slice decrements by 1 each time, so never will exceed threshold in a 10-size window")
	}
}
