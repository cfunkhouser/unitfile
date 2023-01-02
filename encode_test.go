package unitfile

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshal(t *testing.T) {
	cmpOpts := cmp.Options{cmp.AllowUnexported(testUnit{}, testUnitSection{})}
	for tn, tc := range map[string]struct {
		from any
		opts []MarshalOption
		want string
		err  error
	}{
		"invalid empty":       {err: ErrInvalidSource},
		"invalid map":         {from: make(map[string]any), err: ErrInvalidSource},
		"invalid nil pointer": {from: ptr[*struct{ Field string }](t, nil), err: ErrInvalidSource},
		"invalid float":       {from: float64(0), err: ErrInvalidSource},
		"invalid int":         {from: int(0), err: ErrInvalidSource},
		"invalid string":      {from: "", err: ErrInvalidSource},
		"zero bool": {
			from: &struct{ Field bool }{},
			// Booleans are always printed, for now.
			want: `[Unit]
Field=false` + "\n",
		},
		"zero float32": {
			from: &struct{ Field float32 }{},
			want: "[Unit]\n",
		},
		"zero float64": {
			from: &struct{ Field float64 }{},
			want: "[Unit]\n",
		},
		"zero int": {
			from: &struct{ Field int }{},
			want: "[Unit]\n",
		},
		"zero int64": {
			from: &struct{ Field int64 }{},
			want: "[Unit]\n",
		},
		"zero int32": {
			from: &struct{ Field int32 }{},
			want: "[Unit]\n",
		},
		"zero int16": {
			from: &struct{ Field int16 }{},
			want: "[Unit]\n",
		},
		"zero int8": {
			from: &struct{ Field int8 }{},
			want: "[Unit]\n",
		},
		"zero string": {
			from: &struct{ Field string }{},
			want: "[Unit]\n",
		},
		"zero []string": {
			from: &struct{ Field []string }{},
			want: "[Unit]\n",
		},
		"simple bool": {
			from: &struct{ One, Yes, True, On, Zero, No, False, Off bool }{true, true, true, true, false, false, false, false},
			want: `[Unit]
False=false
No=false
Off=false
On=true
One=true
True=true
Yes=true
Zero=false` + "\n",
		},
		"simple float32": {
			from: &struct{ Field float32 }{13.37},
			want: `[Unit]
Field=13.37` + "\n",
		},
		"simple float64": {
			from: &struct{ Field float64 }{13.37},
			want: `[Unit]
Field=13.37` + "\n",
		},
		"simple int": {
			from: &struct{ Field int }{42},
			want: `[Unit]
Field=42` + "\n",
		},
		"simple int64": {
			from: &struct{ Field int64 }{42},
			want: `[Unit]
Field=42` + "\n",
		},
		"simple int32": {
			from: &struct{ Field int32 }{42},
			want: `[Unit]
Field=42` + "\n",
		},
		"simple int16": {
			from: &struct{ Field int16 }{42},
			want: `[Unit]
Field=42` + "\n",
		},
		"simple int8": {
			from: &struct{ Field int8 }{42},
			want: `[Unit]
Field=42` + "\n",
		},
		"simple string": {
			from: &struct{ Field string }{"I am become death"},
			want: `[Unit]
Field=I am become death` + "\n",
		},
		"simple []string": {
			from: &struct{ Field []string }{[]string{"I am become death", "destroyer of worlds"}},
			want: `[Unit]
Field=I am become death
Field=destroyer of worlds` + "\n",
		},
		"nil pointer bool": {
			from: &struct{ Field *bool }{},
			want: `[Unit]` + "\n",
		},
		"nil pointer float32": {
			from: &struct{ Field *float32 }{},
			want: "[Unit]\n",
		},
		"nil pointer float64": {
			from: &struct{ Field *float64 }{},
			want: "[Unit]\n",
		},
		"nil pointer int": {
			from: &struct{ Field *int }{},
			want: "[Unit]\n",
		},
		"nil pointer int64": {
			from: &struct{ Field *int64 }{},
			want: "[Unit]\n",
		},
		"nil pointer int32": {
			from: &struct{ Field *int32 }{},
			want: "[Unit]\n",
		},
		"nil pointer int16": {
			from: &struct{ Field *int16 }{},
			want: "[Unit]\n",
		},
		"nil pointer int8": {
			from: &struct{ Field *int8 }{},
			want: "[Unit]\n",
		},
		"nil pointer string": {
			from: &struct{ Field *string }{},
			want: "[Unit]\n",
		},
		"nil pointer []string": {
			from: &struct{ Field *[]string }{},
			want: "[Unit]\n",
		},
		"pointer bool": {
			from: &struct{ One, Yes, True, On, Zero, No, False, Off *bool }{ptr(t, true), ptr(t, true), ptr(t, true), ptr(t, true), ptr(t, false), ptr(t, false), ptr(t, false), ptr(t, false)},
			want: `[Unit]
False=false
No=false
Off=false
On=true
One=true
True=true
Yes=true
Zero=false` + "\n",
		},
		"pointer float32": {
			from: &struct{ Field *float32 }{ptr(t, float32(13.37))},
			want: `[Unit]
Field=13.37` + "\n",
		},
		"pointer float64": {
			from: &struct{ Field *float64 }{ptr(t, 13.37)},
			want: `[Unit]
Field=13.37` + "\n",
		},
		"pointer int": {
			from: &struct{ Field *int }{ptr(t, 42)},
			want: `[Unit]
Field=42` + "\n",
		},
		"pointer int64": {
			from: &struct{ Field *int64 }{ptr(t, int64(42))},
			want: `[Unit]
Field=42` + "\n",
		},
		"pointer int32": {
			from: &struct{ Field *int32 }{ptr(t, int32(42))},
			want: `[Unit]
Field=42` + "\n",
		},
		"pointer int16": {
			from: &struct{ Field *int16 }{ptr(t, int16(42))},
			want: `[Unit]
Field=42` + "\n",
		},
		"pointer int8": {
			from: &struct{ Field *int8 }{ptr(t, int8(42))},
			want: `[Unit]
Field=42` + "\n",
		},
		"pointer string": {
			from: &struct{ Field *string }{ptr(t, "destroyer of worlds")},
			want: `[Unit]
Field=destroyer of worlds` + "\n",
		},
		"pointer []string": {
			from: &struct{ Field *[]string }{ptr(t, []string{"I am become death", "destroyer of worlds"})},
			want: `[Unit]
Field=I am become death
Field=destroyer of worlds` + "\n",
		},
		"documentation example 1a": {
			from: &struct{ KeyOne, KeyTwo string }{"value 1", "value 2"},
			want: `[Unit]
KeyOne=value 1
KeyTwo=value 2` + "\n",
		},
		"documentation example 1b": {
			from: &struct {
				Setting []string
				KeyTwo  string
			}{
				[]string{"something", "some thing", "…"},
				"value 2 \n\tvalue 2 continued",
			},
			want: `[Unit]
KeyTwo=value 2 \
	value 2 continued
Setting=some thing
Setting=something
Setting=…` + "\n",
		},
		"documentation example 1c": {
			from: &struct{ KeyThree string }{"value 3\n\t\t\t\tvalue 3 continued"},
			want: `[Unit]
KeyThree=value 3\
				value 3 continued` + "\n",
		},
	} {
		t.Run(tn, func(t *testing.T) {
			got, err := Marshal(tc.from, tc.opts...)
			if !errors.Is(err, tc.err) {
				t.Errorf("Marshal(): error mismatch, got: %v, want: %v", err, tc.err)
			}
			if diff := cmp.Diff(string(got), tc.want, cmpOpts...); diff != "" && tc.err == nil {
				t.Errorf("Unmarshal(): mismatch (-got,+want):\n%v", diff)
			}
		})
	}
}
