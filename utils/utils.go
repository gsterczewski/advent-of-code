package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func GetInput(inputPath string) *bufio.Scanner {

	wd := "C:/Users/GS/Programowanie/Advent-Of-Code/inputs"
	fullPath := path.Join(wd, inputPath)
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("invalid input file path")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

type Stack[T any] struct {
	internalStack []T
	len           int
}

func (st Stack[T]) Size() int {
	return st.len
}
func (st Stack[T]) Top() T {
	if st.IsEmpty() {
		return *new(T)
	}
	return st.internalStack[st.Size()-1]
}
func (st Stack[T]) IsEmpty() bool {
	return len(st.internalStack) == 0
}
func (st *Stack[T]) Pop() T {
	if st.IsEmpty() {
		return *new(T)
	}
	lastElement := st.Top()
	lastIndex := st.len - 1
	st.internalStack = st.internalStack[:lastIndex]
	st.len--
	return lastElement
}
func (st *Stack[T]) Push(el T) {
	st.internalStack = append(st.internalStack, el)
	st.len++
}
func CreateStack[T any]() Stack[T] {
	return Stack[T]{internalStack: make([]T, 0), len: 0}
}
