package main

import "fmt"
import "time"

func funcA() {
  for i := 0; i < 10; i++ {
    fmt.Print("A")
    time.Sleep(10 * time.Millisecond)
  }
}

func funcB(chA chan <- string) {
  time.Sleep(3 * time.Second)
  chA <- "Finished"
}

func main() {
  fmt.Println("[goroutine 1]")
  go funcA()
  for i := 0; i < 10; i++ {
    fmt.Print("M")
    time.Sleep(20 * time.Millisecond)
  }

  fmt.Println("\n[goroutine 2]")
  chA := make(chan string)
  defer close(chA)
  go funcB(chA)
  msg := <- chA
  fmt.Println(msg)
}
