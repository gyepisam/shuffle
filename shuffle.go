/*
Package shuffle permutes arrays of integers with anchoring.
In addition to the standard of permuting all entries, shuffle
can maintain the positions, fixed or relative, of marked entries,
which is critical in a number of applications.
*/
package shuffle

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

type item struct {
	value  int
	anchor Anchor
}

// Shuffle permutes integers (generally representing positions) while maintaining anchored positions.
type Shuffle struct {
  items []item
}

// New returns an initialized shuffle.
func New() *Shuffle {
  return &Shuffle{[]item{}}
}

// Add appends new items to be shuffled.
// The anchor argument specifies whether and how the item is anchored.
func (s *Shuffle) Add(value int, anchor Anchor) {
	s.items = append(s.items, item{value: value, anchor: anchor})
}

// Shuffle shuffles and returns the list of integers added with Add,
// using seed to initialize the random number generator.
func (s *Shuffle) Shuffle(seed int64) []int {
    q := s.items
    n := len(q)

	if n > 0 {
		if q[0].anchor == ToPrevious {
			q[0].anchor = Position
		}

		if q[n-1].anchor == ToNext {
			q[n-1].anchor = Position
		}
	}

	permuted := []int{}
    for i := 0; i < n; i++ {
      if j := i+1; j < n {
        if q[i].anchor == ToNext && q[j].anchor == ToPrevious {
          q[i].anchor = None
        }
      }

      if q[i].anchor == None {
        permuted = append(permuted, i)
      }
    }

	out := []int{}
	emit := func(k int) {
        j := k
		for i := k - 1; i > -1 && q[i].anchor == ToNext; i-- {
		    j--
		}

        l := k+1
        for i := k + 1; i < n && q[i].anchor == ToPrevious; i++ {
            l++
		}

        for i := j; i < l; i++ {
          out = append(out, q[i].value)
        }
	}

	p := rand.New(rand.NewSource(seed)).Perm(len(permuted))

	for i, k := 0, 0; i < n; i++ {
		switch q[i].anchor {
		case Position:
			emit(i)
			break
		case None:
			emit(permuted[p[k]])
			k++
			break
		}
	}

	return out
}

// Seed is a convenience function which produces a 64bit int value read from crypto/rand.Reader.
func Seed() (int64, error) {
	val, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return 0, err
	}
	return val.Int64(), nil
}
