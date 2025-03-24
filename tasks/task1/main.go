package main

import (
	"fmt"
)

func main() {
    var studentName string
    var subjectCount int

    fmt.Print("Enter your name: ")
    fmt.Scanln(&studentName)
    fmt.Print("Enter number of subjects: ")
    fmt.Scanln(&subjectCount)

    subjectGrades := make(map[string]float64)
    subjectOrder := make([]string, 0, subjectCount)

    for i := 1; i <= subjectCount; i++ {
        var subject string
        var grade float64

        fmt.Printf("Enter subject %d name: ", i)
        fmt.Scanln(&subject)
        subjectOrder = append(subjectOrder, subject)

        for {
            fmt.Printf("Enter grade for %s (0-100): ", subject)
            fmt.Scanln(&grade)
            if grade < 0 || grade > 100 {
                fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
            } else {
                break
            }
        }
        subjectGrades[subject] = grade
    }

    average := calculateAverage(subjectGrades)

    fmt.Println("\nStudent Name: ", studentName)
    fmt.Println()
    fmt.Println("Subjects and their corresponding grades:")
    for _, subj := range subjectOrder {
        fmt.Printf("%s: %.2f\n", subj, subjectGrades[subj])
    }
    fmt.Printf("\nAverage Grade is: %.2f\n", average)
}

func calculateAverage(grades map[string]float64) float64 {
    if len(grades) == 0 {
        return 0
    }
    var total float64
    for _, grade := range grades {
        total += grade
    }
    return total / float64(len(grades))
}
