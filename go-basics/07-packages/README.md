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
    and installing go packages. So, go tool is a package manager like apt, and rpm. Also, a build
    system that compute dependencies, invoke compiler, assembler, and linker.
  - This is similar to the design of Swiss army knife to perform multiple tasks.
  - Some of the frequently used commands:
    - build    - compile packages and dependencies
    - clean    - remove object files
    - doc      - show documentation for package or symbol
    - env      - print Go environment variables
    - fmt      - run gofmt and update source files
    - get      - download and install packages and its dependencies
    - install  - compile and install packages and its dependencies
    - list     - list packages in local/remote, including custom lists such as dependencies with -f
    - run      - compile and run Go program
    - test     - test packages
    - version  - print Go version
    - vet      - run go vet on packages to identify potential issues in the program.
  - To keep the configurations minimum, Go relies on convensions. For example, a package name can
    be identified by its underlying directory. The import path also includes the repository of the
    package.
### Workspace Organization
  - Run "go env" to get the complete list of go environment configurations.
  - GOPATH environment variable is the one every developers need to know.
  - This refers to root of our workspace which contains three subdirectories:
    1. src/ 
       - All our source code resides here with subdirectories for each repo
       - Example: src/github.com/rajkumar-km/go-play/go-basics/01-hello/hello.go
    2. pkg/
       - Stores compiled packages here
       - Example: $GOPATH/pkg/mod/github.com/rajkumar-km/go-play@v0.0.0-20230729032233-5ceac969bcd6
    3. bin/
       - Executable binary files are stored here
  - GOROOT specifies the root directory of Go distribution which contains all Go standard packages.
  - GOOS specify the target operating system and GOARCH is the target architecture to compile for.
### Downloading Packages
  - The import path not only refers the local path but it can also locate the remote repository.
  - The "go get" command download the package and its dependencies, build, and install them.
  - Examples:
    go get golang.org/x/lint
    go get -u golang.org/x/lint # -u flag to get latest version and update its dependencies as well
    go get -v github.com/some-repo/sub/... # use ... to install all packages in a subtree
  - This supports popular repositories like github.com, bitbucket, and lanuchpad. For other sites,
    we need to specify the protocol use in import path (such as git or mecurial).
  - The import path need not be a actual hosting repository. Instead the import path can contain a
    metadata that can redirect to a actual repository. Go tool supports redirection with metadata.
### Building Packages
  - The go build command compiles the specified package
  - If the package is a library, the result is ignored but prints compile errors if any. This is
    just to ensure that library is free of compile errors.
  - If the package is main, Go invokes the linker and generates the executable file. The executable
    file is named by the last segment of path.
  - One or more files can also be specified for go build in which case executable name is the
    basename of first file.
  - Examples:
    go build github.com/rajkumar-km/go-play/go-basics/01-hello # produces exe "01-hello"
    cd go-play/ && go build ./go-basics/01-hello               # produces exe "01-hello"
    cd go-play/ && go build go-basics/01-hello                 # not supported without "./" prefix
    cd go-play/ && go build ./go-basics/01-hello/hello.go      # produces exe "hello"
  - The go run command is handy to quickly build and run the programs. 
    go run ./echo/echo.go "hello"
    go run ./echo/echo.go -- "sample.go" # use -- to pass any ".go" string as arguments
  - go build compiles the requested packages and all its dependencies by default. It is faster for
    smaller projects, but it can slow down when the number of packages and lines of code increases.
  - The go install command comes to the rescue. It is similar to go build except it can save the
    compiled code under GOPATH/pkg and GOPATH/bin for each packages instead of throwing away.
    Thereafter go build or go install won't compile the package unless there are changes.
  - "go build -i" is equal to "go install" command to build and save the compiled code.
  - Cross platform compilation is easy in Go by setting the GOOS and GOARCH environment variables.
  - Sometime, we may have files that should be built only on specific platform.
  - Go tool also scans for OS and archiecture names in the filename and builds only on the
    particular OS or architecture. For example, "net_linux.go" will be build only if the GOOS is
    linux. or net_amd64.go is build only if GOARCH is amd64.
  - We can also use the special package comments to control the build.
  - To build only on linux and mac os
    // +build linux darwin
    package net
  - To ignore during the build
    // +build ignore
    package dummy
  - See go doc go/build for more details.
