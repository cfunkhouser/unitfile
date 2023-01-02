package unitfile

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranscodingIdempotency(t *testing.T) {
	cmpOpts := cmp.Options{cmp.AllowUnexported(testUnit{}, testUnitSection{})}
	for tn, tc := range map[string]struct {
		mopts []MarshalOption
		uopts []UnmarshalOption
		got   any
		want  any
	}{
		"zero bool": {
			got:  &struct{ Field bool }{},
			want: &struct{ Field bool }{},
		},
		"zero float32": {
			got:  &struct{ Field float32 }{},
			want: &struct{ Field float32 }{},
		},
		"zero float64": {
			got:  &struct{ Field float64 }{},
			want: &struct{ Field float64 }{},
		},
		"zero int": {
			got:  &struct{ Field int }{},
			want: &struct{ Field int }{},
		},
		"zero int64": {
			got:  &struct{ Field int64 }{},
			want: &struct{ Field int64 }{},
		},
		"zero int32": {
			got:  &struct{ Field int32 }{},
			want: &struct{ Field int32 }{},
		},
		"zero int16": {
			got:  &struct{ Field int16 }{},
			want: &struct{ Field int16 }{},
		},
		"zero int8": {
			got:  &struct{ Field int8 }{},
			want: &struct{ Field int8 }{},
		},
		"zero string": {
			got:  &struct{ Field string }{},
			want: &struct{ Field string }{},
		},
		"zero []string": {
			got:  &struct{ Field []string }{},
			want: &struct{ Field []string }{},
		},
		"simple bool": {
			got:  &struct{ One, Yes, True, On, Zero, No, False, Off bool }{},
			want: &struct{ One, Yes, True, On, Zero, No, False, Off bool }{true, true, true, true, false, false, false, false},
		},
		"simple float32": {
			got:  &struct{ Field float32 }{},
			want: &struct{ Field float32 }{13.37},
		},
		"simple float64": {
			got:  &struct{ Field float64 }{},
			want: &struct{ Field float64 }{13.37},
		},
		"simple int": {
			got:  &struct{ Field int }{},
			want: &struct{ Field int }{42},
		},
		"simple int64": {
			got:  &struct{ Field int64 }{},
			want: &struct{ Field int64 }{42},
		},
		"simple int32": {
			got:  &struct{ Field int32 }{},
			want: &struct{ Field int32 }{42},
		},
		"simple int16": {
			got:  &struct{ Field int16 }{},
			want: &struct{ Field int16 }{42},
		},
		"simple int8": {
			got:  &struct{ Field int8 }{},
			want: &struct{ Field int8 }{42},
		},
		"simple string": {
			got:  &struct{ Field string }{},
			want: &struct{ Field string }{"this is a string"},
		},
		"simple slice repeated field": {
			got:  &struct{ Field []string }{},
			want: &struct{ Field []string }{[]string{"first string", "second string"}},
		},
		"simple slice repeated value": {
			got:  &struct{ Field []string }{},
			want: &struct{ Field []string }{[]string{"first string", "second string"}},
		},
		"pointer bool": {
			got:  &struct{ One, Yes, True, On, Zero, No, False, Off *bool }{},
			want: &struct{ One, Yes, True, On, Zero, No, False, Off *bool }{ptr(t, true), ptr(t, true), ptr(t, true), ptr(t, true), ptr(t, false), ptr(t, false), ptr(t, false), ptr(t, false)},
		},
		"pointer float32": {
			got:  &struct{ Field *float32 }{},
			want: &struct{ Field *float32 }{ptr(t, float32(13.37))},
		},
		"pointer float64": {
			got:  &struct{ Field *float64 }{},
			want: &struct{ Field *float64 }{ptr(t, 13.37)},
		},
		"pointer int": {
			got:  &struct{ Field *int }{},
			want: &struct{ Field *int }{ptr(t, 42)},
		},
		"pointer int64": {
			got:  &struct{ Field *int64 }{},
			want: &struct{ Field *int64 }{ptr(t, int64(42))},
		},
		"pointer int32": {
			got:  &struct{ Field *int32 }{},
			want: &struct{ Field *int32 }{ptr(t, int32(42))},
		},
		"pointer int16": {
			got:  &struct{ Field *int16 }{},
			want: &struct{ Field *int16 }{ptr(t, int16(42))},
		},
		"pointer int8": {
			got:  &struct{ Field *int8 }{},
			want: &struct{ Field *int8 }{ptr(t, int8(42))},
		},
		"pointer string": {
			got:  &struct{ Field *string }{},
			want: &struct{ Field *string }{ptr(t, "this is a string")},
		},
		"pointer slice repeated field": {
			got:  &struct{ Field *[]string }{},
			want: &struct{ Field *[]string }{ptr(t, []string{"first string", "second string"})},
		},
		"pointer slice repeated value": {
			got:  &struct{ Field *[]string }{},
			want: &struct{ Field *[]string }{ptr(t, []string{"first string", "second string"})},
		},
		"pure tags": {
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
			got: &struct {
				FieldA string `unit:"Something"`
				FieldB int
			}{},
			want: &struct {
				FieldA string `unit:"Something"`
				FieldB int
			}{"this is a string", 42},
		},
		"documentation example 1a": {
			got:  &struct{ KeyOne, KeyTwo string }{},
			want: &struct{ KeyOne, KeyTwo string }{"value 1", "value 2"},
		},
		"documentation example 1b": {
			got: &struct {
				Setting []string
				KeyTwo  string
			}{},
			want: &struct {
				Setting []string
				KeyTwo  string
			}{
				[]string{"some thing", "something", "…"},
				"value 2 \n\tvalue 2 continued"},
		},
		"documentation example 1c": {
			got:  &struct{ KeyThree string }{},
			want: &struct{ KeyThree string }{"value 3\n\t\t\t\tvalue 3 continued"},
		},
		"full struct test": {
			got: &testUnit{},
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
			got: &Service{},
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
			data, err := Marshal(tc.want, tc.mopts...)
			if err != nil {
				t.Fatalf("Marshal(): unwanted error: %v", err)
			}
			if err := Unmarshal(data, tc.got, tc.uopts...); err != nil {
				t.Fatalf("error unmarshalling intermediate:\n=== BEGIN UNIT ===\n%v\n=== END UNIT ===\n%v", string(data), err)
			}
			if diff := cmp.Diff(tc.got, tc.want, cmpOpts...); diff != "" {
				t.Errorf("mismatch (-got,+want):\n%v", diff)
			}
		})
	}
}
