package main

import (
	"fmt"
	// "io/ioutil"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	type person struct {
		fname string
		lname string
	}

	var filepath string
	fmt.Printf("Enter the  filepath=>\n")
	fmt.Scan(&filepath)
	people := make([]person, 6, 10)
	csvfile, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		people[i].fname = strings.SplitAfter(record[0], " ")[0]
		people[i].lname = strings.SplitAfter(record[0], " ")[1]
		i++
	}
	for _, v := range people {
		fmt.Printf("fname: %s , lname: %s \n", v.fname, v.lname)
	}

}
