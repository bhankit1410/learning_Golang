package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	// "io"
	// "bytes"
	// "log"
	// "net/http"
)

func main() {
	var name string
	fmt.Printf("Enter the  name=>\n")
	// fmt.Scan(&s)
	reader := bufio.NewReader(os.Stdin)
	name, errName := reader.ReadString('\n')

	var address string
	fmt.Printf("Enter the  address=>\n")
	// fmt.Scan(&s)
	readerAddr := bufio.NewReader(os.Stdin)
	address, errAddr := readerAddr.ReadString('\n')
	map1 := make(map[string]string)
	if (errAddr == nil) && (errName == nil) {
		map1["name"] = name
		map1["address"] = address
	}
	data, e := json.Marshal(map1)

	if e == nil {
		fmt.Printf(string(data))
	}
}
