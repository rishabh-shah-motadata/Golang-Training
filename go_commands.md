# Go Run  
This section describes the internal working of the `go run` command.

## Steps executed by `go run`:

1. **Extract main files and arguments**  
   - Parse command-line arguments.  
   - Identify all `.go` files that belong to the main package.  
   - Separate remaining arguments to pass to the final binary.

2. **Create a temporary build directory**  
   - Generate a temporary folder where the compiled binary will be placed.

3. **Compile the program**  
   - Build the identified main Go files into a temporary binary.

4. **Execute the binary**  
   - Run the compiled executable.  
   - Attach terminal `stdin`, `stdout`, and `stderr` streams to the running process.
---

# Go Mod Init  
This section describes the internal working of the `go mod init <module>` command.

## Steps executed by `go mod init`:

1. **Determine the module path**  
   - Extract the module name from arguments or derive it from the directory.

2. **Detect the Go version**  
   - Read the currently installed Go version.

3. **Create an empty module file structure**  
   - Initialize an in-memory representation of `go.mod`.

4. **Insert module path**  
   - Add the `module <module-name>` statement.

5. **Insert Go version**  
   - Add the `go <version>` statement.

6. **Write the go.mod file**  
   - Save the constructed module information into a new `go.mod` file.
---

# Go Mod Tidy  
This section describes the internal working of the `go mod tidy` command.

## Steps executed by `go mod tidy`:

1. **Load the current module and packages**  
   - Read the existing `go.mod` and load all referenced source code packages.

2. **Build the import dependency graph**  
   - Collect all direct and indirect imports used in the project.

3. **Compare the import graph with go.mod**  
   - Identify missing dependencies.  
   - Identify unused dependencies.

4. **Update go.mod**  
   - Add missing modules.  
   - Remove unused modules.  
   - Ensure module definitions match the import graph.

5. **Update go.sum**  
   - Recalculate and update checksums for all required dependencies.
---
