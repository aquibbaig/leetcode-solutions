## Golang cheatsheet
--------------
- Declare variables
```
var (
    a = 5
    a int
)
a := 20
```
- Declare an array
```
var arr []int
arr := []int{1, 2, 3}
```

- Using a random number generator [0, n)

`* Make sure to import "math/rand" package.`
```
rng := rand.Intn(end)
```

- While-loops in Golang
```
for <condition> {
    *body*
}
```

- For-loops in range

`* arr must be an iterator of some native data type.`

```
for idx, val := range arr {
    *body*
}
```

- Struct in Golang
```
type Person struct {
    a int
}
```