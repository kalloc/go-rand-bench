package main

import (
    "math/rand"
    "fmt"
)


func generateRandomString(length int) string {
    const alphanum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    const size = len(alphanum)
    b := make([]byte, length)
    for i := range b {
        b[i] = alphanum[rand.Intn(size)]
    }
    return string(b)
}

func generateRandomString2(length int) string {
    b := make([]byte, length)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}

func main() {
    fmt.Println(generateRandomString(10))
    fmt.Println(generateRandomString2(5))
}
