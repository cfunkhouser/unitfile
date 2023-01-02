package unitfile

import (
	"embed"
	"testing"
)

//go:embed testdata
var testdataFS embed.FS

type testUnitSection struct {
	StringField string
	IntField    int
	FloatField  float64
	BoolField   bool
}

type testUnit struct {
	StringField      string
	IntField         int
	FloatField       float64
	BoolField        bool
	StringPtrField   *string
	IntPtrField      *int
	FloatPtrField    *float64
	BoolPtrField     *bool
	StructSection    testUnitSection
	StructPtrSection *testUnitSection
}

func readTestFile(tb testing.TB, path string) []byte {
	tb.Helper()
	data, err := testdataFS.ReadFile(path)
	if err != nil {
		tb.Fatalf("failed to read file %q: %v", path, err)
	}
	return data
}

func ptr[T any](tb testing.TB, v T) *T {
	tb.Helper()
	return &v
}
