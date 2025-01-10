```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    switch v := i.(type) {
    case string:
        fmt.Println(v)
    default:
        fmt.Println("Type assertion failed. i is not a string.")
    }
}
```