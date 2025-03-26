// // filepath: /d:/Egos/a2sv/internship/a2sv-internship-backend/tasks/task4/library_management/main.go
package main

import (
	"bufio"
	"fmt"
	"library_management/concurrency"
	"library_management/controllers"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func main() {
    library := services.NewLibrary()
    controller := controllers.NewLibraryController(library)

    // Start the reservation worker.
    concurrency.StartReservationWorker(library)

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
        fmt.Println("8. Reserve Book")
        fmt.Println("9. Exit")
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
        case 2:
            controller.RemoveBook()
        case 3:
            controller.BorrowBook()
        case 4:
            controller.ReturnBook()
        case 5:
            controller.ListAvailableBooks()
        case 6:
            controller.ListBorrowedBooks()
        case 7:
            controller.AddMember()
        case 8:
            controller.ReserveBook()
        case 9:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}