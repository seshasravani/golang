
//?? why %x n all showing up?
package main
import "fmt"

func main(){
	var a int
	a=2

	var ip *int
	ip = &a 
	 fmt.Println("address of a is %x",&a)
	 fmt.Println("address at ip is  %x",ip)
	 fmt.Println("ip contains an address and the value at that address is * ip = %d",*ip)
}


// package main

// import "fmt"

// func main() {
//    var a int = 20   /* actual variable declaration */
//    var ip *int      /* pointer variable declaration */

//    ip = &a  /* store address of a in pointer variable*/

//    fmt.Printf("Address of a variable: %x\n", &a  )

//    /* address stored in pointer variable */
//    fmt.Printf("Address stored in ip variable: %x\n", ip )

//    /* access the value using the pointer */
//    fmt.Printf("Value of *ip variable: %d\n", *ip )
// }