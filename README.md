<A name="toc1-0" title="What" />
# What

Package shuffle permutes arrays of integers while optionally
maintaining the positions, fixed or relative, of designated entries.

A fixed position entry retains the same position after the permutation.
In other words, it is not permuted. A relative position entry anchored
to its previous or next neighbor retains the same relative position
to the neighbor which may itself be permuted, fixed, or anchored to
another entry in a chain of dependencies. The relative positions are
maintained even if the anchor target is permuted.

Terminology:

"A > B" describes two items where A is anchored to its successor, B.
"A B <" describes the inverse, where B is anchored to its predecessor, A.
"A B . C" describes a list where B is anchored by position.

The unanchored items are permuted.

Edge Cases:

In "A > B <", where  A and B are mutually anchored, the dependency from A to B
is removed and the list converted into "A B <".

In "A < B C"  and "A B C >", where the endpoints are anchored to non-existent
neighbors, the anchors are converted into fixed positions.


<A name="toc1-30" title="Why" />
# Why

Permutations with anchoring is particularly useful for surveys and other
applications.

<A name="toc1-36" title="How" />
# How

Here is a complete example, also found in the example directory:

    package main
    
    import (
    	"fmt"
    	"log"
    
    	"github.com/gyepisam/shuffle"
    )
    
    func main() {
    	// Lightly edited list of types of shuffles, copied from Wikipedia
    	shuffles := []struct {
    		name   string
    		anchor shuffle.Anchor
    	}{
    		{"None", shuffle.Position}, // position anchored
    		{"Chemmy", shuffle.None},   // shuffled
    		{"Corgi", shuffle.None},
    		{"Faro", shuffle.None},
    		{"Indian", shuffle.None},
    		{"Irish", shuffle.ToPrevious},   // anchored to previous
    		{"Mexican", shuffle.ToPrevious}, // anchors can be chained
    		{"Mongean", shuffle.None},
    		{"Overhand", shuffle.Position},
    		{"Pile", shuffle.None},
    		{"Riffle", shuffle.None},
    		{"Stripping", shuffle.ToNext}, // anchored to next item.
    		{"Wash", shuffle.None},
    		{"Weave", shuffle.Position},
    	}
    
    	shuf := shuffle.New()
    	for i, shuffle := range shuffles {
    		shuf.Add(i, shuffle.anchor)
    	}
    
    	seed, err := shuffle.Seed()
    	if err != nil {
    		log.Fatal(err)
    	}
    	indices := shuf.Shuffle(seed)
    
    	fmt.Println("Sorted list of shuffles:")
    	for i, shuffle := range shuffles {
    		fmt.Printf("%d %s\n", i, shuffle.name)
    	}
    
    	fmt.Println("Permuted list of shuffles:")
    	for _, j := range indices {
    		fmt.Printf("%d %s\n", j, shuffles[j].name)
    	}
    }

<A name="toc1-43" title="Who" />
# Who

Shuffle is written by Gyepi Sam <self-github@gyepi.com>
