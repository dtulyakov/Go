package main

import (
  "os/exec"
  "os"
)

func main() {
  app := "ls"
  arg0 := " --help"
//  arg0 := "--status"
//  arg1 := "--delete-after"
//  arg2 := "--recursive"
//  arg3 := "--force"
//  arg4 := "--cvs-exclude"
//  arg5 := "--exclude=vendor --exclude=modules --exclude=myoctober --exclude=plugins --exclude=themes --exclude=storage ./ ../bo/"

//  cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
  cmd := exec.Command(app, arg0)
  stdout, err := cmd.Output()

  if err != nil {
    println(err.Error())
    return
  }

  print(string(stdout))
}
