package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func walkDir() {
	var files []string
	dir_path := "."
	err := filepath.WalkDir(dir_path, func(path string, info fs.DirEntry, err error) error {
		dir_name := filepath.Base(dir_path)
		if info.IsDir() == true && info.Name() != dir_name {
			return filepath.SkipDir
		} else {
			files = append(files, path)
			return nil
		}
	})

	if err != nil {
		panic(err)
	}
	for _, file := range files {
		log.Println(file)
	}
}

func walkDirRecursive() {
	var files []string
	root := "static/"
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		files = append(files, path)
		var f string
		if info.IsDir() {
			f = "Directory"
		} else {
			f = "File"
		}
		log.Printf("%s Name: %s\n", f, info.Name())
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		log.Println(file)
	}
}

func onlyDir() {
	var folders []string
	root := "static/"
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		dir_name := filepath.Base(root)
		if info.IsDir() {
			folders = append(folders, info.Name())
			return nil
		} else if info.IsDir() && dir_name != info.Name() {
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, folder := range folders {
		log.Println(folder)
	}
}

func walkAsAnonymousFunction() {
	var files []string
	var dirs []string

	root := "math/"
	// walk function with parameter as os.FileInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		dir_name := filepath.Base(root)

		// skip the path if it is a directory and also the nested directory
		if info.IsDir() == true && info.Name() != dir_name {
			return filepath.SkipDir
		} else {
			files = append(files, path)
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
	log.Println("Files : ")
	for _, file := range files {
		log.Println(file)
	}

	err = filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		dir_name := filepath.Base(root)

		// add the path if the path is a directory
		if info.IsDir() == true && info.Name() != dir_name {
			dirs = append(dirs, path)
		} else {
			// add the path to the files slice
			files = append(files, path)
		}
		if err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
	log.Println("Files : ")
	for _, file := range files {
		log.Println(file)
	}
	log.Println("Dirs: ")
	for _, dir := range dirs {
		log.Println(dir)
	}
}
