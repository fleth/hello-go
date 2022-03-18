package main

import "fmt"

func main() {
  // 定義
  a1 := map[string]int{
    "x": 100,
    "y": 200,
  }

  fmt.Println(a1["x"])

  // 要素追加
  a1["z"] = 300

  // 要素削除
  delete(a1, "z")

  fmt.Println(len(a1))

  _, ok := a1["z"]
  if ok {
    fmt.Println("Exist")
  } else {
    fmt.Println("Not exist")
  }
  // ループ
  for key, value := range a1 {
    fmt.Printf("%s=%d\n", key, value)
  }
}
