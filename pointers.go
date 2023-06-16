// A pointer is a variable which stores the address of another variable

package main
import "fmt"

func main() {
	m1 := 2
	ptr := &m1
	fmt.Println(ptr)
	fmt.Println(*ptr)
	// fmt.Println(&m1)
}
