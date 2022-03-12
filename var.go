package main

import "fmt"

func main() {
  var a1 int
  var a2 int = 123
  var a3 = 123 // 型省略
  a4 := 123 // var省略

  var (
    a5 int = 123
    a6 int = 456
  ) // まとめて定義

  var name string
  var age int
  name, age = "Yamada", 26 // 複数代入

  fmt.Println("var 1:", a1, a2, a3, a4, a5, a6, name, age)

  a1 = 456 // 代入
  name = "Suzuki"; age = 62; // 複数文

  a2 = uint(a2) // キャスト

  fmt.Println("var 2:", a1, a2, a3, a4, a5, a6, name, age)

  const foo = 100
  const (
    hoge = 200
    piyo = 300
  )

  /*コメント*/
  /*
    コメント
  */

  fmt.Println("const:", foo, hoge, piyo)

}
