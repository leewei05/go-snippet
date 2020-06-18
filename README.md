# Go-snippet

Go-snippet is a Go code base

```go
// Check empty slice
var eSlice []byte
if len(eSlice) == 0 {
    // Do something if the slice is empty
}
```

```go
// Reverse slice of integers
func reverse(s []int) {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
                s[i], s[j] = s[j], s[i]
        }
}
```

```go
// Swap values between variables
a, b = b, a
a, b, c = b, c, a
```

```go
// Copy a map
a := map[string]bool{"A": true, "B": true}
b := make(map[string]bool)
for key, value := range a {
	b[key] = value
}
```
