package search

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	Kilobytes int = 1024
	Megabytes     = 1024 * Kilobytes
	Gigabytes     = 1024 * Megabytes
)

// reads entire file into memory and uses strings.Contains to
// search for text
func SearchTextInFile(fname, text string) (bool, error) {
	log.Printf("searching for text '%v' in file %v\n", text, fname)

	data, err := os.ReadFile(fname)
	if err != nil {
		return false, err
	}

	found := strings.Contains(string(data), text)
	return found, nil
}

// reads entire file in memory, divvies up the data into same-sized
// chunks, spawns a goroutine to search for text within each chunk
func SearchTextInFile2(fname, text string) (bool, error) {
	log.Printf("searching for text '%v' in file %v\n", text, fname)

	data, err := os.ReadFile(fname)
	if err != nil {
		return false, err
	}

	buf_size := 10 * Megabytes
	chunks, err := chunkData(data, buf_size)
	if err != nil {
		return false, err
	}

	// spawn a goroutine to search for text within each chunk
	wg := sync.WaitGroup{}
	result_chan := make(chan bool, 100)

	for _, chunk := range chunks {
		wg.Add(1)

		go func(chunk []byte, text string) {
			defer wg.Done()

			found := strings.Contains(string(chunk), text)
			result_chan <- found
		}(chunk, text)
	}

	go func() {
		wg.Wait()
		close(result_chan)
	}()

	for found := range result_chan {
		if found {
			return true, nil
		}

		// continue searching
	}

	return false, nil
}

// chunks data into array of buffers each of size buf_size
func chunkData(data []byte, buf_size int) ([][]byte, error) {
	if buf_size <= 0 {
		return nil, fmt.Errorf("invalid buffer size")
	}

	chunks := [][]byte{}
	bytes_read := 0
	start := 0
	end := buf_size

	for bytes_read < len(data) {
		// we're on the last buff
		if end > len(data) {
			end = len(data)
		}

		buff := data[start:end]
		chunks = append(chunks, buff)

		start = end
		end += buf_size
		bytes_read += (end - start)
	}
	return chunks, nil
}
