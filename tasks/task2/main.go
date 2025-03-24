package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
    inputText := "Hello, world! Hello, GoLang. world..."
    freq := wordFrequency(inputText)
    fmt.Println("Word Frequency Count for", freq)
    for word, count := range freq {
        fmt.Printf("%s: %d\n", word, count)
    }

    // Example usage for palindrome check
    palindromeTest := "A man, a plan, a canal: Panama"
    isPal := isPalindrome(palindromeTest)
    fmt.Printf("\nIs \"%s\" a palindrome? %v\n", palindromeTest, isPal)
}

func wordFrequency(s string) map[string]int {
    s = strings.ToLower(s)

    // Remove punctuation. Only allow letters, digits, and white space.
    clean := func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
            return r
        }
        return -1
    }
    s = strings.Map(clean, s)

    words := strings.Fields(s)

    freq := make(map[string]int)
    for _, word := range words {
        freq[word]++
    }
    return freq
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)

    var cleaned []rune
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            cleaned = append(cleaned, r)
        }
    }

    n := len(cleaned)
    for i := 0; i < n/2; i++ {
        if cleaned[i] != cleaned[n-1-i] {
            return false
        }
    }
    return true
}