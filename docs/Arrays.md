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
