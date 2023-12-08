package main

import (
	"compare-tool/read"
	"testing"
)

func TestCompare(t *testing.T) {
	h := read.ReadHandler{}

	h.ReadLargeByChunk("target.txt")
	srcs := make([]string, len(h.Contents))
	copy(srcs, h.Contents)

	h.ReadLargeByChunk("src.txt")
	dests := make(map[string]int, len(h.Contents))
	for _, v := range h.Contents {
		dests[v] = 0
	}

	findDiff(srcs, dests)

	if len(diff) != 4 {
		t.Errorf("compare two files failed, expected %d, got %d", 4, len(diff))
	}
}
