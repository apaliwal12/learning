package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// ─────── THE io/fs PACKAGE (Go 1.16+) ───────
// Historically, Go tied file system operations directly to the OS (`os` package).
// In Go 1.16, the `io/fs` package introduced a standard `fs.FS` interface.
//
// This is incredibly powerful! A function that accepts an `fs.FS` can now read
// from the real disk, an embedded ZIP file (`archive/zip`), a mock file system in tests,
// or files embedded directly into the Go binary (`go:embed`).

// PrintDirTree takes ANY file system and prints its tree structure.
func PrintDirTree(fileSystem fs.FS, root string) error {
	fmt.Printf("Walking directory: %s\n", root)
	
	// fs.WalkDir is safer and faster than the old filepath.Walk
	err := fs.WalkDir(fileSystem, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		info, _ := d.Info()
		size := info.Size()
		
		if d.IsDir() {
			fmt.Printf("[DIR]  %s\n", path)
		} else {
			fmt.Printf("[FILE] %s (%d bytes)\n", path, size)
		}
		return nil
	})
	
	return err
}

func main() {
	fmt.Println("--- io/fs Project Demo ---")
	
	// We will create a temporary directory on the real disk for this demo
	tmpDir, err := os.MkdirTemp("", "fs_demo_*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir) // Clean up when done

	// Create some files
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("hello"), 0644)
	os.Mkdir(filepath.Join(tmpDir, "subdir"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "subdir", "file2.json"), []byte("{}"), 0644)

	// Now we use os.DirFS to create an fs.FS implementation anchored at tmpDir.
	// This acts like a chroot jail — the code can't look above this directory.
	myFS := os.DirFS(tmpDir)
	
	// Because PrintDirTree takes an interface, we can pass our real disk FS.
	// If we were testing, we could pass a memory-backed FS here!
	err = PrintDirTree(myFS, ".")
	if err != nil {
		fmt.Println("Error walking tree:", err)
	}
}
