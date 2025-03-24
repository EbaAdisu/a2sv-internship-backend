```markdown
//// filepath: /d:/Egos/a2sv/internship/a2sv-internship-backend/tasks/task3/library_management/docs/documentation.md

# Library Management System Documentation

## Overview

The Library Management System is a simple console-based Go application that allows you to manage books and library members. It enables you to add and remove books, manage book borrowing and returns, list available/borrowed books, and add library members.

## Folder Structure
```

library_management/
├── controllers/
│ └── library_controller.go # Handles user input and interacts with the library service.
├── docs/
│ └── documentation.md # Project documentation.
├── models/
│ ├── book.go # Definition of the Book model.
│ └── member.go # Definition of the Member model.
├── services/
│ └── library_service.go # Implements the LibraryManager interface.
├── main.go # Entry point; renders the main menu and handles user interaction.
└── go.mod # Module file.

````

## Models

### Book

Defined in `models/book.go`:

- **Id**: Unique identifier for the book.
- **Title**: Book title.
- **Author**: Book author.
- **Status**: Book status (Available or Borrowed).

```go
type BookStatus string

const (
    StatusAvailable BookStatus = "Available"
    StatusBorrowed  BookStatus = "Borrowed"
)

type Book struct {
    Id     int
    Title  string
    Author string
    Status BookStatus
}
````

### Member

Defined in `models/member.go`:

-   **Id**: Unique member identifier.
-   **Name**: Member's name.
-   **BorrowedBooks**: A list of books the member has borrowed.

```go
type Member struct {
    Id            int
    Name          string
    BorrowedBooks []Book
}
```

## Services

The service layer (in library_service.go) implements the `LibraryManager` interface, which defines the following methods:

-   **AddBook(book Book)**: Adds a new book to the library.
-   **RemoveBook(bookID int)**: Removes a book by its ID.
-   **BorrowBook(bookID int, memberID int) error**: Checks if a book is available and allows a member to borrow it.
-   **ReturnBook(bookID int, memberID int) error**: Allows a member to return a previously borrowed book.
-   **ListAvailableBooks() []Book**: Lists all available books.
-   **ListBorrowedBooks(memberID int) []Book**: Lists all the books borrowed by a specific member.
-   **AddMember(member Member)**: Adds a new member to the library.

The service maintains a map for books and another for members to ensure fast lookups.

## Controllers

The controller layer (in library_controller.go) is responsible for:

-   Prompting users for input.
-   Converting the user input to the required type.
-   Calling the corresponding library service method.
-   Displaying a success or error message based on the result.

Each method in the controller creates and utilizes a buffered reader to take full line inputs, ensuring that spaces within strings (e.g., book titles or member names) are properly captured.

### Example: Adding a Book

The `AddBook` method prompts for the Book ID, title, and author. It then constructs a `Book` model and calls the service's `AddBook` method.

## Main

The entry point (`main.go`) displays a console menu with options to:

-   Add, remove, borrow, or return a book.
-   List available or borrowed books.
-   Add a library member.
-   Exit the application.

The main loop reads the user choice as a string (using a buffered reader), converts it to an integer and calls the corresponding controller method.

## How to Run

1. Open a terminal in the project root folder (`library_management/`).
2. Run the command:
    ```
    go run main.go
    ```
3. Follow the on-screen prompts to interact with the library management system.

---

```

```
