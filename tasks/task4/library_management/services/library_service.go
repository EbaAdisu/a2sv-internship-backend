// // filepath: /d:/Egos/a2sv/internship/a2sv-internship-backend/tasks/task3/library_management/services/library_service.go
package services

import (
	"errors"
	"sync"

	"library_management/models"
)

// LibraryManager defines the set of methods for managing the library.
type LibraryManager interface {
    AddBook(book models.Book) error
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
    AddMember(member models.Member)
    ReserveBook(bookID int, memberID int) error
    UnreserveBook(bookID int) error
}

// Library implements the LibraryManager interface.
type Library struct {
    books   map[int]models.Book
    members map[int]models.Member
    mu      sync.Mutex
}

// NewLibrary creates and returns a new Library instance.
func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

// AddBook adds a new book to the library. It returns an error if the ID is already taken.
func (l *Library) AddBook(book models.Book) error {
    l.mu.Lock()
    defer l.mu.Unlock()
    if _, exists := l.books[book.Id]; exists {
        return errors.New("book ID already exists")
    }
    l.books[book.Id] = book
    return nil
}

// RemoveBook deletes a book from the library.
func (l *Library) RemoveBook(bookID int) {
    l.mu.Lock()
    defer l.mu.Unlock()
    delete(l.books, bookID)
}

// BorrowBook allows a member to borrow a book if it is available or reserved for them.
func (l *Library) BorrowBook(bookID int, memberID int) error {
    l.mu.Lock()
    defer l.mu.Unlock()

    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }
    // Allow borrowing if available or reserved (assuming the reservation was made by the same member)
    if book.Status != models.StatusAvailable && book.Status != models.StatusReserved {
        return errors.New("book is not available to borrow")
    }
    // In a complete system you'd validate that the reserved book belongs to memberID.

    member, exists := l.members[memberID]
    if !exists {
        return errors.New("member not found")
    }
    book.Status = models.StatusBorrowed
    l.books[bookID] = book
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.members[memberID] = member
    return nil
}

// ReturnBook allows a member to return a borrowed book.
func (l *Library) ReturnBook(bookID int, memberID int) error {
    l.mu.Lock()
    defer l.mu.Unlock()

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
    member.BorrowedBooks = append(member.BorrowedBooks[:foundIndex], member.BorrowedBooks[foundIndex+1:]...)
    l.members[memberID] = member

    book.Status = models.StatusAvailable
    l.books[bookID] = book
    return nil
}

// ListAvailableBooks returns a list of all available books.
func (l *Library) ListAvailableBooks() []models.Book {
    l.mu.Lock()
    defer l.mu.Unlock()
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
    l.mu.Lock()
    defer l.mu.Unlock()
    member, exists := l.members[memberID]
    if !exists {
        return []models.Book{}
    }
    return member.BorrowedBooks
}

// AddMember adds a new member to the library.
func (l *Library) AddMember(member models.Member) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.members[member.Id] = member
}

// ReserveBook reserves a book if it is available. Returns an error if the book is not available.
func (l *Library) ReserveBook(bookID int, memberID int) error {
    l.mu.Lock()
    defer l.mu.Unlock()

    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }
    if book.Status != models.StatusAvailable {
        return errors.New("book is not available for reservation")
    }
    book.Status = models.StatusReserved
    l.books[bookID] = book
    return nil
}

// UnreserveBook cancels a reservation if the book is still reserved.
func (l *Library) UnreserveBook(bookID int) error {
    l.mu.Lock()
    defer l.mu.Unlock()

    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }
    if book.Status == models.StatusReserved {
        book.Status = models.StatusAvailable
        l.books[bookID] = book
        return nil
    }
    return errors.New("book is not reserved")
}