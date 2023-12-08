package main

import (
	"compare-tool/read"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/panjf2000/ants/v2"
	"os"
	"sync"
)

// src is the source file
// des is the destination file
// top is the number of lines to show
// con is the number of find function goroutines
var (
	src  string
	des  string
	top  int
	diff = make([]string, 0)
	con  = 10000
)

// fileExists checks if a file exists and is not a directory
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func findFunc(s []string, d map[string]int) []string {
	results := make([]string, 0)
	for _, k := range s {
		if _, ok := d[k]; !ok {
			results = append(results, k)
		}
	}
	return results
}

// findDiff finds the difference between two files
func findDiff(srcs []string, dests map[string]int) {
	var wg sync.WaitGroup
	diff = make([]string, 0)
	p, _ := ants.NewPoolWithFunc(con, func(i interface{}) {
		results := findFunc(i.([]string), dests)
		diff = append(diff, results...)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < len(srcs); i += con {
		wg.Add(1)
		if i+con > len(srcs) {
			p.Invoke(srcs[i:])
		} else {
			p.Invoke(srcs[i : i+con])
		}
	}
	wg.Wait()
}

// compare compares two files
func compare(ip1 string, ip2 string, showTop int) {
	h := read.ReadHandler{}

	h.ReadLargeByChunk(ip1)
	srcs := make([]string, len(h.Contents))
	copy(srcs, h.Contents)
	fmt.Printf("%s has %d lines...\n", ip1, len(srcs))
	h.ReadLargeByChunk(ip2)
	dests := make(map[string]int, len(h.Contents))
	destlines := make([]string, len(h.Contents))
	copy(destlines, h.Contents)
	fmt.Printf("%s has %d lines...\n", ip2, len(destlines))
	for _, v := range h.Contents {
		dests[v] = 0
	}

	findDiff(srcs, dests)

	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("\nin %s but %s in %s: %s lines\n", ip1, red("not"), ip2, red(len(diff)))
	if len(diff) > 0 {
		if len(diff) < 10 {
			showTop = len(diff)
		}
		fmt.Printf("diff top %d lines:\n", showTop)
		for _, v := range diff[:showTop] {
			fmt.Printf("%s\n", v)
		}
	}
	fmt.Printf("\n\n")
}

func main() {
	flag.StringVar(&src, "src", "", "one of the input file")
	flag.StringVar(&des, "dest", "", "another input file")
	flag.IntVar(&top, "top N", 10, "show top N lines")

	flag.Parse()

	if !fileExists(src) || !fileExists(des) {
		fmt.Printf("src or dest file does not exist")
		return
	}

	compare(src, des, top)
	compare(des, src, top)
}
