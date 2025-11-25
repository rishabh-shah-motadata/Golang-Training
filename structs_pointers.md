# Definitive Guide: Using Structs With and Without Pointers in Go

This document summarizes **why**, **when**, and **how** to use value types vs pointer types in Go—based on Go community best practices, performance behavior, and language semantics.

---

# 1. Struct Fields: Value (`T`) vs Pointer (`*T`)

## Use Value Fields by Default

Value fields should be your default choice in almost all cases.

### Why Use Value Fields?

#### 1. **Safety — No nil checks**
Value fields always have a zero value, so you never need to defend against `nil`:

```go
type User struct {
    Name string // always "", never nil
}
```

This eliminates an entire class of runtime panics.

#### 2. **Ownership is Clear**
The struct *owns* the data. The lifetime of the field matches the parent struct.

#### 3. **Performance — Better cache locality**
Value fields are stored *inside* the struct, making memory contiguous and easier for CPUs to fetch.

#### 4. **Lower Garbage Collector (GC) Pressure**
No pointer → no GC tracing required.
This improves GC speed, especially with large arrays of structs.

### Example: Use Value Fields

```go
type Employee struct {
    ID   int
    Name string
}
```

The fields are small, safe, and simple to work with.

---

## When to Use Pointer Fields (`*T`)

Pointer fields should be used *only when needed*.

### 1. Representing Optional Data
If a field can be “missing,” a value field cannot express this.

Example: You cannot distinguish “not provided” from actual zero value.

```go
type User struct {
    DateOfBirth *time.Time // nil → not provided
}
```

Perfect for:

- JSON APIs
- Database models
- Partial updates

### 2. Sharing a Single Instance
If multiple structs should point to the *same* data:

```go
type Department struct {
    Name string
}

type Employee struct {
    Dept *Department
}
```

Modifying `Dept.Name` updates it for every employee.

### 3. Avoid Copying Large Structs
If a struct field contains a large nested struct (many KB), copying would be expensive.

Pointer fields avoid unnecessary duplication.

### Example of Pointer Field Use

```go
type Config struct {
    Timeout *int // optional
}
```

---

# 2. Passing Structs: Value vs Pointer

## Pass by Value (Default)

```go
func Process(u User) { /* works on a copy */ }
```

### Why Pass by Value?

#### 1. **Immutability**
The function receives a copy → no side effects.

#### 2. **Performance for Small Structs**
Small structs (<64 bytes) are faster to copy than pointers.

#### 3. **Stack Allocation → No GC Cost**
Passing by value often keeps structs on the stack, avoiding heap allocation.

### Example

```go
func PrintUser(u User) {
    fmt.Println(u.Name) // safe read, no mutation
}
```

---

## Pass by Pointer

### 1. When You Need to Mutate the Original

```go
func (u *User) SetName(name string) {
    u.Name = name
}
```

### 2. When the Struct Is Large
Passing pointers (8 bytes) is cheaper than copying large structs.

### Trade-off  
Pointers may cause the struct to “escape to the heap,” increasing GC load.

---

# 3. Method Receivers: Value vs Pointer

Go encourages **consistency**.

If one method needs a pointer receiver, use pointer receivers for all methods.

## Use Pointer Receiver When:
- Method needs to modify the receiver
- Struct is large
- Other methods already use pointer receivers

Example:

```go
type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func (c *Counter) Get() int {
    return c.Value
}
```

## Use Value Receiver When:
- Struct is small
- Method is read-only
- You want copy semantics

Example:

```go
func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

---

# 4. Practical Examples

## Example 1: Optional Fields (Correct Use of Pointer)
```go
type Product struct {
    Name        string
    Description *string // optional
}
```

## Example 2: Copying vs Mutating
```go
type Person struct {
    Name string
}

func (p Person) RenameWrong(newName string) {
    p.Name = newName // modifies only the copy
}

