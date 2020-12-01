package weightrand

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	//probs := []int{1, 2, 3, 4, 20}
	probs := []int{1, 1, 2, 1, 25}
	wr := NewWeights(probs)
	//fmt.Println(wr)
	var buckets = make([]int, len(probs))
	for i := 0; i < 10000; i++ {
		buckets[wr.Generate()]++
	}
	for idx, b := range buckets {
		if b == 0 {
			t.Errorf("Empty Bucket#%d", idx)
		}
	}
	fmt.Printf("Probs..:")
	fmt.Println(probs)
	fmt.Printf("Buckets:")
	fmt.Println(buckets)
}
