package main

import (
	"fmt"
	"runtime"
)

// ─────── 1. GO BUILD SYSTEM ───────
// Go's build system is incredibly fast and produces single, statically linked
// binaries. This means you don't need Go installed on the server to run the app!
//
// Common Commands:
// `go build -o myapp main.go` -> Compiles the code into an executable named `myapp`
// `GOOS=linux GOARCH=amd64 go build` -> Cross-compiles for a Linux server (from Mac/Windows!)

// ─────── 2. BUILD TAGS (Conditional Compilation) ───────
// Sometimes you need different code for different environments (e.g., Windows vs Linux,
// or "Free" vs "Pro" versions).
// 
// You can use special comments at the VERY TOP of the file (before the package declaration)
// like `//go:build pro` or `//go:build windows`.
// The file will ONLY be compiled if that tag is active.
//
// Example: `go build -tags pro`

func demonstrateBuildInfo() {
	fmt.Println("\n--- Build System Info ---")
	
	// The `runtime` package gives us info about where the binary is CURRENTLY running.
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Go Version: %s\n", runtime.Version())
}

// ─────── 3. LINKER FLAGS (ldflags) ───────
// You can inject variables into your binary AT COMPILE TIME.
// This is how companies bake the "Version" or "Git Commit Hash" into their apps.

// This variable will be overwritten if we build with:
// go build -ldflags="-X 'main.AppVersion=v1.2.0'"
var AppVersion = "dev"

func demonstrateLdflags() {
	fmt.Println("\n--- Linker Flags ---")
	fmt.Printf("Current App Version: %s\n", AppVersion)
	fmt.Println("(Try building this with: go build -ldflags=\"-X 'main.AppVersion=v2.0.0'\")")
}
