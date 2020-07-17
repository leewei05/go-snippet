# Go-snippet

Go-snippet is a Go code base

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


