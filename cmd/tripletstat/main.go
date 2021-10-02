package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/akhenakh/codingtestnr/wordsplit"
	"github.com/akhenakh/codingtestnr/wordstat"
)

func main() {
	s := wordstat.New()

	for _, a := range os.Args[1:] {
		f, err := os.Open(a)
		if err != nil {
			log.Fatal(err)
		}

		itr := wordsplit.ParseAndSplit(context.Background(), f)
		for {
			t, err := itr.Next()
			if err == wordsplit.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			s.Inc(wordstat.Triplet(t))
		}
	}

	stats := s.Compute()
	fmt.Printf("%v\n", stats)
}
