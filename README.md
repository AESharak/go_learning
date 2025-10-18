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
- [x] **Section 4**: Slices, Append Function, and Iteration

---

## Section 4: Slices, Append Function, and Iteration

### New Concepts Learned

**Understanding Slices vs Arrays:**

Go has two data structures for handling lists of records:
- **Array**: Fixed length, primitive data structure
- **Slice**: Dynamic length, can grow/shrink, more features

**Key Differences:**
- Arrays have fixed length (must specify size at creation)
- Slices can grow and shrink dynamically
- Both must contain elements of the same type
- Slices are more commonly used in Go

**Slice Declaration Syntax:**
```go
// Create a slice of strings
cards := []string{"Ace of Diamonds", "Five of Diamonds"}

// Empty slice
var cards []string

// Slice with specific type
cards := []string{} // Empty slice of strings
```

**Important Rules:**
- All elements in a slice must be of the same type
- Use `[]type` syntax (singular type name, not plural)
- Slices are zero-indexed (first element is at index 0)

### The Append Function

**How Append Works:**
```go
// Original slice
cards := []string{"Ace of Diamonds", "Five of Diamonds"}

// Add new element - append returns a NEW slice
cards = append(cards, "Ace of Spades")
```

**Critical Points:**
- `append()` does NOT modify the original slice
- It returns a completely new slice
- You MUST assign the result back to your variable
- This is a fundamental concept in Go's design

**Common Mistake:**
```go
// WRONG - this won't work
append(cards, "New Card") // Result is discarded

// CORRECT - assign the result back
cards = append(cards, "New Card")
```

### Iterating Over Slices

**Range Keyword Syntax:**
```go
for i, card := range cards {
    fmt.Println(i, card)
}
```

**How Range Works:**
- `i` gets the index (0, 1, 2, etc.)
- `card` gets the actual value at that index
- `range` keyword tells Go to iterate over every element
- `:=` is used because we're declaring new variables each iteration

**Why Use `:=` in Loops:**
- Each iteration creates new `i` and `card` variables
- Previous values are discarded
- `:=` re-declares variables for each loop iteration

### Updated Code Examples

**Cards Example - Version 3 (`cards/main.go`)**
```go
package main

func main() {
    // Create a slice of strings called 'cards' with initial values
    // []string declares a slice that can hold multiple string values
    // We initialize it with two elements: a string literal and a function call
    cards := []string{"Ace of Diamongs", newCard()}

    // Iterate over the slice using 'range' keyword
    // 'i' gets the index (0, 1, 2, etc.) and 'card' gets the actual value
    // This is the first iteration, so we'll see the original 2 cards
    for i, card := range cards {
        println("first one", i, card)
    }

    // Use the append function to add a new element to the slice
    // append() takes the existing slice and the new element(s) to add
    // IMPORTANT: append() returns a NEW slice, it doesn't modify the original
    // We must assign the result back to 'cards' to update our slice
    cards = append(cards, "Ace of Spades")

    // Iterate over the slice again after adding the new card
    // Now we should see 3 cards total (original 2 + the new "Ace of Spades")
    for i, card := range cards {
        println("second time", i, card)
    }
}

// newCard function returns a string value
// The "string" after the parentheses tells Go this function returns a string type
func newCard() string {
    // Return a string value - this must match the declared return type
    return "Five of Diamonds"
}
```

**Expected Output:**
```
first one 0 Ace of Diamongs
first one 1 Five of Diamonds
second time 0 Ace of Diamongs
second time 1 Five of Diamonds
second time 2 Ace of Spades
```

### Key Learning Points

1. **Slice Declaration**: `[]string` creates a slice of strings
2. **Append Function**: Always returns a new slice, doesn't modify original
3. **Range Iteration**: `for i, value := range slice` iterates over all elements
4. **Variable Declaration in Loops**: Use `:=` because variables are re-declared each iteration
5. **Type Consistency**: All elements in a slice must be the same type

### Advanced Interview Questions

**9. What's the difference between an array and a slice in Go?**
- Arrays have fixed length, slices can grow/shrink
- Arrays: `[5]string` (fixed size 5)
- Slices: `[]string` (dynamic size)

**10. Why does append() return a new slice instead of modifying the original?**
- Go's design philosophy emphasizes immutability
- Prevents unexpected side effects
- Makes code more predictable and safer

**11. Can you mix different types in a slice?**
- No, all elements must be the same type
- `[]string` can only contain strings
- `[]int` can only contain integers

**12. What happens if you don't assign the result of append() back to a variable?**
- The new slice is created but immediately discarded
- Original slice remains unchanged
- You lose the new element

**13. How do you iterate over a slice without the index?**
```go
for _, card := range cards {
    fmt.Println(card) // Only print the value, ignore index
}
```

**14. What's the zero value of a slice?**
- `nil` - an uninitialized slice
- Length and capacity are both 0
- Can still call `append()` on a nil slice

### Best Practices Demonstrated

1. **Always assign append() results back to a variable**
2. **Use descriptive variable names in range loops**
3. **Comment complex slice operations for clarity**
4. **Initialize slices with meaningful default values**
5. **Use range for iteration instead of manual indexing**

### Common Pitfalls to Avoid

1. **Forgetting to assign append() result:**
   ```go
   // WRONG
   append(cards, "New Card")
   
   // CORRECT
   cards = append(cards, "New Card")
   ```

2. **Mixing types in slices:**
   ```go
   // WRONG - compilation error
   cards := []string{"Ace", 5}
   
   // CORRECT
   cards := []string{"Ace", "Five"}
   ```

