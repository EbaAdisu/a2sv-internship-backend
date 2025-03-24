package main

import "testing"

func TestWordFrequency(t *testing.T) {
    input := "Hello, hello! World? world..."
    expected := map[string]int{
        "hello": 2,
        "world": 2,
    }
    result := wordFrequency(input)
    for word, count := range expected {
        if result[word] != count {
            t.Errorf("Expected %d for word %q, got %d", count, word, result[word])
        }
    }
}

func TestIsPalindrome(t *testing.T) {
    palindrome := "Madam, in Eden, I'm Adam"
    notPalindrome := "This is not a palindrome"

    if !isPalindrome(palindrome) {
        t.Errorf("Expected %q to be a palindrome", palindrome)
    }
    if isPalindrome(notPalindrome) {
        t.Errorf("Expected %q not to be a palindrome", notPalindrome)
    }
}