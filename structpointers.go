package main
import "fmt"

type Garments struct{

	 style string
	 color string
	 size int
}

func main(){

var garment1 Garments
var garment2 Garments

garment1.style="trousers"
garment1.color="blue"
garment1.size=5

garment2.style="shirt"
garment2.color="white"
garment2.size=6



getGarments(&garment1) 
getGarments(&garment2)

}

func getGarments(garment *Garments){

	fmt.Println(garment.style)
	fmt.Println(garment.color)
	fmt.Println(garment.size)
}



//following commented code is my code without pointers which gives the same output

// getGarments(garment1) 
// getGarments(garment2)

// }

// func getGarments(garment Garments){

// 	fmt.Println(garment.style)
// 	fmt.Println(garment.color)
// 	fmt.Println(garment.size)
// }