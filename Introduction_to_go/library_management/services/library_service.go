package services

import (
	"bufio"
	"fmt"
	"library_management/models"
	"os"
	"strings"
)

type ErrBAB struct {} // Book is already borrowed

func (err ErrBAB)Error() string {
	return "Error: Book is already borrowed! Please come back later."
}

type ErrBNF struct {} // Book not found

func (err ErrBNF) Error() string {
	return "Error: Book not found! Please make sure you entered a valid Book ID."
}

type ErrBNB struct {} // Book is not borrowed

func (err ErrBNB) Error() string {
	return "Error: Book is not borrowed!"
}

type Library struct {
	Books map[int]models.Book 
	Members map[int]models.Member
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func (l Library) AddBook(book models.Book) {
	_, ok := l.Books[book.ID]
	if ok {
		fmt.Println("Book already exists!")
		return
	}
	l.Books[book.ID] = book 
	fmt.Println("Book added succesfully!")
}

func (l Library) RemoveBook(bookID int) {
	_, ok := l.Books[bookID]
	if !ok {
		fmt.Println("Book not found!")
		return 
	}
	fmt.Printf("\nBook %s removed successfully!\n", l.Books[bookID].Title)
	delete(l.Books, bookID)
}

func (l Library) BorrowBook(bookID, memberID int) error {
	reader := bufio.NewReader(os.Stdin)
	_, ok := l.Books[bookID]
	if !ok {
		return ErrBNF{}
	}
	if l.Books[bookID].Status == "Borrowed" {
		return ErrBAB{}
	}

	_, ok = l.Members[memberID] 
	if !ok {
		// New member registers his/her name.
		fmt.Printf("\nYou are a new member!\nEnter your name: ")
		name, err := reader.ReadString('\n')
		checkerr(err)

		tmp := l.Members[memberID]
		tmp.Name = strings.TrimSpace(name) 
		l.Members[memberID] = tmp 
	}

	tmp := l.Books[bookID]
	tmp.Status = "Borrowed"
	l.Books[bookID] = tmp 
	tmp2 := l.Members[memberID]
	tmp2.BorrowedBooks = append(tmp2.BorrowedBooks, l.Books[bookID])
	l.Members[memberID] = tmp2

	fmt.Printf("\n%s Borrowed  %s Successfully!\n", l.Members[memberID].Name, l.Books[bookID].Title)

	return nil
}

func (l Library) ReturnBook(bookID , memberID int) error {
	if l.Books[bookID].Status == "Available" {
		return ErrBNB{}
	}
	for i, v := range l.Members[memberID].BorrowedBooks {
		if v.ID == bookID {
			tmp := l.Books[bookID]
			tmp.Status = "Available"
			l.Books[bookID] = tmp 

			tmp2 := l.Members[memberID] 
			tmp2.BorrowedBooks = tmp2.BorrowedBooks[:i]
			if i < len(tmp2.BorrowedBooks) - 1 {
				tmp2.BorrowedBooks = append(tmp2.BorrowedBooks, l.Members[memberID].BorrowedBooks[i+1:]...)
			}
			l.Members[memberID] = tmp2 
			return nil
		}
	}
	return ErrBNF{}
}

func (l Library) ListAvailableBooks() []models.Book {
	tmp := []models.Book{}
	for _, book := range l.Books {
		if book.Status == "Available" {
			tmp = append(tmp, book)
		}
	}
	return tmp 
}

func (l Library) ListBorrowedBooks(memberID int) []models.Book {
	_, ok := l.Members[memberID]
	if !ok {
		fmt.Println("Member not found! Please enter a valid member ID.")
		return []models.Book{}
	}
	return l.Members[memberID].BorrowedBooks
}

func checkerr(err error) {
	if err != nil {
		os.Exit(0)
	}
}

func Service(X int, l Library) {
	reader := bufio.NewReader(os.Stdin)
	switch X {
	case 1:
		{
			var book models.Book
			var err error

			fmt.Printf("\nEnter the book title: ")
			book.Title , err = reader.ReadString('\n') 
			book.Title = strings.TrimSpace(book.Title)
			checkerr(err)

			fmt.Printf("Enter the Author of the book: ")
			book.Author, err = reader.ReadString('\n')
			book.Author = strings.TrimSpace(book.Author)

			checkerr(err)

			for {
				var id int

				fmt.Printf("Enter the book ID: ")
				fmt.Scanln(&id)

				_, ok := l.Books[id]
				if ok {
					fmt.Println("Book ID already in use!")
					continue
				}
				book.ID = id
 				break
			}
			
			book.Status = "Available"
			l.Books[book.ID] = book
			fmt.Printf("\nBook %s added successfully!\n", book.Title)
		}
	case 2:
		{
			var id int
			fmt.Printf("Enter ID of book to remove: ")
			fmt.Scanln(&id)

			l.RemoveBook(id)
		}
	case 3:
		{
			var book_id int
			fmt.Printf("\nEnter the id of the book you would like to borrow: ")
			fmt.Scanln(&book_id)
			
			var member_id int
			fmt.Printf("Enter your member id: ")
			fmt.Scanln(&member_id)

			err := l.BorrowBook(book_id, member_id)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case 4:
		{
			var member_id int 
			fmt.Printf("Enter your member id: ")
			fmt.Scanln(&member_id)

			var book_id int
			fmt.Printf("Enter your book id: ")
			fmt.Scanln(&book_id)

			err := l.ReturnBook(book_id, member_id)
			if err != nil {
				fmt.Println(err)
				return 
			}
			fmt.Printf("\n%s returned %s successfully!\n", l.Members[member_id].Name, l.Books[book_id].Title)
		}
	case 5:
		{
			fmt.Println("\nHere is a list of all available books:")
			fmt.Println(l.ListAvailableBooks())
		}
	case 6:
		{
			var id int
			
			fmt.Printf("Enter member ID: ")
			fmt.Scanln(&id)

			fmt.Println(l.ListBorrowedBooks(id))
		}
	case 7:
		fmt.Println("\nExiting program! Have a nice day.")
	}
}
