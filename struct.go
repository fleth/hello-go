package main

import "fmt"

type Person struct {
  name string
  age int
}
func (p *Person) SetPerson(name string, age int) {
  p.name = name
  p.age = age
}
func (p *Person) GetPerson() (string, int) {
  return p.name, p.age
}

// アクセス
type Person2 struct {
  Name string // public
  Age int     // public
  status int  // パッケージ外でprivate
}

// interface
type Printable interface {
  ToString() string
}

func PrintOut(p Printable) {
  fmt.Println(p.ToString())
}

// interface{} any型のような感じで使える
func IPrintOut(a interface{}) {
  q, ok := a.(Printable)
  if ok {
    fmt.Println(q.ToString())
  } else {
    fmt.Println("Not printable.")
  }
}

type Person3 struct {
  name string
}
func (p Person3) ToString() string {
  return p.name
}

type Book struct {
  title string
}
func (b Book) ToString() string {
  return b.title
}

func funcA(a interface{}) {
  switch a.(type) {
  case bool:
    fmt.Printf("[bool] %t\n", a)
  case int:
    fmt.Printf("[int] %d\n", a)
  case string:
    fmt.Printf("[string] %s\n", a)
  }
}

func main() {
  var p1 Person
  p1.SetPerson("Yamada", 26)
  name, age := p1.GetPerson()

  fmt.Printf("%s(%d)\n", name, age)

  a1 := Person{"Yamada", 26}
  a2 := Person{name: "Yamada", age: 26}
  a3 := Person2{Name: "Suzuki", Age: 32, status: 1}
  fmt.Println(a1, a2, a3)

  b1 := Person3{name: "山田太郎"}
  b2 := Book{title: "吾輩は猫である"}
  PrintOut(b1)
  PrintOut(b2)
  IPrintOut(b1)
  IPrintOut(b2)
  IPrintOut(a1)

  funcA(true)
  funcA(12345)
  funcA("test")

  type any interface{} // any型
  type dict map[string]any // pythonのdictのような型

  p2 := dict {
    "name": "Yamada",
    "age": 26,
    "address": dict {
      "zip": "123-4567",
      "tel": "012-3456-7890",
    },
  }

  fmt.Println(p2["name"], p2["address"].(dict)["tel"])

}
