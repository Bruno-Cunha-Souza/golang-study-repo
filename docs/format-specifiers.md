# Format Specifiers in Go

Complete guide on format specifiers used with `fmt.Printf`, `fmt.Sprintf`, and `fmt.Fprintf`.

## Table of Contents

- [Basic Specifiers](#basic-specifiers)
- [Number Formatting](#number-formatting)
- [String Formatting](#string-formatting)
- [Advanced Formatting](#advanced-formatting)
- [Formatting Flags](#formatting-flags)
- [Practical Examples](#practical-examples)

## Basic Specifiers

### General Types

| Specifier | Description                       | Example                                             |
| --------- | --------------------------------- | --------------------------------------------------- |
| `%v`      | Default value (any type)          | `fmt.Printf("%v", 42)` → `42`                       |
| `%+v`     | Value with field names (structs)  | `fmt.Printf("%+v", person)` → `{Name:Bruno Age:25}` |
| `%#v`     | Go-syntax representation of value | `fmt.Printf("%#v", "hello")` → `"hello"`            |
| `%T`      | Type of variable                  | `fmt.Printf("%T", 42)` → `int`                      |
| `%%`      | Literal percent sign              | `fmt.Printf("%%")` → `%`                            |

### Boolean

| Specifier | Description   | Example                           |
| --------- | ------------- | --------------------------------- |
| `%t`      | true or false | `fmt.Printf("%t", true)` → `true` |

### Strings and Bytes

| Specifier | Description                    | Example                                 |
| --------- | ------------------------------ | --------------------------------------- |
| `%s`      | String                         | `fmt.Printf("%s", "hello")` → `hello`   |
| `%q`      | Double-quoted string           | `fmt.Printf("%q", "hello")` → `"hello"` |
| `%x`      | Hexadecimal, lowercase letters | `fmt.Printf("%x", "Go")` → `476f`       |
| `%X`      | Hexadecimal, uppercase letters | `fmt.Printf("%X", "Go")` → `476F`       |

## Number Formatting

### Integers

| Specifier | Description                     | Example                            |
| --------- | ------------------------------- | ---------------------------------- |
| `%d`      | Decimal (base 10)               | `fmt.Printf("%d", 42)` → `42`      |
| `%b`      | Binary (base 2)                 | `fmt.Printf("%b", 42)` → `101010`  |
| `%o`      | Octal (base 8)                  | `fmt.Printf("%o", 42)` → `52`      |
| `%O`      | Octal with `0o` prefix          | `fmt.Printf("%O", 42)` → `0o52`    |
| `%x`      | Hexadecimal lowercase (base 16) | `fmt.Printf("%x", 42)` → `2a`      |
| `%X`      | Hexadecimal uppercase (base 16) | `fmt.Printf("%X", 42)` → `2A`      |
| `%c`      | Unicode character               | `fmt.Printf("%c", 65)` → `A`       |
| `%U`      | Unicode format                  | `fmt.Printf("%U", 'A')` → `U+0041` |

### Float and Complex

| Specifier | Description                             | Example                                     |
| --------- | --------------------------------------- | ------------------------------------------- |
| `%f`      | Float (6 decimal places default)        | `fmt.Printf("%f", 3.14)` → `3.140000`       |
| `%.2f`    | Float with 2 decimal places             | `fmt.Printf("%.2f", 3.14159)` → `3.14`      |
| `%e`      | Scientific notation lowercase           | `fmt.Printf("%e", 1234.5)` → `1.234500e+03` |
| `%E`      | Scientific notation uppercase           | `fmt.Printf("%E", 1234.5)` → `1.234500E+03` |
| `%g`      | `%e` or `%f`, whichever is more compact | `fmt.Printf("%g", 1234.5)` → `1234.5`       |
| `%G`      | `%E` or `%f`, whichever is more compact | `fmt.Printf("%G", 1234.5)` → `1234.5`       |

### Pointers

| Specifier | Description    | Example                                 |
| --------- | -------------- | --------------------------------------- |
| `%p`      | Memory address | `fmt.Printf("%p", &x)` → `0xc000014088` |

## String Formatting

### Width and Alignment

```go
// Minimum width of 10 characters (right-aligned)
fmt.Printf("%10s\n", "Go")        // "        Go"

// Minimum width of 10 characters (left-aligned)
fmt.Printf("%-10s\n", "Go")       // "Go        "

// Numbers with minimum width
fmt.Printf("%5d\n", 42)           // "   42"
fmt.Printf("%-5d\n", 42)          // "42   "
```

### Precision

```go
// Float with specific precision
fmt.Printf("%.2f\n", 3.14159)     // "3.14"
fmt.Printf("%.4f\n", 3.14159)     // "3.1416"

// Truncated string
fmt.Printf("%.5s\n", "Hello World")  // "Hello"
```

### Combining Width and Precision

```go
// 10 characters width, 2 decimal places
fmt.Printf("%10.2f\n", 3.14159)   // "      3.14"

// 10 characters, 3 digits from string
fmt.Printf("%10.3s\n", "Hello")   // "       Hel"
```

## Formatting Flags

| Flag        | Description                   | Example                            |
| ----------- | ----------------------------- | ---------------------------------- |
| `+`         | Always print sign for numbers | `fmt.Printf("%+d", 42)` → `+42`    |
| `-`         | Left-align                    | `fmt.Printf("%-5d", 42)` → `42   ` |
| `#`         | Alternate format              | `fmt.Printf("%#x", 42)` → `0x2a`   |
| `0`         | Pad with leading zeros        | `fmt.Printf("%05d", 42)` → `00042` |
| ` ` (space) | Space for positive numbers    | `fmt.Printf("% d", 42)` → ` 42`    |

## Practical Examples

### Formatting Tables

```go
package main

import "fmt"

func main() {
    fmt.Printf("%-10s | %5s | %8s\n", "Name", "Age", "Salary")
    fmt.Printf("%-10s-|-%5s-|-%8s\n", "----------", "-----", "--------")
    fmt.Printf("%-10s | %5d | %8.2f\n", "Alice", 30, 5500.50)
    fmt.Printf("%-10s | %5d | %8.2f\n", "Bob", 25, 4200.00)
    fmt.Printf("%-10s | %5d | %8.2f\n", "Charlie", 35, 6800.75)
}
```

Output:

```
Name       |   Age |   Salary
-----------|-------|----------
Alice      |    30 |  5500.50
Bob        |    25 |  4200.00
Charlie    |    35 |  6800.75
```

### Formatting Monetary Values

```go
price := 1234.567
fmt.Printf("Price: $%.2f\n", price)           // Price: $1234.57
fmt.Printf("Price: $%10.2f\n", price)         // Price: $   1234.57
fmt.Printf("Price: $%010.2f\n", price)        // Price: $0001234.57
```

### Debugging with Structs

```go
type Person struct {
    Name   string
    Age    int
    Height float32
}

p := Person{Name: "Bruno", Age: 25, Height: 1.85}

fmt.Printf("%v\n", p)     // {Bruno 25 1.85}
fmt.Printf("%+v\n", p)    // {Name:Bruno Age:25 Height:1.85}
fmt.Printf("%#v\n", p)    // main.Person{Name:"Bruno", Age:25, Height:1.85}
```

### Base Conversions

```go
num := 255

fmt.Printf("Decimal: %d\n", num)      // Decimal: 255
fmt.Printf("Binary: %b\n", num)       // Binary: 11111111
fmt.Printf("Octal: %o\n", num)        // Octal: 377
fmt.Printf("Hex: %x\n", num)          // Hex: ff
fmt.Printf("Hex: %X\n", num)          // Hex: FF
fmt.Printf("Hex: %#x\n", num)         // Hex: 0xff
```

### Unicode and Characters

```go
char := 'A'

fmt.Printf("Char: %c\n", char)        // Char: A
fmt.Printf("Unicode: %U\n", char)     // Unicode: U+0041
fmt.Printf("Value: %d\n", char)       // Value: 65
```

### Argument Indices

```go
// Use the same argument multiple times
fmt.Printf("%[1]d in decimal = %[1]b in binary = %[1]x in hex\n", 42)
// 42 in decimal = 101010 in binary = 2a in hex

// Reorder arguments
fmt.Printf("%[2]s %[1]s\n", "World", "Hello")
// Hello World
```

## fmt Family Functions

```go
// Printf - prints to standard output
fmt.Printf("Age: %d\n", 25)

// Sprintf - returns formatted string
str := fmt.Sprintf("Age: %d", 25)

// Fprintf - writes to an io.Writer
fmt.Fprintf(os.Stderr, "Error: %s\n", "file not found")

// Errorf - creates a formatted error
err := fmt.Errorf("failed to open file: %s", filename)
```

## References

- [Official fmt documentation](https://pkg.go.dev/fmt)
- [Go by Example: String Formatting](https://gobyexample.com/string-formatting)
