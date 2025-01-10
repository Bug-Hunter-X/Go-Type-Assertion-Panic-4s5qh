```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    s := i.(string)
    fmt.Println(s)
}
```