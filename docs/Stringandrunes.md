# Strings and Runes in Go

## Strings

- A string in Go is a read-only slice of bytes.
- Strings are immutable; once created, they cannot be changed.
- The `len()` function returns the number of bytes in the string, **not** the number of characters.
- To count the number of characters (runes) in a string, use `utf8.RuneCountInString()` from the `unicode/utf8` package.
- Example:
    ```go
    import "unicode/utf8"

    s := "Go语言"
    count := utf8.RuneCountInString(s)
    fmt.Println(count) // Output: 4
    ```
- This function correctly counts Unicode characters, not just bytes.

```go
s := "Go语言"
fmt.Println(len(s)) // Output: 8 (because of multi-byte UTF-8 characters)
```

## Runes

- A rune is an alias for `int32` and represents a Unicode code point(A Unicode code point is just a number that uniquely identifies a character in the Unicode standard).
- To get the number of runes (characters) in a string, use `utf8.RuneCountInString()` from the `unicode/utf8` package.

```go
import "unicode/utf8"

s := "Go语言"
fmt.Println(utf8.RuneCountInString(s)) // Output: 4
```

## int8 vs int32

- `int8` is an 8-bit signed integer (range: -128 to 127).
- `int32` is a 32-bit signed integer (range: -2,147,483,648 to 2,147,483,647).
- Runes use `int32` to store Unicode code points, allowing for a wide range of characters.

## Efficient String Building

- Concatenating strings repeatedly with `+` is inefficient due to repeated memory allocations.
- Use `strings.Builder` for efficient string construction.

```go
import "strings"

var builder strings.Builder
builder.WriteString("Hello, ")
builder.WriteString("world!")
result := builder.String()
```

- `strings.Builder` minimizes memory allocations and improves performance when building large or complex strings.
