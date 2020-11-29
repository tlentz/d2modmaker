package weightrand

import (
	"fmt"
	"log"
	"math/rand"
)

// Weights Contains cache of weighted probabilities created with NewWeights
type Weights struct {
	alias []int
	prob  []float32
}

// NewWeights from https://www.keithschwarz.com/darts-dice-coins/
/*
Initialization:
    Create arrays Alias and Prob, each of size n
	Create two worklists, Small and Large
	Multiply each probability by n
	For each scaled probability pi
		If pi<1, add i to Small
		Otherwise (pi≥1), add i to Large

	While Small and Large are not empty: (Large might be emptied first)
	    Remove the first element from Small; call it l
		Remove the first element from Large; call it g
		Set Prob[l]=pl
		Set Alias[l]=g
		Set pg:=(pg+pl)−1 (This is a more numerically stable option.)
		If pg<1, add g to Small
		Otherwise (pg≥1), add g to Large
	While Large is not empty:
		Remove the first element from Large; call it g
		Set Prob[g]=1
	While Small is not empty: This is only possible due to numerical instability.
		Remove the first element from Small; call it l
		Set Prob[l]=1

Generation:
    Generate a fair die roll from an n-sided die; call the side i
	Flip a biased coin that comes up heads with probability Prob[i]
	If the coin comes up "heads," return i
	Otherwise, return Alias[i]
*/
// OBC: Can't pass an array of <blah> into a function using interface, so therefore
// skip passing the items themselves, just return an index into weights.
// Feed this an array of integer counts similar to "rarity" column in d2, then
// calls to Generate will return a weighted random index
func NewWeights(weights []int) *Weights {

	fmt.Printf("NewWeights: len:%d [0]=%d [1]=%d\n", len(weights), weights[0], weights[1])
	wr := Weights{}
	if len(weights) == 0 {
		log.Fatal("weightrand.NewWeights: no weights supplied")
	}
	numWeights := len(weights)
	p := make([]float32, numWeights)
	wr.prob = make([]float32, numWeights)
	wr.alias = make([]int, numWeights)

	small := make([]int, 0, numWeights)
	large := make([]int, 0, numWeights)
	//fmt.Printf("#weights=%d\n", len(weights))
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}
	multiplier := float32(numWeights) / float32(totalWeight)
	for idx, w := range weights {
		p[idx] = float32(w) * multiplier
	}
	//fmt.Println(p)
	for idx := range weights {
		if p[idx] < 1 {
			push(&small, idx)
		} else {
			push(&large, idx)
		}
		/* fmt.Printf("small:")
		fmt.Println(small)
		fmt.Printf("large:")
		fmt.Println(large) */
	}

	for (len(small) > 0) && (len(large) > 0) {
		l, _ := pop(&small)
		g, _ := pop(&large)
		wr.prob[l] = p[l]
		wr.alias[l] = g
		p[g] = (p[g] + p[l]) - 1
		if p[g] < 1 {
			push(&small, g)
		} else {
			push(&large, g)
		}
	}
	//fmt.Printf("wr after small & lg:")
	//fmt.Println(wr)
	for len(large) > 0 {
		g, _ := pop(&large)
		wr.prob[g] = 1
	}
	for len(small) > 0 {
		l, _ := pop(&small)
		wr.prob[l] = 1
	}
	if len(wr.prob) == 0 {
		panic(6)
	}
	return &wr
}

// Generate weighted random roll
func (wr *Weights) Generate() int {
	//fmt.Printf("WR:")
	//fmt.Println(wr)
	if wr == nil {
		log.Panic("Null pointer argument")
	}
	if wr.prob == nil {
		panic(8)
	}
	if len(wr.prob) == 0 {
		log.Panic("no entries in prob")
	}
	r := rand.Float32()
	r = r * float32(len(wr.prob))
	i := int(r)
	r = r - float32(i)

	val := 0
	if r > wr.prob[i] {
		val = wr.alias[i]
	} else {
		val = i
	}
	//fmt.Printf("Generate:%d\n", val)
	return val
}
func push(stack *[]int, item int) {
	*stack = append(*stack, item)
}
func pop(stack *[]int) (int, bool) {
	var item int
	l := len(*stack)
	if l == 0 {
		return -1, false
	}
	item, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return item, true

}
