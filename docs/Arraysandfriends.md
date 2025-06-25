# Arrays in Go

An **array** in Go is a 
- fixed-size
-  ordered collection
-  same type
-  Indexable
-  Contiguous Memory of elements. The size of an array is part of its type, and once defined, it cannot be changed.

## Declaration

```go
var a [5]int // declares an array of 5 integers
```

## Initialization

```go
b := [3]string{"Go", "Rust", "Python"}
```

## Key Points

- Arrays have a fixed length.
- The length is part of the array's type.
- Arrays are value types; assigning one array to another copies all elements.

For more flexible collections, consider using **slices**.

# Slices

A **slice** in Go is a
- dynamically-sized
- ordered collection
- same type
- indexable
- reference type that provides a flexible view into arrays.

## Declaration

```go
var s []int // declares a slice of integers
```

## Initialization

```go
t := []string{"Go", "Rust", "Python"}
```

## Key Points

- Slices can grow and shrink dynamically.
- Slices are references to underlying arrays.
- Modifying a slice may affect the underlying array and other slices sharing it.
```go
arr := [3]int{1, 2, 3}
s1 := arr[0:2] // s1 is a slice referencing elements 1 and 2
s1[0] = 10     // modifies arr[0] as well

fmt.Println(arr) // Output: [10 2 3]
fmt.Println(s1)  // Output: [10 2]
```
You can access slice elements using indices, just like arrays:
```go
fmt.Println(s1[1]) // Output: 2
```
- Slices have both length and capacity.

Slices are the idiomatic way to work with sequences in Go.

can also use make for creating slices:
```go
s := make([]int, 5)
s := make([]float32, 6)
//make also has capacity
s := make([]int, 5 , 10)
fmt.Println(len(s), cap(s)) // Output: 5 10
```
# Maps

A **map** in Go is an
- unordered collection
- of key-value pairs
- where all keys are of the same type
- and all values are of the same type.

Maps are reference types and provide fast lookups, inserts, and deletes.

## Declaration

```go
var m map[string]int // declares a map from string to int
```

## Initialization

```go
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
}
```

You can also use `make` to create a map:

```go
m := make(map[string]int)
```

## Key Points

- Maps are dynamic; they grow as you add more key-value pairs.
- Accessing a non-existent key returns the zero value for the value type.
- Use `delete(map, key)` to remove a key-value pair.
- The order of iteration over map keys is not specified and can change.

Example usage:

```go
m := map[string]int{"foo": 1, "bar": 2}
m["baz"] = 3
fmt.Println(m["foo"]) // Output: 1

delete(m, "bar")
fmt.Println(m) // Output: map[foo:1 baz:3]
```

To check if a key exists:

```go
value, ok := m["foo"]
if ok {
    fmt.Println("Key exists:", value)
}
```

# Loops

Go supports only one looping construct: the `for` loop. It can be used in several ways:

## Basic for loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## While-style loop

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

## Infinite loop

```go
for {
    // loop forever
}
```

## Looping over collections

### Arrays and Slices

```go
nums := []int{1, 2, 3}
for i, v := range nums {
    fmt.Println(i, v)
}
```

### Maps

```go
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Println(k, v)
}
```

### Strings (runes)

```go
for i, r := range "hello" {
    fmt.Println(i, r)
}
```

Use `break` to exit early and `continue` to skip to the next iteration.