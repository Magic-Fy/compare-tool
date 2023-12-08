package read

import (
	"bufio"
	"github.com/panjf2000/ants/v2"
	"os"
	"strings"
	"sync"
)

type ReadHandler struct {
	Contents []string
}

var MAXGOROUTINE = 10000

// ReadLargeByChunk Read a large file by read a chunk of 16MB
func (handler *ReadHandler) ReadLargeByChunk(filename string) error {
	handler.Contents = make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// set the buffer size to 16MB
	buf := make([]byte, 16*1024*1024)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	// create a wait group to ensure all goroutines finish before returning
	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(MAXGOROUTINE, func(line interface{}) {
		handler.Contents = append(handler.Contents, line.(string))
		wg.Done()
	})

	defer p.Release()

	// process each line in parallel
	for scanner.Scan() {
		wg.Add(1)
		p.Invoke(strings.Trim(scanner.Text(), "\n"))
	}

	// wait for all goroutines to finish
	wg.Wait()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
