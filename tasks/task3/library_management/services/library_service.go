package services

import (
	"errors"

	"library_management/models"
)

// LibraryManager defines the set of methods for managing the library.
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMember(member models.Member)

}

// Library implements the LibraryManager interface.
type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

// NewLibrary creates and returns a new Library instance.
func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

// AddBook adds a new book to the library.
func (l *Library) AddBook(book models.Book) {
	l.books[book.Id] = book
	
}

// RemoveBook deletes a book from the library.
func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

// BorrowBook allows a member to borrow a book if it's available.
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != models.StatusAvailable {
		return errors.New("book is not available to borrow")
	}
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	// mark book as borrowed
	book.Status = models.StatusBorrowed
	l.books[bookID] = book
	// add the book to the member's borrowed list
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}

// ReturnBook allows a member to return a borrowed book.
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	foundIndex := -1
	for i, b := range member.BorrowedBooks {
		if b.Id == bookID {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return errors.New("member did not borrow this book")
	}
	// remove the book from member's BorrowedBooks slice
	member.BorrowedBooks = append(member.BorrowedBooks[:foundIndex], member.BorrowedBooks[foundIndex+1:]...)
	l.members[memberID] = member
	// mark the book as available
	book.Status = models.StatusAvailable
	l.books[bookID] = book
	return nil
}

// ListAvailableBooks returns a list of all available books.
func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range l.books {
		if book.Status == models.StatusAvailable {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// ListBorrowedBooks returns a list of books borrowed by a specific member.
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
func (l *Library) AddMember(member models.Member) {
    l.members[member.Id] = member
}