func (p *Person) RenameRight(newName string) {
    p.Name = newName // modifies original
}
```

---

# 5. Best Practices

### Use Values by Default
Simpler, safer, faster.

### Use Pointers Only When Needed
- Data is optional
- Data must be shared
- Struct is large
- You need mutation

### Be Consistent With Method Receivers
Pointer → all pointer.  
Value → all value.

---

# 6. The Shortest Possible Summary

- **Use values unless you need nil, sharing, mutation, or performance reasons.**
- **Use pointers intentionally, not habitually.**

# 7. Why You Should Not Use Pointers in Go (Unless Necessary)

This document explains the key reasons **why pointers should *not* be used by default in Go**, and why value types are usually the safer, faster, and more idiomatic choice.

Pointers are powerful—but they come with real costs. Use them only when you have a clear, intentional reason.

---

### 7.1 Pointers Increase Garbage Collector (GC) Pressure

Every pointer inside a struct forces the GC to **trace** that reference.

- More pointers → more GC work  
- More GC work → longer GC cycles  
- Longer GC cycles → slower overall performance  

Example:

```go
type User struct {
    Name *string // pointer → GC must scan
}
```

vs.

```go
type User struct {
    Name string // value → GC skips scanning
}
```

Structs containing only value fields become **GC leaf objects**, meaning the GC can scan them quickly and move on.

---

# 7.2 Pointers Force Heap Allocation (Escape Analysis)

Using pointers increases the likelihood that the Go compiler will decide:

> “This value must escape to the heap.”

Heap allocations are:
- slower to allocate  
- slower to free  
- more GC work  

Example: heap allocation due to returning a pointer.

```go
func NewUser() *User {
    return &User{} // may escape to the heap
}
```

Value-returning version:

```go
func NewUser() User {
    return User{} // often stays on the stack
}
```

Stack allocations are fast and have zero GC cost, which is why value types often outperform pointer-heavy designs.

---

# 7.3 Risk of Nil Pointer Panics

When using pointers, you must always check for nil:

```go
if user.Name != nil {
    fmt.Println(*user.Name)
}
```

Without the check:

```go
fmt.Println(*user.Name) // panic: invalid memory address or nil pointer dereference
```

Value fields eliminate this entire category of bugs:

```go
fmt.Println(user.Name) // always safe, always a valid string
```

Zero values make Go safer—pointers remove that safety.

---

# 7.4 Harder to Serialize (JSON, DB Models, APIs)

Value fields serialize predictably:

```json
{
  "name": ""
}
```

Pointer fields complicate things:

- `nil` vs `""`
- `omitempty` behaves differently
- `null` appears in JSON
- Partial updates become tricky

Example:

```go
type Product struct {
    Description *string `json:"description,omitempty"`
}
```

When serialized:
- `nil` → field omitted  
- pointer to empty string → field included  

This is useful **only when intentional**; otherwise, it adds complexity.

---

# 7.5 Bad for Small Structs

For small structs (≈ up to 64 bytes), pointers are usually **slower** than values.

Why?
- Dereferencing pointers causes extra CPU work  
- Pointers scatter data across the heap → cache misses  
- Copying small structs is extremely cheap (just a few bytes)  

Example small struct:

```go
type Point struct {
    X, Y int
}
```

Passing by value is faster and safer than using:

```go
*p Point
```

Pointers harm performance unless dealing with *large* structs.

---

# 8. Final Summary

Use **values by default**.

Use pointers **only when necessary**, such as:
- optional data  
- shared references  
- mutation of caller data  
- very large structs  
- concurrency primitives (mutex, atomic values)  

---

# 9. Quick Decision Table

| Avoid Pointers Because… | Explanation |
|--------------------------|-------------|
| GC Pressure | More pointers → more GC scanning → slower system |
| Heap Allocation | Pointers can force values to escape to the heap |
| Nil Panics | Pointers require nil checks or risk panics |
| Serialization Issues | `nil` vs zero-value causes complexity |
| Poor for Small Structs | Slower, worse locality, unnecessary overhead |

---

Pointers are a tool—use them intentionally, not by habit.
