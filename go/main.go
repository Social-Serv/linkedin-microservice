package main

import (
       "net/http"
       "fmt"
)

func main() {
     resp, err := http.Get("https://google.com")
     fmt.Printf("%v", resp)
     fmt.Printf("%v", err)
}
