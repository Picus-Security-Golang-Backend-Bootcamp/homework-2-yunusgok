package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/yunusgok/go-patika/library"
)

var ErrInvalidInput = errors.New("input is invalid")

func init() {
	library.InitBooks()
}

func main() {
	args := os.Args

	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("%s operations are : \n search \n list \n buy \n delete \n", projectName)
		return
	}
	// Invalid arguments
	if len(args) > 1 && (args[1] != "list" && args[1] != "search") {
		fmt.Printf("Kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n")
		return
	}

	if len(args) == 2 && args[1] == "list" { // Listing books
		library.ListBooks()
		return
	} else if len(args) == 2 && args[1] == "search" {
		fmt.Printf("Enter the words you want to search...")
		return
	} else if args[1] == "search" { //
		word := strings.Join(args[2:], " ")
		library.ListGivenBooks(library.FindBooks(word))
		return
	} else if args[1] == "buy" {

		if len(args) > 4 {
			fmt.Println("Enter an <ID> and <count> to buy given number of books")
			return
		}
		check1, id := library.IsInt(args[2])
		check2, count := library.IsInt(args[3])
		if !check1 || !check2 {
			fmt.Println("ID and count must be integer")
			return
		}
		library.Buy(id, count)
	} else if args[1] == "delete" {

		if len(args) > 3 {
			fmt.Println("Enter an <ID> to delete a book")
			return
		}
		check1, id := library.IsInt(args[2])
		if !check1 {
			fmt.Println("ID must be integer")
			return
		}
		library.DeleteBook(id)
	}

}
