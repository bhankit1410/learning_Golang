package main
import "fmt"

func main(){
	var n float32
	fmt.Printf("Enter a floating point num:  \n")
	fmt.Scan(&n)
	new := int32(n)
	fmt.Printf("Truncated => %d ",new)
	
}