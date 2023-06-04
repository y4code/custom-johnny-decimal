package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	fmt.Println("Starting renaming...")

	// 1. Extract the first 2 digits of the parent folder name
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("Current working directory: " + getwd)
	base := filepath.Base(getwd)
	fmt.Println("Current parent folder name: " + base)
	category := base[:2]

	// 2. Sort files in the current folder by modified time
	files, err := os.ReadDir(getwd)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	sort.Slice(files, func(i, j int) bool {
		fileInfoI, err := files[i].Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			return false
		}
		fileInfoJ, err := files[j].Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			return false
		}
		return fileInfoI.ModTime().Before(fileInfoJ.ModTime())
	})

	for i, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name()[0] == '.' {
			continue
		}
		// 3. Rename the files with the first 2 digits of the parent folder name and a 2-digit incremental number
		newName := category + fmt.Sprintf(".%02d ", i+1) + file.Name()
		fmt.Println()
		fmt.Println(file.Name())
		fmt.Println(">>>>>>>>>>>>")
		fmt.Println(newName)
		fmt.Println()
		err = os.Rename(file.Name(), newName)
		if err != nil {
			fmt.Println("Error renaming file:", err)
			return
		}
	}
}
