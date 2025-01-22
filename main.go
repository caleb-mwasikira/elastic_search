package main

import (
	"elastic_search/search"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	search_text string
	files       []string = []string{}
	recursive   bool
)

func init() {
	var (
		help      bool
		files_arg []string = []string{}
		err       error
	)

	flag.StringVar(&search_text, "search", "", "Text to search for within provided files")
	flag.BoolVar(&recursive, "recursive", false, "If a directory is provided, the recursive option will search all files in that directory and its subdirectories for text.")
	flag.BoolVar(&help, "help", false, "Display usage")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if search_text == "" {
		log.Fatalln("no text to search for")
	}

	files_arg = flag.Args()
	files, err = extractFilesFromArgs(files_arg, false)
	if err != nil {
		log.Fatalf("error parsing argument filenames; %v\n", err)
	}

	if len(files) == 0 {
		log.Fatalln("zero files found")
	}
}

func extractFilesFromArgs(args []string, recursive bool) ([]string, error) {
	files := []string{}

	for _, arg := range args {
		stat, err := os.Stat(arg)
		if err != nil {
			return nil, fmt.Errorf("extractFilesFromArgs: invalid file/dir path %v given as argument", arg)
		}

		if stat.IsDir() {
			dirFiles, err := getFilesInDir(arg, recursive)
			if err != nil {
				return nil, err
			}

			files = append(files, dirFiles...)
			continue
		}

		if stat.Mode().IsRegular() {
			files = append(files, arg)
		}
	}
	return files, nil
}

func getFilesInDir(dir string, recursive bool) ([]string, error) {
	stat, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, fmt.Errorf("getFilesInDir: path is not a directory")
	}

	files := []string{}

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if dir == path {
			// current directory
			return nil
		}

		if d.IsDir() && !recursive {
			return fs.SkipDir
		}

		if d.Type().IsRegular() {
			files = append(files, path)
		}

		return nil
	})

	return files, nil
}

func main() {
	start := time.Now()
	// log.Printf("files %v\n", files)

	for _, file := range files {
		found, err := search.SearchTextInFile3(file, search_text)
		if err != nil {
			log.Fatalf("error searching for text in file %v; %v\n", file, err)
			os.Exit(1)
		}

		if found {
			log.Printf("text '%v' found in file %v\n", search_text, file)
		}
	}

	elapsed := time.Since(start)
	log.Printf("completed in %v\n", elapsed)
}
