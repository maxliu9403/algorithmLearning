package maxEnvelopes

import (
	"log"
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	log.Fatal(MaxEnvelopes(envelopes))
}
