package main

import (
  "strings"
  "io/ioutil"
  "net/http"
  "fmt"
  "sort"
)

func main() {

  resp, err := http.Get("https://dotnetperls-controls.googlecode.com/files/enable1.txt")

  println(resp)
  println(err)

  if resp != nil {

    body, _ := ioutil.ReadAll(resp.Body)
    if body != nil {

      content := string(body)
      words := strings.Split(content, "\r\n")

      for _, word := range words {
        if strings.HasPrefix(word, "at") {
          replaced := strings.Replace(word, "at", "@", 1)
          output := strings.Join([]string{replaced, word}, " : ")
          fmt.Println(output)
        }
      }
    }
  }
}

