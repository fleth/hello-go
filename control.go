package main

import "fmt"

func compare(x int, y int) string {
  if x > y {
    return "Big"
  } else if x < y {
    return "Small"
  } else {
    return "Equal"
  }
}

func toStatusMessage(mode string) string {
  switch mode {
  case "run":
    fallthrough // 下へ
  case "running": // run or running
    return "実行中"
  case "stop":
    return "停止中"
  default:
    return "不明"
  }
}

func funcA() {
   for a:=0; a<5; a++ {
      fmt.Println(a)
      if a >= 3 {
         goto Loop
      }
   }
   Loop:
      fmt.Println("test")
}

func main() {
  fmt.Println("[if]")
  fmt.Println(compare(1, 2))
  fmt.Println(compare(2, 1))
  fmt.Println(compare(2, 2))

  fmt.Println("[switch]")
  fmt.Println(toStatusMessage("run"))
  fmt.Println(toStatusMessage("running"))
  fmt.Println(toStatusMessage("stop"))
  fmt.Println(toStatusMessage("hoge"))

  fmt.Println("[for]")
  var x = 0
  var y = 10
  for x < y { // while
    x++
  }

  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  n := 0
  for { // 無限ループ
    n++
    if n > 10 {
      break
    } else if n % 2 == 1 {
      continue
    } else {
      fmt.Println(n)
    }
  }

  colors := [...]string{"Red", "Green", "Blue"}

  for i, color := range colors {
    fmt.Printf("%d: %s\n", i, color)
  }

  fmt.Println("goto")
  funcA()
}
