package main

import (
	"fmt"
	"os"
)

// ─────── 2. WORKING WITH FILES (os package) ───────
// In Go, `os.File` implements both `io.Reader` and `io.Writer`.
// Always remember to close files to prevent resource leaks!

func demonstrateFiles() {
	fmt.Println("\n--- Files ---")
	
	filename := "test_demo.txt"
	
	// 1. Write to a file
	// os.WriteFile is a convenience function that creates/truncates the file,
	// writes the data, and closes it in one shot. (Replaces ioutil.WriteFile)
	data := []byte("Hello, File System!\nSecond line.")
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully using os.WriteFile.")

	// 2. Read from a file (Whole file)
	// os.ReadFile reads the entire file into memory. Good for small files.
	readData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("File contents:\n%s\n", readData)

	// 3. Appending to a file
	// We need to use os.OpenFile to specify exactly how we want to open it.
	// os.O_APPEND: append data to the file when writing.
	// os.O_WRONLY: open the file write-only.
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for append:", err)
		return
	}
	
	// DEFER THE CLOSE! Crucial for preventing leaks.
	defer f.Close()

	if _, err := f.WriteString("\nAppended third line."); err != nil {
		fmt.Println("Error appending to file:", err)
	} else {
		fmt.Println("Appended to file successfully.")
	}

	// 4. Cleanup
	// We'll remove the file so we don't clutter your disk during the demo.
	// Normally you wouldn't do this immediately after writing.
	os.Remove(filename)
}
