# go-basics
Examples to demonstrate the core Go programming concepts
1. Hello, World
2. Variables
3. Types
4. Constants
5. Control Flow
6. Functions
7. Packages
8. Arrays
9. Slices
10. Maps
11. Pointers
12. Structures
13. Methods
14. Interfaces
15. Composition
16. Concurrency
17. Reflection
18. Generics

## Program Structure

### 1. Names
The names of Go variables, constants, functions, types, statement labels, and packages:
  - Must start with a letter or underscore
  - Can be in any length with letters, digits, and underscores. But Go recommends shorter names.
  - Case sensitive: getName() and GetName() are two different functions
  - Go has 25 keywords that can not be used as names
  - Go programmers use the "camel case" style for names.
  - The acronyms are left as it is in the capital. For example, escapeHTML() is valid and not escapeHtml()

#### Go keywords that can not be used as names
  - break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto
  - if, import, interface, map, package, range, return, select, struct, switch, type, var

#### Go predeclared names that can be reused
The following are the predeclared Go names that aren't reserved. This can be used in the program
but wherever it really makes sense. Please be aware of the confusion caused by redeclaration.
  - Constants - true, false, iota, nil
  - Types - int, int8, int16, int32, int64, uint, uint16, uint32, uint64, float32, float64,
    complex64, complex128, byte, rune, bool, string, error
  - Functions - new, make, len, cap, copy, append
    close, delete, complex, real, imag, panic, recover
