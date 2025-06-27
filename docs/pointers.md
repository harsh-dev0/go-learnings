# Pointers in Go

Pointers are variables that store the memory address of another variable.

## Declaring Pointers

The `new` keyword in Go allocates memory for a variable of the specified type and returns a pointer to it. For example, `new(int32)` allocates memory for an `int32` and returns a pointer to that memory location.

Using `new` is useful when you want a pointer to a newly allocated zero-value variable, rather than the address of an existing variable.

It's important to check for `nil` pointers before dereferencing them. Dereferencing a `nil` pointer (trying to access the value it points to) will cause a runtime error (panic) in Go. Always ensure a pointer is properly initialized (using `new`, the address-of operator `&`, or assignment) before using it.

```go
var p *int32 = new(int32)

var a int = 10
var p *int = &a // p holds the address of a
```

### The `&` (Address-of) and `*` (Dereference) Operators in Pointers

- The `&` operator is used to get the memory address of a variable. For example, `p := &a` assigns the address of `a` to the pointer `p`. This does **not** create a new copy of `a` in memory; instead, both `a` and `*p` refer to the same memory location. This is efficient because you avoid duplicating data.
- The `*` operator is used to access (dereference) the value stored at the address a pointer refers to. For example, `*p` gives you the value of `a` if `p` is `&a`. Again, this does not copy the value; it simply reads or writes to the original memory location.

**Why is this efficient?**  
When you use `&` to get the address of a variable, you are just passing around a reference (the memory address), not duplicating the underlying data. This is especially useful for large structs or data types, as it avoids unnecessary memory usage and copying. When you use `*` to access or modify the value, you are directly working with the original data.

#### Common Pitfalls with Pointers

- **Dereferencing nil pointers:** Attempting to use `*p` when `p` is `nil` will cause a runtime panic.
- **Uninitialized pointers:** Declaring a pointer without assigning it (e.g., `var p *int`) leaves it as `nil`.
- **Forgetting to use `&`:** If you want a pointer to an existing variable, you must use `&` to get its address.
- **Confusing `*` in declaration vs. usage:** In `var p *int`, the `*` means "pointer to int". In `*p = 10`, the `*` means "the value pointed to by p".

Example:
```go
var a int = 42
var p *int = &a // & gets the address of a
fmt.Println(*p) // * gets the value at that address (42)
```

## Dereferencing Pointers

```go
fmt.Println(*p) // Outputs: 10
```

## Modifying Values via Pointers

```go
*p = 20
fmt.Println(a) // Outputs: 20
```

## Nil Pointers

Uninitialized pointers have a `nil` value.

```go
var ptr *int
fmt.Println(ptr) // Outputs: <nil>
```

## Why Use Pointers?

- To share and modify data across functions
- To avoid copying large structs

## Example: Passing Pointers to Functions

```go
func increment(n *int) {
    *n++
}

num := 5
increment(&num)
fmt.Println(num) // Outputs: 6
```

## Pointers and Slices

Slices in Go are reference types, which means they internally use pointers to reference underlying arrays. When you pass a slice to a function, the function receives a copy of the slice header (which contains a pointer to the array, length, and capacity), but both the original and the copy point to the same underlying array. Therefore, modifications to the elements of a slice inside a function are visible to the caller.

Example:
```go
func modifySlice(s []int) {
    s[0] = 100
}

arr := []int{1, 2, 3}
modifySlice(arr)
fmt.Println(arr) // Outputs: [100 2 3]
```

However, if you reassign the slice itself (e.g., `s = append(s, 4)`), the changes may not be reflected outside the function because the slice header is copied.

If you need to modify the slice's length or capacity and have those changes reflected outside the function, you can pass a pointer to the slice:

```go
func appendValue(s *[]int, value int) {
    *s = append(*s, value)
}

nums := []int{1, 2, 3}
appendValue(&nums, 4)
fmt.Println(nums) // Outputs: [1 2 3 4]
```