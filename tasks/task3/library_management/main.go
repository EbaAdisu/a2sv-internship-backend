// // filepath: /d:/Egos/a2sv/internship/a2sv-internship-backend/tasks/task3/library_management/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/controllers"
	"library_management/services"
)

func main() {
    library := services.NewLibrary()
    controller := controllers.NewLibraryController(library)
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println()
        fmt.Println("Library Management System")
        fmt.Println("1. Add Book")
        fmt.Println("2. Remove Book")
        fmt.Println("3. Borrow Book")
        fmt.Println("4. Return Book")
        fmt.Println("5. List Available Books")
        fmt.Println("6. List Borrowed Books")
        fmt.Println("7. Add Member")
        fmt.Println("8. Exit")
        fmt.Print("Choose an option: ")

        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }
        input = strings.TrimSpace(input)
        choice, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid choice. Please enter a number.")
            continue
        }

        switch choice {
        case 1:
            controller.AddBook()
			continue
        case 2:
            controller.RemoveBook()
			continue
        case 3:
            controller.BorrowBook()
			continue
        case 4:
            controller.ReturnBook()
			continue
        case 5:
            controller.ListAvailableBooks()
			continue
        case 6:
            controller.ListBorrowedBooks()
			continue
        case 7:
            controller.AddMember()
			continue
        case 8:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}