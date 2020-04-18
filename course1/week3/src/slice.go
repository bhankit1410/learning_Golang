package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	i := 1
	max := 20
	for i < max {
		var n string
		fmt.Scan(&n)

		if !(string(n) == "X") {
			newInt, err := strconv.Atoi(n)
			if err == nil {
				sli = append(sli, newInt)
				sort.Ints(sli)
				for _, elmt := range sli {
					fmt.Printf("%d", elmt)
				}
			}

		} else {
			break
		}
	}

}
