package unitfile

import "testing"

func benchmarkCleanValue(b *testing.B, in string) {
	for n := 0; n < b.N; n++ {
		_ = cleanValue(in)
	}
}

func BenchmarkCleanValueContlineAndWhitespace(b *testing.B) {
	benchmarkCleanValue(b, `
		This is a test \
  with line \
	continuations and surrounding whitespace	  
	`)
}

func BenchmarkCleanValueWhitespaceOnly(b *testing.B) {
	benchmarkCleanValue(b, `
		This is a test with only surrounding whitespace	  
	`)
}

func BenchmarkCleanValueNoContlineOrWhitespace(b *testing.B) {
	benchmarkCleanValue(b, "false")
}

func BenchmarkCleanValueEmptyInput(b *testing.B) {
	benchmarkCleanValue(b, "")
}

func benchmarkUnmarshal(b *testing.B, in []byte, out any) {
	for n := 0; n < b.N; n++ {
		_ = Unmarshal([]byte(in), out)
	}
}

func BenchmarkUnmarshalMultipleSectionUnit(b *testing.B) {
	data := readTestFile(b, "testdata/test-full.unit")
	benchmarkUnmarshal(b, data, &testUnit{})
}
