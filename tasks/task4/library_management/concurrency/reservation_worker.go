// // filepath: /d:/Egos/a2sv/internship/a2sv-internship-backend/tasks/task3/library_management/concurrency/reservation_worker.go
package concurrency

import (
	"fmt"
	"library_management/services"
	"time"
)

// ReservationRequest represents a reservation request.
type ReservationRequest struct {
    BookID   int
    MemberID int
}

// ReservationChan is a channel for queuing reservation requests.
var ReservationChan chan ReservationRequest

func init() {
    ReservationChan = make(chan ReservationRequest, 100)
}

// StartReservationWorker starts a Goroutine that processes reservation requests concurrently.
func StartReservationWorker(library services.LibraryManager) {
    go func() {
        for req := range ReservationChan {
            fmt.Printf("Processing reservation for Book %d by Member %d\n", req.BookID, req.MemberID)
            timer := time.NewTimer(5 * time.Second)
            <-timer.C

            // Attempt to auto-cancel reservation after 5 seconds.
            err := library.UnreserveBook(req.BookID)
            if err == nil {
                fmt.Printf("Reservation for Book %d cancelled due to timeout.\n", req.BookID)
            }
        }
    }()
}