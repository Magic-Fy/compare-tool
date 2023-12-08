// read_test.go
package read

import "testing"

func TestRead(t *testing.T) {
	handler := ReadHandler{}
	handler.ReadLargeByChunk("../src.txt")

	if len(handler.Contents) != 3 {
		t.Errorf("readLarge() failed, expected %d, got %d", 5000000, len(handler.Contents))
	}
}

func BenchmarkRead(b *testing.B) {
	handler := ReadHandler{}
	for i := 0; i < b.N; i++ {
		handler.ReadLargeByChunk("../src.txt")
	}
}
