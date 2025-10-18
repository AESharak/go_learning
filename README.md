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
- [Section 3: Function Declaration & Return Types](#section-3-function-declaration--return-types)
- [Updated Code Examples](#updated-code-examples)
- [Additional Interview Questions](#additional-interview-questions)

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

---

## Section 3: Function Declaration & Return Types

### New Concepts Learned

**Function Declaration with Return Types:**

Go requires explicit return type declaration for functions. Unlike variables where Go can infer types, functions must explicitly declare what they return.

**Basic Function Syntax:**
```go
func functionName() returnType {
    // function body
    return value
}
```

**Key Concepts:**
- Functions must declare their return type explicitly
- Go is very strict about type matching
- Return type comes after the parentheses: `func name() string`
- If no return value, omit the return type: `func name()`

**Common Return Types:**
- `string` - text data
- `int` - whole numbers
- `float64` - decimal numbers
- `bool` - true/false values

**Error Handling:**
- Go's error messages are very descriptive
- "too many arguments to return" means return type mismatch
- Always match declared return type with actual return value

**Key Interview Points:**
- Go requires explicit return type declaration
- Type inference works for variables, not function returns
- Error messages help identify type mismatches
- Functions can return any Go type

## Updated Code Examples

### Cards Example - Version 2 (`cards/main.go`)
```go
package main

import "fmt"

func main() {
    // Call the newCard function and assign its return value to card
    // Go infers that card is of type string because newCard() returns a string
    card := newCard()
    
    // Print the card value to the console
    fmt.Println(card)
}

// newCard function returns a string value
// The "string" after the parentheses tells Go this function returns a string type
func newCard() string {
    // Return a string value - this must match the declared return type
    return "Five of Diamonds"
}
```

**Key Learning Points:**
- Function declaration with return types
- Type inference for variables from function calls
- Explicit return type declaration requirement
- Function call and assignment syntax
- Go's strict type checking system

## Additional Interview Questions

**5. How do you declare a function that returns a string?**
```go
func functionName() string {
    return "some string"
}
```

**6. What happens if you don't declare a return type for a function?**
- Go assumes the function returns nothing
- If you try to return a value, you'll get "too many arguments to return" error

**7. Can Go infer return types like it does with variables?**
- No, Go requires explicit return type declaration
- This is different from variable type inference with `:=`

**8. What's the difference between function return types and variable type inference?**
- Variables: Go can infer types with `:=`
- Functions: Must explicitly declare return types
- This ensures type safety at function boundaries

## Updated Course Progress

- [x] **Section 1**: Basic Go Setup & Hello World
- [x] **Section 2**: Variable Declaration & Assignment
- [x] **Section 3**: Function Declaration & Return Types

