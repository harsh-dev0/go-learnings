# Structs and Interfaces in Go

## Structs

- **Struct**: Custom data type that groups fields.
- **Example:**
    ```go
    type gasengine struct {
            gallons uint8
            mpg     uint8
    }
    ```

## Methods vs. Functions

### Version 1: Function with Struct Parameter
```go
func milesLeft(e gasengine) uint8 {
        return e.gallons * e.mpg
}
```
- **What:** Standalone function, takes struct as parameter.
- **Call:** `milesLeft(e)`
- **Gotchas:** Not attached to struct; less idiomatic for struct-specific logic.

### Version 2: Method on Struct
```go
func (e gasengine) milesLeft() uint8 {
        return e.gallons * e.mpg
}
```
- **What:** Method attached to struct type.
- **Call:** `e.milesLeft()`
- **Gotchas:** More idiomatic; enables chaining and encapsulation.

### Key Differences

| Aspect         | Standalone Function         | Method on Struct         |
|----------------|----------------------------|-------------------------|
| Ownership      | External                   | Belongs to struct       |
| Call syntax    | `milesLeft(e)`             | `e.milesLeft()`         |
| Readability    | Less expressive            | Cleaner, OOP-style      |
| Encapsulation  | Logic outside struct       | Logic inside struct     |

- **Pro Tip:**  
    - Use methods for struct-specific behaviour.
    - Use functions for generic logic.

    ## Interfaces

    - **Interface**: A type that specifies method signatures. Any type implementing those methods satisfies the interface.
    - **Why:** Enables polymorphism and abstraction.

    ### Example: Defining and Using an Interface

    ```go
    type engine interface {
        milesLeft() uint8
    }

    type gasengine struct {
        gallons uint8
        mpg     uint8
    }

    func (e gasengine) milesLeft() uint8 {
        return e.gallons * e.mpg
    }

    func printMilesLeft(e engine) {
        fmt.Println(e.milesLeft())
    }
    ```

    - **How it works:**
        - `engine` interface requires a `milesLeft()` method.
        - `gasengine` implements `milesLeft()`, so it satisfies `engine`.
        - Any type with a `milesLeft()` method can be used as an `engine`.

    ### See Also

    - See `struct.go` for more examples of structs and interfaces in action.

## Summary

- Use **methods** for struct-specific logic (idiomatic Go).
- Use **functions** for generic, decoupled logic.
- **Interfaces** enable polymorphism and abstraction.