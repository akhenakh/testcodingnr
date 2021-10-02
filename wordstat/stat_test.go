// package wordstat is computing statistics for input Triplet
package wordstat

import (
	"reflect"
	"testing"
)

func TestSink_Compute(t *testing.T) {
	tests := []struct {
		name  string
		input []Triplet
		want  Stats
	}{
		{"simple count", []Triplet{{"a", "b", "c"}}, []Stat{{Triplet{"a", "b", "c"}, 1}}},
		{"simple count", []Triplet{{"a", "b", "c"}, {"a", "b", "c"}}, []Stat{{Triplet{"a", "b", "c"}, 2}}},
		{"simple count", []Triplet{{"a", "b", "c"}, {"a", "b", "c"}, {"b", "c", "d"}}, []Stat{{Triplet{"a", "b", "c"}, 2}, {Triplet{"b", "c", "d"}, 1}}},
		{"simple count", []Triplet{{"a", "b", "c"}, {"b", "c", "d"}, {"a", "b", "c"}}, []Stat{{Triplet{"a", "b", "c"}, 2}, {Triplet{"b", "c", "d"}, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			for _, v := range tt.input {
				s.Inc(v)
			}
			if got := s.Compute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sink.Compute() = %v, want %v", got, tt.want)
			}
		})
	}
}
