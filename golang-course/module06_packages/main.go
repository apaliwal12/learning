package main

import (
	"fmt"
	
	// We import our local sub-package. 
	// The path must include the module name from go.mod.
	"github.com/user/golang-course/module06_packages/geometry"
)

// ─────── 2. MODULES & PACKAGES ───────
// Historically, Go used `GOPATH` for dependency management, which was messy.
// Since Go 1.11, `go modules` is the standard.
//
// A Go Module is a collection of related Go packages that are released together.
// It is defined by a `go.mod` file at its root.
//
// A Go Package is a directory containing one or more `.go` files.
// All files in a package must declare the same `package name` at the top.
//
// Toolchain Commands you must know:
// `go mod init <module-name>`: Creates a go.mod file.
// `go mod tidy`: Adds missing modules and removes unused modules from go.mod/go.sum.
// `go get <path>`: Downloads a specific dependency.
// `go run .`: Compiles and runs the package in the current directory.
// `go build`: Compiles the package into an executable binary.

func main() {
	fmt.Println("Welcome to Module 6: Packages & Modules")

	// We can access exported types from the geometry package.
	c := geometry.Circle{Radius: 5.0}
	
	// We CANNOT access unexported fields:
	// c.cachedArea = 100 // COMPILER ERROR!
	
	fmt.Printf("Circle radius: %.2f\n", c.Radius)
	fmt.Printf("Circle area: %.2f\n", c.Area())

	// Using the constructor
	c2 := geometry.NewCircle(10.0)
	fmt.Printf("Circle 2 area: %.2f\n", c2.Area())
}
