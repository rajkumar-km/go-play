# Package and the Go Tool

Go packages are the separate units of code like libraries or modules:
  - Serves modularity, encapsulation, separate compilation, reusability, and maintainability.

Modularity:
  - Provides separate namespace within every package. So the same name can be reused.
    in other packages to reduces naming conflicts.
  - Multiple files can be created inside the package folder. They can access information
    each other as like they are defined in the same file.

Reusability:
  - DRY (Don't Repeat Yourself) principle is one of guidance of good quality software
    Packages are the next steps (after functions) for code.

Encapsulation
  - This also provides the way to hide/expose information outside the package.
  - Anything (variable, type, func) that starts with the Capital letter is accessible after the import
  - Anything that starts with the small letter are private to the package
  - Developers are free to change unexported private members without worrying about external usage.

Speed up the compilation:
  - Three main things about compilation that makes Go faster.
    1. All imports must be listed explicitly on top of source file. So, the compiler need not scan
       the entire source file for its dependencies.
    2. The dependencies of packages form a directed acyclic graph. Since there are no cycles,
       packages can be compiled separatly in parallel.
    3. Records the export information in the object file for the current and its dependency
       packages. The compiler needs to look one object file for every import and not beyond that.

## Importing packages:
  - Go convention is to have the directory name same as package name
  - For example "rand" is the nested package name on import of the "math/rand"
  - External packages can be imported with path like
    "github.com/rajkumar-km/go-play/go-basics/07-packages/tempconv"
  - Go does not define any standards to this path, it upto the tools to intrepret it.
    The "go" tool interpret the path as folders.
  - A short name is assigned to access the package. By default it is the package name
    generally available at the last part of the import. We can also specificy the alternative
    name to avoid naming conflicts when importing multiple packages with the same name.
  - All the external packages starts with the domain that owns or hosts the package making it
    globally unique.

## Package declaration
  - Every source file starts with a package statement on top is called package declaration.
  - Generally the last part of the import path contains the package name, but there are three
    exceptions:
    1. A package defining a command (or Go executable) has the package name "main". This would signal
       the "go build" tool to use the linker and build an executable.
    2. Generally the "_test.go" files can use separate package name ending with "_test". So, this
       allows a single directory to have two different packages. The _test package is specifically
       for "go test" tool which would both the actual and test package before running the test.
       Separate _test packages are used for tests to avoid cycles in the import path.
    3. Some tools suggest to append version number in the package import paths
       (like "gopkg.in/yaml.v2"). The package name excludes the version suffix and use "yaml".

## Import declarations
  - Go source file contains zero or more import declarations immediately followed by the
    package declaration.
  - Multiple import packages can be grouped as like var and const blocks.
  - Alternative names can be provided for the imported package to avoid naming conflicts.
  - Renamed imports can also be used to provide shorter names, but it can be used consistently
    across all source files to avoid confusion.
  - Each import creates a dependency to other package. The go build tool reports an error if
    there is a dependency cycle.