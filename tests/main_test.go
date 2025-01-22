package tests

import (
	"elastic_search/search"
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	project_root string // project root directory
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	parent := filepath.Dir(file)
	project_root = filepath.Dir(parent)
}

func BenchmarkSearchTextInFile(b *testing.B) {
	file := filepath.Join(project_root, "./data/file.txt")
	search_text := "hello"

	b.StartTimer()
	found, err := search.SearchTextInFile(file, search_text)
	if err != nil {
		b.Errorf("unexpected error; %v", err)
	}
	elapsed := b.Elapsed()

	if found {
		log.Printf("text '%v' found in file %v\n", search_text, file)
	}

	b.Logf("done in %v\n", elapsed)
}

func BenchmarkSearchTextInFile2(b *testing.B) {
	file := filepath.Join(project_root, "./data/file.txt")
	search_text := "hello"

	b.StartTimer()
	found, err := search.SearchTextInFile2(file, search_text)
	if err != nil {
		b.Errorf("unexpected error; %v", err)
	}
	elapsed := b.Elapsed()

	if found {
		log.Printf("text '%v' found in file %v\n", search_text, file)
	}

	b.Logf("done in %v\n", elapsed)
}

func BenchmarkSearchTextInFile3(b *testing.B) {
	file := filepath.Join(project_root, "./data/file.txt")
	search_text := "hello"

	b.StartTimer()
	found, err := search.SearchTextInFile3(file, search_text)
	if err != nil {
		b.Errorf("unexpected error; %v", err)
	}
	elapsed := b.Elapsed()

	if found {
		log.Printf("text '%v' found in file %v\n", search_text, file)
	}

	b.Logf("done in %v\n", elapsed)
}
