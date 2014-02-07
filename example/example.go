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
