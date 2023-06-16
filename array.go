package main
import "fmt"
// import "strings"

func main(){
	// var arr []
	arr:= []int{1, 2, 3, 4, 5}
	arr2:= []string{"Sravani","is","awesome"}
	arr2= append(arr2, "and happy")
	fmt.Println(arr)
	fmt.Println(arr2)
}