package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)    
    dict := make(map[string]int)

    for _,v := range words {
        dict[v] += 1
    }
    return dict
}

func main() {
    wc.Test(WordCount)
}

