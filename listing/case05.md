### Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
    msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
    {
        // do something
    }
    return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```
> Так как мы возвращаем конкретный тип, то ошибка будет иметь два поля, сама ошибка <nil> и тип ошибки *main.customErrorerror
> так  что результат будет error
