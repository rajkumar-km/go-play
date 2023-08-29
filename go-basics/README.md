## go-basics
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

### Program Structure

#### Names
The names of Go variables, constants, functions, types, statement labels, and packages follows:
  - Must start with a letter or underscore
  - Can be in any length with letters, digits, and underscores. But go recommends shorter names.
  - Case sensitive: getName() and GetName() are two different functions
  - Go has 25 keywords which can not be used as names

##### Go keywords that can not be used as names
  - break
  - case
  - chan
  - const
  - continue
  - default
  - defer
  - else
  - fallthrough
  - for
  - func
  - go
  - goto
  - if
  - import
  - interface
  - map
  - package
  - range
  - return
  - select
  - struct
  - switch
  - type
  - var

#### Go predeclared names that can be reused
The following are the predeclared Go names that aren't reserved. This can be used in the program
but wherever it really makes sense. Beware of the confusions caused by redeclaration.
  - Constants - true, false, iota, nil
  - Types - int, int8, int16, int32, int64, uint, uint16, uint32, uint64, float32, float64,
    complex64, complex128, byte, rune, bool, string, error
  - Functions - new, make, len, cap, copy, append
    close, delete, complex, real, imag, panic, recover
