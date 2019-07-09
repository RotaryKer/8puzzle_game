package main

import (
	"fmt"
	"reflect"
)

type X struct {
	before []int
	now    []int
}

func main() {
	init := []int{8, 0, 6, 5, 4, 7, 2, 3, 1} //initialize puzzle board
	fini := []int{1, 2, 3, 4, 5, 6, 7, 8, 0} //define final form

	if solve(init, fini) == 1 {
		fmt.Println("I got answer")
	} else {
		fmt.Println("I couldn't find answer")
	}
}

func makeData(data []int) int {
	return data[0]*100000000 + data[1]*10000000 + data[2]*1000000 + data[3]*100000 + data[4]*10000 + data[5]*1000 + data[6]*100 + data[7]*10 + data[8]*1
}

func searchElem(data []int, elem int) int { //return index of needed index or couldnt find element in slice return -1
	for i, v := range data {
		if v == elem {
			return i
		}
	}
	return -1
}

func changeElem(data []int, a, b int, stack [][]int, usedData []X) ([][]int, []X) {
	cpdata := make([]int, len(data))
	copy(cpdata, data)

	cpdata[a], cpdata[b] = cpdata[b], cpdata[a]
	temp := X{data, cpdata}

	return append(stack, cpdata), append(usedData, temp)
}

func check(data []int, a, b int, h map[int]int) bool {
	cpdata := make([]int, len(data))
	copy(cpdata, data)
	cpdata[a], cpdata[b] = cpdata[b], cpdata[a]
	_, ok := h[makeData(cpdata)]
	if ok {
		return false
	} else {
		h[makeData(cpdata)] = 1
		return true
	}
}

func solve(init, fini []int) int {
	history := map[int]int{makeData(init): 1}
	usedData := []X{X{nil, init}}
	stack := [][]int{init}
	finidata := makeData(fini)
	for len(stack) > 0 {
		board := stack[0]
		pos := searchElem(board, 0)
		switch pos {
		case 0:
			if check(board, 0, 1, history) {
				stack, usedData = changeElem(board, 0, 1, stack, usedData)
			}
			if check(board, 0, 3, history) {
				stack, usedData = changeElem(board, 0, 3, stack, usedData)
			}
		case 1:
			if check(board, 1, 0, history) {
				stack, usedData = changeElem(board, 1, 0, stack, usedData)
			}
			if check(board, 1, 4, history) {
				stack, usedData = changeElem(board, 1, 4, stack, usedData)
			}
			if check(board, 1, 2, history) {
				stack, usedData = changeElem(board, 1, 2, stack, usedData)
			}
		case 2:
			if check(board, 2, 1, history) {
				stack, usedData = changeElem(board, 2, 1, stack, usedData)
			}
			if check(board, 2, 5, history) {
				stack, usedData = changeElem(board, 2, 5, stack, usedData)
			}
		case 3:
			if check(board, 3, 0, history) {
				stack, usedData = changeElem(board, 3, 0, stack, usedData)
			}
			if check(board, 3, 4, history) {
				stack, usedData = changeElem(board, 3, 4, stack, usedData)
			}
			if check(board, 3, 6, history) {
				stack, usedData = changeElem(board, 3, 6, stack, usedData)
			}
		case 4:
			if check(board, 4, 1, history) {
				stack, usedData = changeElem(board, 4, 1, stack, usedData)
			}
			if check(board, 4, 3, history) {
				stack, usedData = changeElem(board, 4, 3, stack, usedData)
			}
			if check(board, 4, 5, history) {
				stack, usedData = changeElem(board, 4, 5, stack, usedData)
			}
			if check(board, 4, 7, history) {
				stack, usedData = changeElem(board, 4, 7, stack, usedData)
			}
		case 5:
			if check(board, 5, 2, history) {
				stack, usedData = changeElem(board, 5, 2, stack, usedData)
			}
			if check(board, 5, 4, history) {
				stack, usedData = changeElem(board, 5, 4, stack, usedData)
			}
			if check(board, 5, 8, history) {
				stack, usedData = changeElem(board, 5, 8, stack, usedData)
			}
		case 6:
			if check(board, 6, 3, history) {
				stack, usedData = changeElem(board, 6, 3, stack, usedData)
			}
			if check(board, 6, 7, history) {
				stack, usedData = changeElem(board, 6, 7, stack, usedData)
			}
		case 7:
			if check(board, 7, 6, history) {
				stack, usedData = changeElem(board, 7, 6, stack, usedData)
			}
			if check(board, 7, 4, history) {
				stack, usedData = changeElem(board, 7, 4, stack, usedData)
			}
			if check(board, 7, 8, history) {
				stack, usedData = changeElem(board, 7, 8, stack, usedData)
			}
		case 8:
			if check(board, 8, 5, history) {
				stack, usedData = changeElem(board, 8, 5, stack, usedData)
			}
			if check(board, 8, 7, history) {
				stack, usedData = changeElem(board, 8, 7, stack, usedData)
			}
		}

		stack = stack[1:]

		_, ok := history[finidata]
		if ok {
			checkPath(usedData, init, fini)
			return 1
		}
	}
	return 0
}

func checkPath(data []X, init []int, fini []int) {
	before := fini
	for !reflect.DeepEqual(before, init) {
		for _, i := range data {
			if reflect.DeepEqual(i.now, before) {
				before = i.before
				draw8(i.now)
				fmt.Println("////////////////////")
			}
		}
	}
}

func draw8(d []int) {
	fmt.Println("_____________")
	fmt.Println("|", d[0], "|", d[1], "|", d[2], "|")
	fmt.Println("_____________")
	fmt.Println("|", d[3], "|", d[4], "|", d[5], "|")
	fmt.Println("_____________")
	fmt.Println("|", d[6], "|", d[7], "|", d[8], "|")
	fmt.Println("_____________")
}
