package main

import "fmt"

type Person struct {
  name string
  age int
}

type Book struct {
  title string
}

func fn(b1 int, b2 *int) {
  b1 = 456
  *b2 = 456
}

func main() {
  var a1 int
  var p1 *int

  p1 = &a1
  *p1 = 123

  fmt.Println(a1)

  var a2 int = 123
  var a3 int = 123
  fn(a2, &a3)
  fmt.Println(a2, a3)

  c1 := Person{"Tanaka", 26}
  p2 := &c1
  fmt.Println(c1.name)
  fmt.Println((*p2).name) // ポインタの中身にアクセス
  fmt.Println(p2.name)    // goではこうもかける

  bookList := []*Book{}

  for i := 0; i < 10; i++ {
    book := new(Book)
    book.title = fmt.Sprintf("Title#%d", i)
    bookList = append(bookList, book)
  }
  for _, book := range bookList {
    fmt.Println(book.title)
  }

}
