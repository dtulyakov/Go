package main

import (
  "os/exec"
  "os"
//  "fmt"
//  "strings"
)

func main() {
  //var DDD string
  cmd := exec.Command("hostname")
  cmd.Stdout = os.Stdout
  cmd.Run()
 //os.Setenv("DST", "DDD")
 //fmt.Println("DST:", os.Getenv("DST"))

  //fmt.Println("мазафака")

  //fmt.Println()
  //for _, e := range os.Environ() {
  //  pair := strings.Split(e, "=")
  //  fmt.Println(pair[0])
  //}

}
