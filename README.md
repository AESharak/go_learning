# Go Learning Journey - Stephen Grider Course

This repository contains my Go learning journey following Stephen Grider's Go course. Each directory represents different sections and concepts learned, with detailed explanations for interview preparation.

## 📚 Table of Contents

- [Project Structure](#project-structure)
- [Go Fundamentals](#go-fundamentals)
  - [Package Declaration](#package-declaration)
  - [Variable Declaration](#variable-declaration)
  - [Main Function](#main-function)
- [Code Examples](#code-examples)
- [Interview Preparation Notes](#interview-preparation-notes)

## 📁 Project Structure

```
go_stephen_grider/
├── hello_world/          # Basic Go program structure
├── cards/               # Variable declaration and assignment
├── .gitignore          # Git ignore configuration
└── README.md           # This file
```

## 🚀 Go Fundamentals

### Package Declaration

**Why `package main`?**

- `package main` is a special package declaration in Go
- It tells Go that this file contains the entry point of the program
- Only files with `package main` can have a `main()` function
- When you run `go run main.go` or `go build`, Go looks for the `main` package

**Difference from other packages:**

```go
// package main - executable package
package main
func main() {
    // Entry point of the program
}

// package anythingelse - library package
package utils
// No main() function allowed
// Used to create reusable code libraries
```

**Key Interview Points:**
- `package main` = executable program
- Other packages = libraries/reusable code
- Package name should match the directory name (convention)
- Only one `main()` function allowed per package

### Variable Declaration

Go offers multiple ways to declare variables:

**1. Long form (explicit type):**
```go
var card string = "Ace of Spades"
```

**2. Short form (type inference):**
```go
card := "Ace of Spades"  // Go infers the type
```

**3. Zero value initialization:**
```go
var card string  // Initialized to empty string ""
```

**Key Rules:**
- `:=` can only be used for **initial declaration**
- For reassignment, use `=`
- `:=` creates a new variable (can't redeclare)

**Interview Questions You Should Know:**
1. **Q: What's the difference between `var` and `:=`?**
   - `var` can be used anywhere, `:=` only for initial declaration
   - `var` allows zero-value initialization, `:=` requires assignment

2. **Q: When would you use explicit type vs type inference?**
   - Explicit: When type clarity is important for readability
   - Inference: For cleaner code when type is obvious

3. **Q: What happens if you try to use `:=` twice for the same variable?**
   - Compilation error: "No new variables on left side of :="

### Main Function

The `main()` function is:
- The entry point of every Go program
- Automatically called when the program starts
- Must be in the `main` package
- Cannot take parameters or return values

```go
func main() {
    // Program starts here
    fmt.Println("Hello, World!")
}
```

## 💻 Code Examples

### Hello World (`hello_world/main.go`)
```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

**Key Learning Points:**
- Basic program structure
- Import statement usage
- `fmt.Println()` for output

### Cards Example (`cards/main.go`)
```go
package main

import "fmt"

func main() {
    // Long form variable declaration
    // var card string = "Ace of Spades"
    
    // Short form with type inference
    card := "Ace of Spades"
    
    // Reassignment (note: no := here)
    card = "Five of Diamonds"
    
    fmt.Println(card)
}
```

**Key Learning Points:**
- Variable declaration methods
- Type inference vs explicit typing
- Variable reassignment rules
- Commenting best practices

## 🎯 Interview Preparation Notes

### Common Go Interview Questions

**1. What is the difference between `package main` and other packages?**
- `package main` creates executable programs
- Other packages create reusable libraries
- Only `main` package can have `main()` function

**2. Explain Go's variable declaration syntax**
- `var name type = value` - explicit type
- `name := value` - type inference (initialization only)
- `var name type` - zero value initialization

**3. What are Go's zero values?**
- `string` → `""`
- `int` → `0`
- `bool` → `false`
- `pointer` → `nil`

**4. Why does Go have both `var` and `:=`?**
- `var` for flexibility (can be used anywhere)
- `:=` for convenience (shorter syntax, type inference)

### Best Practices Demonstrated

1. **Use `:=` for initial declarations** when type is obvious
2. **Use explicit types** when clarity is important
3. **Comment your code** to explain different approaches
4. **Follow Go naming conventions** (lowercase for package-level variables)

## 🔄 Course Progress

- [x] **Section 1**: Basic Go Setup & Hello World
- [x] **Section 2**: Variable Declaration & Assignment


## 📝 Notes for Future Sections

As I progress through the course, I add new sections here with:
- Key concepts learned
- Code examples
- Interview questions
- Best practices

