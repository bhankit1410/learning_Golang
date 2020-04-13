package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main()  {
	var s string
	fmt.Printf("Enter the  String=>\n")
	// fmt.Scan(&s)
	reader := bufio.NewReader(os.Stdin)
	s, error := reader.ReadString('\n')
	// s = strings.TrimRight(inp,"\r\n")
	if error == nil {
		var sLower string = strings.TrimSpace(strings.ToLower(s))
		if (strings.HasPrefix(sLower, "i") &&  strings.HasSuffix(sLower,"n") && strings.ContainsAny(sLower,"a")){
			fmt.Printf("Found!")
		} else {
			fmt.Printf("not found")
		}
	}
}