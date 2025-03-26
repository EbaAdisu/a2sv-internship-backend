// // filepath: /d:/Egos/a2sv/internship/a2sv-internship-backend/tasks/task3/library_management/controllers/library_controller.go
package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/concurrency"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
    library services.LibraryManager
}

func NewLibraryController(library services.LibraryManager) *LibraryController {
    return &LibraryController{library: library}
}

// AddBook prompts the user to input book details and adds the book.
func (c *LibraryController) AddBook() {
    reader := bufio.NewReader(os.Stdin)

    // Read book ID
    fmt.Print("Enter book ID: ")
    idStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading book ID:", err)
        return
    }
    idStr = strings.TrimSpace(idStr)
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Invalid book ID.")
        return
    }

    // Read book title (full line)
    fmt.Print("Enter book title: ")
    title, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading title:", err)
        return
    }
    title = strings.TrimSpace(title)

    // Read book author (full line)
    fmt.Print("Enter book author: ")
    author, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading author:", err)
        return
    }
    author = strings.TrimSpace(author)

    book := models.Book{
        Id:     id,
        Title:  title,
        Author: author,
        Status: models.StatusAvailable,
    }
    c.library.AddBook(book)
    fmt.Println("Book added successfully.")
}

// RemoveBook prompts the user to input the book ID to be removed.
func (c *LibraryController) RemoveBook() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter book ID to remove: ")
    idStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading book ID:", err)
        return
    }
    idStr = strings.TrimSpace(idStr)
    bookID, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Invalid book ID.")
        return
    }

    c.library.RemoveBook(bookID)
    fmt.Println("Book removed successfully.")
}

// BorrowBook prompts the user for book ID and member ID then processes borrowing.
func (c *LibraryController) BorrowBook() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter book ID to borrow: ")
    bookIDStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading book ID:", err)
        return
    }
    bookIDStr = strings.TrimSpace(bookIDStr)
    bookID, err := strconv.Atoi(bookIDStr)
    if err != nil {
        fmt.Println("Invalid book ID.")
        return
    }

    fmt.Print("Enter member ID: ")
    memberIDStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading member ID:", err)
        return
    }
    memberIDStr = strings.TrimSpace(memberIDStr)
    memberID, err := strconv.Atoi(memberIDStr)
    if err != nil {
        fmt.Println("Invalid member ID.")
        return
    }

    err = c.library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Book borrowed successfully.")
}

// ReturnBook prompts the user for book ID and member ID then processes the return.
func (c *LibraryController) ReturnBook() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter book ID to return: ")
    bookIDStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading book ID:", err)
        return
    }
    bookIDStr = strings.TrimSpace(bookIDStr)
    bookID, err := strconv.Atoi(bookIDStr)
    if err != nil {
        fmt.Println("Invalid book ID.")
        return
    }

    fmt.Print("Enter member ID: ")
    memberIDStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading member ID:", err)
        return
    }
    memberIDStr = strings.TrimSpace(memberIDStr)
    memberID, err := strconv.Atoi(memberIDStr)
    if err != nil {
        fmt.Println("Invalid member ID.")
        return
    }

    err = c.library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Book returned successfully.")
}

// ListAvailableBooks displays all available books.
func (c *LibraryController) ListAvailableBooks() {
    books := c.library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
        return
    }
    fmt.Println("Available Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
    }
}

// ListBorrowedBooks prompts the user for a member ID and displays the books borrowed by that member.
func (c *LibraryController) ListBorrowedBooks() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter member ID: ")
    memberIDStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading member ID:", err)
        return
    }
    memberIDStr = strings.TrimSpace(memberIDStr)
    memberID, err := strconv.Atoi(memberIDStr)
    if err != nil {
        fmt.Println("Invalid member ID.")
        return
    }

    books := c.library.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println("No borrowed books for this member.")
        return
    }
    fmt.Println("Borrowed Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
    }
}

// AddMember prompts the user to input member details and adds the member.
func (c *LibraryController) AddMember() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter member ID: ")
    idStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading member ID:", err)
        return
    }
    idStr = strings.TrimSpace(idStr)
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Invalid member ID.")
        return
    }

    fmt.Print("Enter member name: ")
    name, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading member name:", err)
        return
    }
    name = strings.TrimSpace(name)

    member := models.Member{
        Id:            id,
        Name:          name,
        BorrowedBooks: []models.Book{},
    }
    c.library.AddMember(member)
    fmt.Println("Member added successfully.")
}
func (c *LibraryController) ReserveBook() {
    reader := bufio.NewReader(os.Stdin)

    // Read book ID
    fmt.Print("Enter book ID to reserve: ")
    idStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading book ID:", err)
        return
    }
    idStr = strings.TrimSpace(idStr)
    bookID, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Invalid book ID.")
        return
    }

    // Read member ID
    fmt.Print("Enter member ID: ")
    memberStr, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading member ID:", err)
        return
    }
    memberStr = strings.TrimSpace(memberStr)
    memberID, err := strconv.Atoi(memberStr)
    if err != nil {
        fmt.Println("Invalid member ID.")
        return
    }

    // Attempt to reserve the book.
    err = c.library.ReserveBook(bookID, memberID)
    if err != nil {
        fmt.Println("Reservation failed:", err)
        return
    }

    // Queue the reservation request for auto-cancellation.
    concurrency.ReservationChan <- concurrency.ReservationRequest{
        BookID:   bookID,
        MemberID: memberID,
    }
    fmt.Println("Book reserved successfully.")
}