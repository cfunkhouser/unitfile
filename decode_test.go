package unitfile

import (
	"embed"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestCleanValue(t *testing.T) {
	for tn, tc := range map[string]struct {
		val  string
		want string
	}{
		"empty":                       {},
		"unix endline":                {"A\nB", "A\nB"},
		"windows endline":             {"A\r\nB", "A\r\nB"},
		"unix contline":               {"A\\\nB", "A\nB"},
		"windows contline":            {"A\\\r\nB", "A\nB"},
		"unix multi-contline":         {"A\\\nB\\\nC", "A\nB\nC"},
		"windows multi-contline":      {"A\\\r\nB\\\r\nC", "A\nB\nC"},
		"non-contline escapes":        {"A\\nB", "A\\nB"},
		"unix contline whitespace":    {" \n\tA\\\nB\t\n ", "A\nB"},
		"windows contline whitespace": {" \n\tA\\\r\nB\t\n ", "A\nB"},
	} {
		t.Run(tn, func(t *testing.T) {
			got := cleanValue(tc.val)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("cleanValue(): mismatch (-got,+want):\n%v", diff)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	cmpOpts := cmp.Options{cmp.AllowUnexported(testUnit{}, testUnitSection{})}
	for tn, tc := range map[string]struct {
		data []byte
		opts []Option
		got  any
		want any
		err  error
	}{
		"invalid empty":       {err: ErrInvalidDestination},
		"invalid map":         {got: make(map[string]any), err: ErrInvalidDestination},
		"invalid nil pointer": {got: ptr[*struct{ Field string }](t, nil), err: ErrInvalidDestination},
		"invalid float":       {got: float64(0), err: ErrInvalidDestination},
		"invalid int":         {got: int(0), err: ErrInvalidDestination},
		"invalid string":      {got: "", err: ErrInvalidDestination},
		"zero bool": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field bool }{},
			want: &struct{ Field bool }{},
		},
		"zero float32": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field float32 }{},
			want: &struct{ Field float32 }{},
		},
		"zero float64": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field float64 }{},
			want: &struct{ Field float64 }{},
		},
		"zero int": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field int }{},
			want: &struct{ Field int }{},
		},
		"zero int64": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field int64 }{},
			want: &struct{ Field int64 }{},
		},
		"zero int32": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field int32 }{},
			want: &struct{ Field int32 }{},
		},
		"zero int16": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field int16 }{},
			want: &struct{ Field int16 }{},
		},
		"zero int8": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field int8 }{},
			want: &struct{ Field int8 }{},
		},
		"zero string": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field string }{},
			want: &struct{ Field string }{},
		},
		"zero []string": {
			data: []byte(`[Unit]
Field=`),
			got:  &struct{ Field []string }{},
			want: &struct{ Field []string }{},
		},
		"simple bool": {
			data: []byte(`[Unit]
One=1
Yes=yes
True=true
On=on
Zero=0
No=no
False=false
Off=off`),
			got:  &struct{ One, Yes, True, On, Zero, No, False, Off bool }{},
			want: &struct{ One, Yes, True, On, Zero, No, False, Off bool }{true, true, true, true, false, false, false, false},
		},
		"simple float32": {
			data: []byte(`[Unit]
Field=13.37`),
			got:  &struct{ Field float32 }{},
			want: &struct{ Field float32 }{13.37},
		},
		"simple float64": {
			data: []byte(`[Unit]
Field=13.37`),
			got:  &struct{ Field float64 }{},
			want: &struct{ Field float64 }{13.37},
		},
		"simple int": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field int }{},
			want: &struct{ Field int }{42},
		},
		"simple int64": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field int64 }{},
			want: &struct{ Field int64 }{42},
		},
		"simple int32": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field int32 }{},
			want: &struct{ Field int32 }{42},
		},
		"simple int16": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field int16 }{},
			want: &struct{ Field int16 }{42},
		},
		"simple int8": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field int8 }{},
			want: &struct{ Field int8 }{42},
		},
		"simple string": {
			data: []byte(`[Unit]
Field=this is a string`),
			got:  &struct{ Field string }{},
			want: &struct{ Field string }{"this is a string"},
		},
		"simple slice repeated field": {
			data: []byte(`[Unit]
Field=first string
Field=second string`),
			got:  &struct{ Field []string }{},
			want: &struct{ Field []string }{[]string{"first string", "second string"}},
		},
		"simple slice repeated value": {
			data: []byte(`[Unit]
Field="first string" "second string"`),
			got:  &struct{ Field []string }{},
			want: &struct{ Field []string }{[]string{"first string", "second string"}},
		},
		"pointer bool": {
			data: []byte(`[Unit]
One=1
Yes=yes
True=true
On=on
Zero=0
No=no
False=false
Off=off`),
			got:  &struct{ One, Yes, True, On, Zero, No, False, Off *bool }{},
			want: &struct{ One, Yes, True, On, Zero, No, False, Off *bool }{ptr(t, true), ptr(t, true), ptr(t, true), ptr(t, true), ptr(t, false), ptr(t, false), ptr(t, false), ptr(t, false)},
		},
		"pointer float32": {
			data: []byte(`[Unit]
Field=13.37`),
			got:  &struct{ Field *float32 }{},
			want: &struct{ Field *float32 }{ptr(t, float32(13.37))},
		},
		"pointer float64": {
			data: []byte(`[Unit]
Field=13.37`),
			got:  &struct{ Field *float64 }{},
			want: &struct{ Field *float64 }{ptr(t, 13.37)},
		},
		"pointer int": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field *int }{},
			want: &struct{ Field *int }{ptr(t, 42)},
		},
		"pointer int64": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field *int64 }{},
			want: &struct{ Field *int64 }{ptr(t, int64(42))},
		},
		"pointer int32": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field *int32 }{},
			want: &struct{ Field *int32 }{ptr(t, int32(42))},
		},
		"pointer int16": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field *int16 }{},
			want: &struct{ Field *int16 }{ptr(t, int16(42))},
		},
		"pointer int8": {
			data: []byte(`[Unit]
Field=42`),
			got:  &struct{ Field *int8 }{},
			want: &struct{ Field *int8 }{ptr(t, int8(42))},
		},
		"pointer string": {
			data: []byte(`[Unit]
Field=this is a string`),
			got:  &struct{ Field *string }{},
			want: &struct{ Field *string }{ptr(t, "this is a string")},
		},
		"pointer slice repeated field": {
			data: []byte(`[Unit]
Field=first string
Field=second string`),
			got:  &struct{ Field *[]string }{},
			want: &struct{ Field *[]string }{ptr(t, []string{"first string", "second string"})},
		},
		"pointer slice repeated value": {
			data: []byte(`[Unit]
Field="first string" "second string"`),
			got:  &struct{ Field *[]string }{},
			want: &struct{ Field *[]string }{ptr(t, []string{"first string", "second string"})},
		},
		"pure tags": {
			data: []byte(`[Unit]
Something=this is a string
Else=42`),
			got: &struct {
				FieldA string `unit:"Something"`
				FieldB int    `unit:"Else"`
			}{},
			want: &struct {
				FieldA string `unit:"Something"`
				FieldB int    `unit:"Else"`
			}{"this is a string", 42},
		},
		"mixed tags": {
			data: []byte(`[Unit]
Something=this is a string
FieldB=42`),
			got: &struct {
				FieldA string `unit:"Something"`
				FieldB int
			}{},
			want: &struct {
				FieldA string `unit:"Something"`
				FieldB int
			}{"this is a string", 42},
		},
		"field not in struct": {
			data: []byte(`[Unit]
Exists=this is a string
DoesNotExist=42`),
			got:  &struct{ Exists string }{},
			want: &struct{ Exists string }{"this is a string"},
		},
		"section not in struct": {
			data: []byte(`[Unit]
Exists=this is a string
[SectionA]
DoesNotExist=42`),
			got:  &struct{ Exists string }{},
			want: &struct{ Exists string }{"this is a string"},
		},
		"strict section not in struct": {
			data: []byte(`[Unit]
Exists=this is a string
[SectionA]
DoesNotExist=42`),
			opts: []Option{Strictly},
			got:  &struct{ Exists string }{},
			err:  ErrInvalidDestination,
		},
		// Examples from https://www.freedesktop.org/software/systemd/man/systemd.syntax.html#id-1.4.6
		"documentation example 1a": {
			data: []byte(`[Section A]
KeyOne=value 1
KeyTwo=value 2`),
			got:  &struct{ KeyOne, KeyTwo string }{},
			want: &struct{ KeyOne, KeyTwo string }{"value 1", "value 2"},
		},
		"documentation example 1b": {
			data: []byte(`[Section B]
Setting="something" "some thing" "…"
KeyTwo=value 2 \
	value 2 continued`),
			got: &struct {
				Setting []string
				KeyTwo  string
			}{},
			want: &struct {
				Setting []string
				KeyTwo  string
			}{
				[]string{"something", "some thing", "…"},
				"value 2 \n\tvalue 2 continued"},
		},
		"documentation example 1c": {
			data: []byte(`[Section C]
KeyThree=value 3\
# this line is ignored
; this line is ignored too
				value 3 continued`),
			got:  &struct{ KeyThree string }{},
			want: &struct{ KeyThree string }{"value 3\n\t\t\t\tvalue 3 continued"},
		},
		"full struct test": {
			data: readTestFile(t, "testdata/test-full.unit"),
			got:  &testUnit{},
			want: &testUnit{
				StringField:    "This is a string",
				IntField:       42,
				FloatField:     420.69,
				BoolField:      false,
				StringPtrField: ptr(t, "This is a string pointer inside of a section."),
				IntPtrField:    ptr(t, 69),
				FloatPtrField:  ptr(t, 13.37),
				BoolPtrField:   ptr(t, true),
				StructSection: testUnitSection{
					StringField: "This is a string inside of a section.",
					IntField:    69,
					FloatField:  13.37,
					BoolField:   true,
				},
				StructPtrSection: ptr(t, testUnitSection{
					StringField: "This is a string inside of a pointer section.",
					IntField:    69,
					FloatField:  13.37,
					BoolField:   true,
				}),
			},
		},
		"real - ubuntu 22.04.1 sshd.service": {
			data: readTestFile(t, "testdata/ubuntu-22.04.1-ssh.service"),
			got:  &Service{},
			want: &Service{
				Unit: Unit{
					Description:         "OpenBSD Secure Shell server",
					Documentation:       []string{"man:sshd(8) man:sshd_config(5)"},
					After:               []string{"network.target auditd.service"},
					ConditionPathExists: []string{"!/etc/ssh/sshd_not_to_be_run"},
				},
				Service: ServiceSection{
					Type:                     "notify",
					ExecStart:                "/usr/sbin/sshd -D $SSHD_OPTS",
					ExecStartPre:             "/usr/sbin/sshd -t",
					ExecReload:               "/bin/kill -HUP $MAINPID",
					Restart:                  "on-failure",
					RestartPreventExitStatus: 255,
				},
				Install: InstallSection{
					Alias:    []string{"sshd.service"},
					WantedBy: []string{"multi-user.target"},
				},
			},
		},
	} {
		t.Run(tn, func(t *testing.T) {
			err := Unmarshal(tc.data, tc.got, tc.opts...)
			if !errors.Is(err, tc.err) {
				t.Errorf("Unmarshal(): error mismatch, got: %v, want: %v", err, tc.err)
			}
			// The destination value is undefined if Unmarshal returns an error value,
			// so only test the gotten value when the expected error is nil.
			if diff := cmp.Diff(tc.got, tc.want, cmpOpts...); diff != "" && tc.err == nil {
				t.Errorf("Unmarshal(): mismatch (-got,+want):\n%v", diff)
			}
		})
	}
}
