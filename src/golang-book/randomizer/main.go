package main

import "fmt"

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang

func main() {
    fmt.Println("Привет Денис")
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
