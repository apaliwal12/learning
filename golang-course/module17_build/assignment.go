package main

// TODO 1: Implement BuildCommand
// Write a function that returns the EXACT string of the bash command required
// to compile `main.go` into an executable named `server`, targeting a 64-bit Linux OS.
// Hint: You need to set two environment variables before the `go build` command.
func LinuxBuildCommand() string {
	return "" // Fix me
}

// TODO 2: Implement LDFlagsCommand
// Write a function that returns the EXACT string of the bash command required
// to compile `main.go`, injecting the value "v1.5.0" into a variable named "Version"
// in the "main" package.
// Hint: use the -ldflags flag.
func LDFlagsCommand() string {
	return "" // Fix me
}
