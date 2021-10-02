// package wordstat is computing statistics for input Triplet
package wordstat

import (
	"fmt"
	"sort"
)

// Triplet 3 strings
type Triplet [3]string

// Sink is computing statistics for input Triplet
// not thread safe
type Sink map[Triplet]int

// Stat is a triplet and its occurence
type Stat struct {
	Triplet   Triplet
	Occurence int
}

// Stats is a slice of Stat
type Stats []Stat

// New returns a new Sink, ready to add compute Triplets stats
func New() Sink {
	return make(Sink)
}

// Inc is incrementing one for the triplet t
func (s Sink) Inc(t Triplet) {
	s[t]++
}

// Compute returns the top 100 triplets
func (s Sink) Compute() Stats {
	var stats []Stat

	for t, count := range s {
		stats = append(stats, Stat{t, count})
	}
	sort.SliceStable(stats, func(i, j int) bool {
		return stats[i].Occurence > stats[j].Occurence
	})
	if len(stats) > 100 {
		return stats[:100]
	}
	return stats
}

func (s Stats) String() string {
	var str string
	for _, stat := range s {
		str += fmt.Sprintf("%s - %d\n", stat.Triplet, stat.Occurence)
	}
	return str
}

func (t Triplet) String() string {
	return t[0] + " " + t[1] + " " + t[2]
}
