// powered by Copilot
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// seed rand with current time and pid
	rand.Seed(int64(os.Getpid() ^ time.Now().Nanosecond()))

	// read command line arguments
	args := os.Args
	if len(args) != 3 {
		printUsageAndExit()
	}

	// get number of bytes to append
	numBytes, err := strconv.Atoi(args[2])
	if err != nil {
		printUsageAndExit()
	}

	// check if file or folder
	path := args[1]
	fileInfo, err := os.Stat(path)
	if err != nil {
		printUsageAndExit()
	}
	if fileInfo.IsDir() {
		folderPath := path
		appendRandomBytesToFilesInFolderRecursive(folderPath, numBytes)
	} else if fileInfo.Mode().IsRegular() {
		appendRandomBytesToFile(path, numBytes)
	} else {
		printUsageAndExit()
	}
}

func printUsageAndExit() {
	fmt.Println("Usage: appendRandomBytesToFiles <filename/folder> <number of bytes>")
	os.Exit(1)
}

func appendRandomBytesToFilesInFolderRecursive(folder string, numBytes int) error {
	// open folder
	dir, err := os.Open(folder)
	if err != nil {
		return err
	}
	defer dir.Close()

	// read folder entries
	entries, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	// iterate over entries
	for _, entry := range entries {
		if entry.IsDir() {
			// traverse subfolders
			err := appendRandomBytesToFilesInFolderRecursive(folder+"/"+entry.Name(), numBytes)
			if err != nil {
				return err
			}
		} else if entry.Mode().IsRegular() {
			// append random bytes to file
			err := appendRandomBytesToFile(folder+"/"+entry.Name(), numBytes)
			if err != nil {
				return err
			}
		} else {
			// ignore other entries
		}
	}

	return nil
}

func appendRandomBytesToFile(filename string, numBytes int) error {
	// open file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// make random bytes
	buffer := make([]byte, numBytes)
	for i := 0; i < numBytes; i++ {
		buffer[i] = byte(rand.Intn(256))
	}

	// write random bytes to file
	_, err = file.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}
