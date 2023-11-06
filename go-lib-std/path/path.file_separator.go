package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func JoinWithPathSeparator() {
	// Using os.PathSeparator
	path := fmt.Sprintf("folder%csubfolder%cfile.txt", os.PathSeparator, os.PathSeparator)
	fmt.Println(path)

	// An alternative way is to use the filepath.Join function, which is even better
	// because it takes care of the separators for you.
	joinedPath := filepath.Join("folder", "subfolder", "file.txt")
	fmt.Println(joinedPath)
}
