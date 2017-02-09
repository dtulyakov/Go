package main

import (
    "fmt"
    "log"
    "os/exec"
    "os"
)

func main() {
    out, err := exec.Command("date", "+%Y-%M-%d").Output()
    if err != nil {
        log.Fatal(err)
    }

    hostname, err := os.Hostname()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Это лять hostname %s\n", hostname)
    fmt.Printf("Это лять out %s\n", out)
}

