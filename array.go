package main

import "fmt"

func main() {
  a1 := [3]string{}
  a1[0] = "Red"
  a1[1] = "Green"
  a1[2] = "Blue"
  fmt.Println(a1[0], a1[1], a1[2])
  fmt.Println(a1)

  a2 := [3]string{"R", "G", "B"}
  a3 := [...]string{"R", "G", "B", "A"}
  fmt.Println(a2, a3)

  // slice
  a4 := []int{} // slice 個数不定
  for i := 0; i < 10; i++ {
    a4 = append(a4, i)
    fmt.Println(len(a4), cap(a4)) // len: 長さ cap: 容量
  }

  bufa := make([]byte, 0, 1024) //メモリ確保
  bufa = append(bufa, 'R')
  bufa = append(bufa, 'e')
  bufa = append(bufa, 'd')
  fmt.Println(bufa, string(bufa))
}
