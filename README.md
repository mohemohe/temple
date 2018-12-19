temple
======

inline string template makes easy!

## example

```go
package main

import (
    "github.com/mohemohe/temple"
)

func main() {
    t := "hello, {{.target}}!"
    m := map[string]interface{}{
        "target": "world"
    }
    result, _ := temple.Execute(t, m)
    println(result)
    // hello, world!
}
```