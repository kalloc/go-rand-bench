

## Code

```go
package main

import (
    "math/rand"
    "fmt"
)


func generateRandomString(length int) string {
    const alphanum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    const size = len(alphanum)
    b := make([]byte, length)
    for i := range b {
        b[i] = alphanum[rand.Intn(size)]
    }
    return string(b)
}

func generateRandomString2(length int) string {
    b := make([]byte, length)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}

func main() {
    fmt.Println(generateRandomString(10))
    fmt.Println(generateRandomString2(5))
}
```


## Output

```shell
SquVJQ5yYD
41100e791b
```

## Explanation

This is a small example showing how `fmt` can be expensive.
Please note that `math/rand` is not a good source for generating passwords and keys, use `crypto/rand` instead.

## Benchmark

```shell
go test -bench . -benchmem
```

```shell

goos: linux
goarch: amd64
pkg: example.com/m
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkGenerateRandomString-12     	 8636812	       129.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkGenerateRandomString2-12    	 8805774	       146.9 ns/op	      40 B/op	       3 allocs/op
PASS
ok  	example.com/m	2.698s
```