3. **Using wrong syntax for slice declaration:**
   ```go
   // WRONG
   cards := []strings{"Ace"} // Should be 'string', not 'strings'
   
   // CORRECT
   cards := []string{"Ace"}
   ```

---

## Section 5: Custom Types and Receiver Functions

### New Concepts Learned

**Creating Custom Types:**

Go allows you to create custom types based on existing types. This gives you the ability to create specialized functions that only work with your custom type.

**Custom Type Syntax:**
```go
// Create a new type called 'deck' that behaves like a slice of strings
type deck []string
```

**Key Concepts:**
- Custom types "extend" or "borrow" all behavior from the base type
- You can replace the base type with your custom type anywhere
- Custom types are 100% equivalent to their base type
- This is NOT inheritance like in object-oriented languages

**Why Create Custom Types?**
- Allows you to create specialized functions (methods) for your type
- Makes code more readable and self-documenting
- Enables you to attach custom behavior to your data structures

### Receiver Functions (Methods)

**What are Receiver Functions?**

Receiver functions are special functions that "belong" to a custom type. They allow you to call functions using dot notation on variables of your custom type.

**Receiver Function Syntax:**
```go
func (receiverVariable receiverType) functionName() {
    // function body
}
```

**Key Components:**
- `(receiverVariable receiverType)` - This is the "receiver"
- `receiverVariable` - The variable name you use inside the function
- `receiverType` - The custom type this function belongs to
- `functionName` - The name of the method

**How It Works:**
```go
// Define custom type
type deck []string

// Create a method for deck type
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

// Usage in main function
func main() {
    cards := deck{"Ace of Diamonds", "Five of Diamonds"}
    cards.print() // Call the method using dot notation
}
```

### Updated Code Examples

**Deck Example (`cards_demo/deck.go`)**
```go
package main

import "fmt"

// Create a new type of deck, which is a slice of strings
// This custom type "extends" or "borrows" all the behavior of a slice of string
type deck []string

// This is a special function called a "receiver function" or "method"
// The syntax (d deck) before the function name is called a "receiver"
// This means the print() function belongs to the deck type
func (d deck) print() {
    // Loop through all cards in the deck and print each one
    for i, card := range d {
        fmt.Println(i, card)
    }
}
```

**Main Example (`cards_demo/main.go`)**
```go
package main

func main() {
    // Create a deck using our custom type
    // Notice we can still use all slice operations like append()
    cards := deck{"Ace of Diamonds", newCard()}
    
    // Call our custom method using dot notation
    cards.print()
    
    // We can still use append() because deck is equivalent to []string
    cards = append(cards, "Ace of Spades")
    
    // Call print() again to see the updated deck
    cards.print()
}

func newCard() string {
    return "Five of Diamonds"
}
```

**Expected Output:**
```
0 Ace of Diamonds
1 Five of Diamonds
0 Ace of Diamonds
1 Five of Diamonds
2 Ace of Spades
```

### Key Learning Points

1. **Custom Type Declaration**: `type name baseType` creates a new type
2. **Receiver Functions**: `func (variable type) methodName()` creates methods
3. **Method Calls**: Use dot notation to call methods on custom types
4. **Type Equivalence**: Custom types are 100% equivalent to their base types
5. **Multiple Files**: Use `go run file1.go file2.go` to run programs with multiple files

### Advanced Interview Questions

**15. What's the difference between a custom type and the base type it's built on?**
- Functionally, there is NO difference - they are 100% equivalent
- Custom types just allow you to create specialized methods
- You can use all the same operations on both types

**16. How do you create a method that belongs to a custom type?**
```go
func (variableName typeName) methodName() {
    // method body
}
```

**17. What is a receiver in Go?**
- The `(variableName typeName)` part before the function name
- It tells Go which type this function belongs to
- The variable name is how you access the instance inside the method

**18. Can you call methods on built-in types like int or string?**
- No, methods can only be created for custom types
- You cannot add methods to built-in Go types
- This is why we create custom types like `type deck []string`

**19. How do you run a Go program that has multiple .go files?**
```bash
go run main.go deck.go
```
- List all the .go files you want to include
- Go will compile and run them together

**20. What happens if you don't include all necessary .go files when running?**
- You'll get compilation errors about undefined types or functions
- All related .go files must be included in the same package

### Best Practices Demonstrated

1. **Create custom types for domain-specific concepts** (like `deck` for cards)
2. **Use descriptive receiver variable names** (like `d` for deck)
3. **Keep related functionality together** (methods with their types)
4. **Use dot notation for method calls** (`cards.print()`)
5. **Remember to include all files when running** (`go run file1.go file2.go`)

### Common Pitfalls to Avoid

1. **Forgetting to include all .go files when running:**
   ```bash
   # WRONG - will cause compilation errors
   go run main.go
   
   # CORRECT - include all necessary files
   go run main.go deck.go
   ```

2. **Using wrong receiver syntax:**
   ```go
   // WRONG - missing receiver variable
   func (deck) print() { }
   
   // CORRECT - include receiver variable
   func (d deck) print() { }
   ```

3. **Trying to add methods to built-in types:**
   ```go
   // WRONG - cannot add methods to built-in types
   func (s string) print() { }
   
   // CORRECT - create custom type first
   type myString string
   func (s myString) print() { }
   ```

### Updated Course Progress

- [x] **Section 1**: Basic Go Setup & Hello World
- [x] **Section 2**: Variable Declaration & Assignment
- [x] **Section 3**: Function Declaration & Return Types
- [x] **Section 4**: Slices, Append Function, and Iteration
- [x] **Section 5**: Custom Types and Receiver Functions

