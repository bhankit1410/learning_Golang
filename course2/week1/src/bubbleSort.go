package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slic []int, index int) {
	// swaps
	// fmt.Printf("swapping %d, %d \n", slic[index], slic[index+1])
	// temp := slic[index]
	// slic[index] = slic[index+1]
	// slic[index+1] = temp
	slic[index], slic[index+1] = slic[index+1], slic[index]
}

func BubbleSort(sli []int) {
	// sorts
	// fmt.Printf("inside bubblesort")
	sliLen := len(sli)
	i, j := 0, 0
	for i < sliLen {
		// fmt.Printf("\ni = %d\n", i)
		j = 0
		for j < (sliLen - i - 1) {

			// fmt.Printf("j = %d \n", j)
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
			j++
		}
		i++
	}

}
func main() {
	fmt.Printf("Enter the  sequence=>\n")
	reader := bufio.NewReader(os.Stdin)
	inputLine, error := reader.ReadString('\n')
	var lineInp []string
	seq := make([]int, 0, 10)
	if error == nil {

		lineInp = strings.SplitAfter(inputLine, " ")
		fmt.Printf("no error")
	} else {
		fmt.Printf("cant read line")
	}
	for i, v := range lineInp {

		newInt, err := strconv.Atoi(strings.TrimSpace(v))
		if err == nil {
			// seq[i] = newInt
			seq = append(seq, newInt)
			fmt.Printf("seq[%d] added\n", i)
		} else {
			// fmt.Printf("cant convert %d because of : %s", newInt, err)
		}

	}
	fmt.Printf("given sequence:\n")
	PrintSequence(seq)
	BubbleSort(seq)
	fmt.Printf("\n sorted sequence:\n")
	PrintSequence(seq)
}

func PrintSequence(seq []int) {
	for _, v := range seq {
		fmt.Printf(" %d  ", v)
	}
}
