// package wordstat is computing statistics for input Triplet
package wordstat

import "sort"

type Triplet [3]string

// Sink is computing statistics for input Triplet
// not thread safe
type Sink map[Triplet]int

type Stats struct {
	Triplet   Triplet
	Occurence int
}

// New returns a new Sink, ready to add compute Triplets stats
func New() Sink {
	return make(Sink)
}

// Inc is incrementing one for the triplet t
func (s Sink) Inc(t Triplet) {
	s[t]++
}

// Compute returns the top 100 triplets
func (s Sink) Compute() []Stats {
	var stats []Stats

	for t, count := range s {
		stats = append(stats, Stats{t, count})
	}
	sort.SliceStable(stats, func(i, j int) bool {
		return stats[i].Occurence > stats[j].Occurence
	})
	if len(stats) > 100 {
		return stats[:100]
	}
	return stats
}
