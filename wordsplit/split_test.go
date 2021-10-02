// package wordsplit is splitting input into 3 words triplet
package wordsplit

import (
	"context"
	"os"
	"reflect"
	"testing"
)

func TestParseAndSplit(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		splits []Triplet
	}{
		{"simple file", "../testdata/test00.txt", []Triplet{
			{"one", "two", "three"},
			{"two", "three", "four"},
			{"three", "four", "five"},
		}},
		{"simple file", "../testdata/test01.txt", []Triplet{
			{"one", "two", "three"},
			{"two", "three", "four"},
			{"three", "four", "two"},
			{"four", "two", "three"},
			{"two", "three", "four"},
			{"three", "four", "one"},
			{"four", "one", "two"},
		}},
		{"simple file punctuation", "../testdata/test02.txt", []Triplet{
			{"one", "two", "three"},
			{"two", "three", "four"},
		}},
		{"simple file punctuation and caps", "../testdata/test03.txt", []Triplet{
			{"one", "two", "three"},
			{"two", "three", "four"},
		}},
		{"simple file shouldn't split '", "../testdata/test04.txt", []Triplet{
			{"shouldn't", "be", "split"},
		}},
		{"simple file unicode", "../testdata/test05.txt", []Triplet{
			{"süsse", "in", "straße"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.path)
			if err != nil {
				t.Fatal(err)
			}
			itr := ParseAndSplit(context.Background(), f)
			var res []Triplet
			for {
				triplet, err := itr.Next()
				if err == Done {
					break
				}
				if err != nil {
					t.Fatal(err)
				}
				res = append(res, triplet)
			}
			if !reflect.DeepEqual(res, tt.splits) {
				t.Errorf("ParseAndSplit() = %v, want %v", res, tt.splits)
			}
		})
	}
}
