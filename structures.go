package main
import "fmt"

type Books struct{
	book_id  int
	title string
	author string
}

func main(){

	var book1 Books
	//var book2 Books
	book1.book_id=5
	book1.title ="Physics"
	book1.author ="HC Verma"
	fmt.Println(book1.title)
	fmt.Println(book1.author)
	getbook(book1)

	
}
func getbook(book Books){

	fmt.Println(book.title)
}
