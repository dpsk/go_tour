// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    var words []string = strings.Split(s, " ")
    for _, value := range words {
      count, ok := m[value]
        if ok {
            m[value] = count + 1
        } else {
            m[value] = 1
        }
    }

    return m
}

func main() {
    wc.Test(WordCount)
}