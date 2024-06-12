###  Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

```go
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```
> 2,  1
> так как defer захватывает значение переменной вов ремя её объявления. В 1  случае переменная была  объявлена позже поэтмоу бедет 1+1 = 2
> Во втором случае у нас переменная была объявлена заранее и равна 0, соответственно defer захватит её и сделает 0+1 = 1.
