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

## Blank imports
  - It is an error if we import a package and don't use its name. So, we only import packages that
    we use in the source file.
  - But, there are situations to import them without really using its name. We might need their
    init() functions to be called or some of the initializers to be invoked by their package level
    variables.
  - In such cases, we can import a package but rename the import to blank identifier "_"
      *import _ "image/png" // registers PNG decoder*
  - The sample import is required to decode a PNG image using image.Decode() function. Otherwise,
    decoding can not understand the PNG format without importing png package.
  - Internally the init() in "image/png package registers the PNG decoder which is triggered by
    the blank import.
      *image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)*
  - There are many decoder packages available like PNG, so they are not included by default by the
    image. This can reduce the executable size and users can import only what is required.
  - Similarly the "database/sql" requires blank imports for the supported SQL drivers:
      *import _ "github.com/lib/pq" // enable support for Postgres*

## Packages and Naming
  - Keep the names short but not so short to be cryptic. Say "fmt", "io", "flag", "strings",
    "errors", "bufio", "os", "sort", "sync", "time", and "json"
  - Be descriptic and unambiguous where possible. For example, use like "ioutil" instead of "util".
  - Avoid common variable names, otherwise clients might have to rename your package during import.
  - Prefer singular form. Go packages "strings", "bytes", and "errors" are named plural to avoid
    hiding the type names "string", "byte", and "error".
  - Avoid package names that already have other connotations: For example, do not use "temp" for
    indicating temperature because "temp" is an universal term for temporary.
  - Naming the members of package is also as important. Ensure that combining the package name
    and member name provides the meanigful name. Say bytes.Equal(), http.Get() and json.Marshal().
    Avoid including the package name in the member itelf.
  - Packages like "math/rand" exposes one principle data type "Rand" and provides a New() function
    to create instances of Rand. Several methods can be added for Rand type.

## The Go Tool
  - The "go" is a combined set of tools for downloading, querying, formatting, building, testing,
    and installing go packages. So, go tool is a
  - Package manager like apt, rpm
  - Build system that compute dependencies, invoke compiler, assembler, and linker
  - Some of the frequently used commands:
    - build
    - clean
    - doc
    - env
    - fmt
    - get
    - install
    - list
    - run
    - test
    - version
    - vet
