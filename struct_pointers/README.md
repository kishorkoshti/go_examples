# Go Struct with Pointers Example

This repository demonstrates how to use **structs with pointers** in Go,  
mutate values through pointers, and write unit tests.

## 📂 Structure
- `examples/struct_pointers/main.go` → Runnable example
- `examples/struct_pointers/struct_pointers.go` → Pointer struct logic
- `examples/struct_pointers/struct_pointers_test.go` → Unit tests

## 🚀 Run the Example
```bash
go run examples/struct_pointers/main.go
```
Expected output:
```
Hello, Charlie!
Charlie is now 41 years old.
```

## 🧪 Run Tests
```bash
go test ./examples/struct_pointers
```

## 🔑 Key Points
- **Pointers in structs** allow changing the original data without copying.
- Use `&value` to get the pointer, and `*pointer` to dereference.
- Methods with pointer receivers (`func (u *User) ...`) can mutate struct fields.
