# Generics in Go

If you've worked in a statically typed language before, you might have run into this problem: you have a function that sums up the values of an `int` slice, but later you want a similar function for a `float32` slice, and maybe for `float64` too. You end up repeating almost the same function for each type.

Different typed programming languages solve this problem in different ways. In Go, we use **generics** to collapse all this into a single function that handles all cases.

## Example: Summing Slices

Instead of writing separate functions for each type, you can write:

```go
func SumSlice[T int | float32 | float64](s []T) T {
    var sum T
    for _, v := range s {
        sum += v
    }
    return sum
}
```

Here, generics allow the function to receive additional parameters within square brackets. Instead of a value, we're passing in a type. This function will only accept `int`, `float32`, or `float64` types. Everywhere you see `T`, you can replace it in your head with `int`, `float32`, or `float64`.

You can call this same function for a slice of `float32`, for example, and it works.

## The `any` Type

Go also has the `any` type, which you can use in certain situations with generics. For example, you can't use it in the sum function above, because not all types are compatible with the addition operator (e.g., you can't add booleans). But here's an example where it makes sense:

```go
func IsEmpty[T any](s []T) bool {
    return len(s) == 0
}
```

This works no matter what the slice is made of.

Note: In these cases, you can even omit the square bracket type inputs, because the Go compiler is smart enough to infer the types. If you pass in a slice of `int`, then `T` is inferred to be `int`.

## When Type Inference Doesn't Work

There are cases where Go can't infer the type of your generic parameter. For example, if you're loading some JSON file and unmarshalling it into one of several struct types, you need to pass in the struct type so the function knows what to unmarshal into.

```go
func LoadJSON[T any](data []byte) (T, error) {
    var result T
    err := json.Unmarshal(data, &result)
    return result, err
}
```

Here, you need to specify the struct type because the function won't know what struct to unmarshal the JSON string into.

## Generics with Structs

Generics can also be used with struct types, not just functions. For example, if you have `GasEngine` and `ElectricEngine` types, you can create a `Car` type that has a generic `Engine` field:

```go
type Car[T Engine] struct {
    Make   string
    Model  string
    Engine T
}
```

Now you can instantiate two types of cars by setting the engine value to `GasEngine` or `ElectricEngine`.

---

Generics in Go are very simple and very useful—a good combination!