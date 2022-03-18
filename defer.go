package main

import "fmt"

func funcA() {
  fmt.Println("funcA start")
  defer fmt.Println("funcA release")
  fmt.Println("funcA end")
}

func main() {
  fmt.Println("start")
  defer fmt.Println("release")
  funcA()
  fmt.Println("end")
}
