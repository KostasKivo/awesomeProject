/*
I am going to use those identifiers in main.go so I should write them LikeThis or else
compiler doesn't recognize ex. libBooks.bookidCounter but recognizes libBooks.BookidCounter
*/
package libBooks

//Book struct
type Book struct {
	bookid int
	title string
	author string
}

//Counter to increment book ids
var bookidCounter = 0

//Constructor, somewhat
func NewBook(t string, a string) Book {
	b := Book{bookid:bookidCounter+1,title:t,author:a}
	return b
}


//Creating a slice
// Type Book, len = 3 , cap = 3
var AvailBooks = make([]Book,1)

func PrintAvailBooks() {
	for _, bookOb := range AvailBooks{
		println("ID: " + string(bookOb.bookid))
		println("Title: " + string(bookOb.title))
		println("Author: " + string(bookOb.author))
		println()
	}
}

func AddBook(b Book) {
	AvailBooks = append(AvailBooks,b)
}

func RemoveBook(id int) {
	flag := false
	for i, bookOb := range AvailBooks {
		//If we find the Book with the id we are searching for
		if bookOb.bookid == id {
			//Put the last element in the position of the searched
			AvailBooks[i] = AvailBooks[len(AvailBooks)-1]
			//Copy the array from the start until len(availBooks)-1
			AvailBooks = AvailBooks[:len(AvailBooks)-1]
			flag = true
			break
		}
	}

	if flag == false {
		println("There was no such id in the list!")
	} else {
		println("Book with " + string(id) + " deleted successfully!")
	}
} // end of method