package models

// BookStatus represents the status of a book.
type BookStatus string

// Enum-like constants for BookStatus.
const (
    StatusAvailable BookStatus = "Available"
    StatusBorrowed  BookStatus = "Borrowed"
    StatusReserved  BookStatus = "Reserved"
)

type Book struct {
    Id     int
    Title  string
    Author string
    Status BookStatus
}