package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/akhenakh/codingtestnr/wordsplit"
	"github.com/akhenakh/codingtestnr/wordstat"
)

func main() {
	s := wordstat.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(os.Args) < 2 {
		parseData(ctx, os.Stdin, s)
	} else {
		for _, a := range os.Args[1:] {
			f, err := os.Open(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
			parseData(ctx, f, s)
			f.Close()
		}
	}

	stats := s.Compute()
	fmt.Printf("%v\n", stats)
}

func parseData(ctx context.Context, r io.Reader, s wordstat.Sink) error {
	itr := wordsplit.ParseAndSplit(ctx, r)
	for {
		t, err := itr.Next()
		if err == wordsplit.Done {
			break
		}
		if err != nil {
			return err
		}
		s.Inc(wordstat.Triplet(t))
	}

	return nil
}
