package library

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Books []*Book
var BookNames []string

var ErrBookNotFound = errors.New("Book not found")

var ErrNotEnoughStock = errors.New("not enough stock")

//initilize books with book names and authors in list
func InitBooks() {
	for _, b := range BooksList {
		Books = append(Books, NewBook(b[0], b[1]))
	}
}

// List all books line by line by their name
func ListBooks() {
	for _, b := range Books {
		fmt.Println(b.name)
	}
}

// List given books line by line by their name
func ListGivenBooks(books []Book) {
	for _, b := range books {
		fmt.Println(b.name)
	}
}

// Searches given words in books and return matched books names
func FindBooks(word string) []Book {
	result := make([]Book, 0)
	// word is turned to lowercase to search case insensitive
	searchWord := strings.ToLower(word)
	for _, book := range Books {
		// book name is turned to lowercase to search case insensitive
		if strings.Contains(strings.ToLower(book.name), searchWord) {
			result = append(result, *book)
			// author name is turned to lowercase to search case insensitive
		} else if strings.Contains(strings.ToLower(book.author), searchWord) {
			result = append(result, *book)
			// check word is integer so ISBN number can be searched
		} else if check, value := IsInt(searchWord); check {
			if book.ISBN == value {
				result = append(result, *book)
			}
		}
	}
	if len(result) == 0 {
		fmt.Printf("'%s' not found", word)
	}
	return result
}

//Find book with id
func FindBook(id int) (Book, error) {
	book := *Books[id]
	if book.isDeleted {
		return *new(Book), ErrBookNotFound
	}
	return *Books[id], nil
}

//Buy book if enoubh count exist in stock
func Buy(id int, count int) error {
	book, err := FindBook(id)
	if err != nil {
		return err
	}
	return book.Buy(count)

}

//check given string is int and return to value
// s-> stirng to be checked
// return (check bool, value int)
// check -> whether string is int
// value -> value of the string
func IsInt(s string) (bool, int) {
	if value, err := strconv.Atoi(s); err == nil {
		return true, value
	}
	return false, 0
}

// deletes book if exist
func DeleteBook(id int) {
	b, err := FindBook(id)
	if err == nil {
		err = b.Delete()
	} else {
		fmt.Println(err.Error())
	}

}
