package main
import "fmt"

func main(){


var a int = 10
var b int = 20


fmt.Println("Values before swapping",a,b)

swap(&a, &b)
fmt.Println("Values after swapping",a,b)
fmt.Println("yayy")



}
func swap(x *int,y *int){
var temp int
temp = *x
*x = *y
*y = temp
}

