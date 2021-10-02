package tripletsvc

import (
	"bytes"
	context "context"
	"fmt"

	"github.com/akhenakh/codingtestnr/wordsplit"
	"github.com/akhenakh/codingtestnr/wordstat"
)

type Server struct{}

func (s *Server) Compute(ctx context.Context, req *ComputeRequest) (*ComputeResponse, error) {
	stat := wordstat.New()

	b := bytes.NewBuffer(req.Text)
	itr := wordsplit.ParseAndSplit(ctx, b)
	for {
		t, err := itr.Next()
		if err == wordsplit.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't parse text: %v", err)
		}
		stat.Inc(wordstat.Triplet(t))
	}

	stats := stat.Compute()
	gstats := make([]*Stat, len(stats))
	for i, stat := range stats {
		gstats[i] = &Stat{
			Triplet: &Triplet{
				Words: []string{stat.Triplet[0], stat.Triplet[1], stat.Triplet[2]},
			},
			Occurence: int32(stat.Occurence),
		}
	}

	return &ComputeResponse{Stats: gstats}, nil
}
