/*
Package main prints the message "Hello, World!"

package
  - Go code is organized into packages (like modules or libraries).
  - The first line of every source file is a package statement.
  - A package consists of one or more source files in a single directory.
  - The package name is same as the directory name.
  - However, package main is special here. It defines a standalone executable file and not a library.

import
  - Import tells the compiler what are other packages we reference. We need to import what is really
    required. Not less or not more. Go strictly checks for unused packages and raises errors.

func
  - Functions are defined using the func keyword. Like the main package, the main function is also
    special. The main() function in the main package is invoked automatically when running the
    executable. (Note: The init() function is also special that is invoked before main)

syntax
  - Go does require semicolons at the end of statements except when combining multiple statements
    in a single line.
  - In effect, Go should automatically add semicolons at the end of newlines. So, we should be
    careful when spliting the statements to multiple lines.
  - Certain cases allowed to continue the statement after a new line.
  - For instance, "return a + b" can be spread across multiple lines but you can add new line
    only after the "+" operator.
  - And for instance, the "{" must be in the same line of function declaration and in a new line.

statements
  - Statements changes the program states
  - In Go, i = i+1, or i += 1, or i++ are statements and not expressions
  - So, like C, we can not perform j = i++ which produces error.
  - Also, the prefix style like --i or ++j are not allowed.

expressions
  - Expressions focuses on producing some result and does not affect the program state
  - For instance, i * j, or fmt.Println() are expressions

comments:
  - Single line comments begins with "//" which is mostly used.
  - Multiline comments begins with "/*"" and ends with "*" followed by a "/" which is mostly used
    as package comments.

formatting
  - gofmt is the standard tool that can format the go source code
  - goimports (golang.org/x/tools/cmd/goimports) can add/remove import statements based on the code.
*/
package main

import "fmt"

// main displays the message "Hello, World!"
func main() {
	fmt.Println("Hello, World!")
}
