# Tricky Go Question: Running Multiple Files

This directory demonstrates a **tricky Go concept** that often confuses beginners: how to run a Go program that consists of multiple files in the same package.

## 🎯 The Tricky Question

**Q: How do you run a Go program that has multiple `.go` files in the same package?**

**A: You need to specify ALL the files when running the program!**

## 📁 File Structure

```
tricky_question/
├── main.go      # Contains the main() function
├── state.go     # Contains the printState() function
└── README.md    # This documentation
```

## 💻 Code Breakdown

### `main.go`
```go
package main

func main() {
    printState()
}
```

### `state.go`
```go
package main

import "fmt"

func printState() {
    fmt.Println("California")
}
```

## 🚨 The Tricky Part

### ❌ What DOESN'T Work
```bash
# This will FAIL with an error
go run main.go
```

**Error you'll see:**
```
./main.go:4:2: undefined: printState
```

### ✅ What DOES Work
```bash
# You must specify ALL files in the package
go run main.go state.go
```

**Output:**
```
California
```

## 🧠 Why This Happens

### Go Package Compilation Process

1. **Single File Compilation**: When you run `go run main.go`, Go only compiles `main.go`
2. **Missing Dependencies**: `main.go` calls `printState()` but Go doesn't know where it's defined
3. **Multiple File Compilation**: When you run `go run main.go state.go`, Go compiles both files together
4. **Function Resolution**: Go can now find the `printState()` function in `state.go`

### Key Concepts

- **Same Package**: Both files declare `package main`
- **Function Sharing**: Functions in the same package can call each other
- **Compilation Unit**: Go needs ALL files to resolve dependencies
- **No Forward Declarations**: Unlike C/C++, Go doesn't need function prototypes

## 🎓 Interview Questions

### 1. **Q: Why do you need to specify multiple files when running a Go program?**

**A:** Go compiles files individually by default. When `main.go` calls `printState()`, Go needs to know where this function is defined. By specifying both files, Go compiles them together and can resolve the function call.

### 2. **Q: What's the difference between `go run main.go` and `go run main.go state.go`?**

**A:** 
- `go run main.go`: Only compiles `main.go`, causing "undefined function" error
- `go run main.go state.go`: Compiles both files together, allowing function resolution

### 3. **Q: Can you have multiple files in the same package?**

**A:** Yes! Multiple files can share the same package name. They can call each other's functions, but you must compile them together.

### 4. **Q: What happens if you have a function in `state.go` that calls a function in `main.go`?**

**A:** It works! As long as both files are compiled together, functions can call each other regardless of which file they're in.

## 🔧 Alternative Solutions

### Option 1: Use `go run .`
```bash
# Run all .go files in the current directory
go run .
```

### Option 2: Use `go build`
```bash
# Build the program first
go build main.go state.go
# Then run the executable
./main
```

### Option 3: Use `go run *.go`
```bash
# Run all .go files (wildcard)
go run *.go
```

## 🎯 Key Takeaways

1. **Multiple Files = Multiple Compilation**: When using multiple files, specify them all
2. **Package Scope**: Functions in the same package can call each other
3. **No Forward Declarations**: Go doesn't need function prototypes
4. **Compilation Dependencies**: Go needs all files to resolve function calls
5. **Alternative Commands**: Use `go run .` or `go run *.go` for convenience

## 🚀 Try It Yourself

1. Navigate to this directory: `cd tricky_question`
2. Try the wrong way: `go run main.go` (watch it fail!)
3. Try the right way: `go run main.go state.go` (see it work!)
4. Try the convenient way: `go run .` (also works!)

## 💡 Pro Tips

- **Use `go run .`** for convenience when you have multiple files
- **Remember the rule**: Same package = compile together
- **This is different from other languages** like Java or C# where you can compile files separately
- **Go's simplicity** means you need to be explicit about what you're compiling

---

*This example demonstrates why understanding Go's compilation model is crucial for working with multi-file projects!*
