### Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.

```go
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}
```
> 3, 2, 3 
> Так как у нас слайс изначально имеет cap равный 3-ем  и при добавление  4, у нас  начнется эвакуация данных и будет создан новый слайс, где первые 3 элемента будут скопированы, а 4 встанет в конце.  Дальнейшие изменения не имеют отношения к первончачальному слайсу