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


# Why

Permutations with anchoring is particularly useful for surveys and other
applications.

# How

Here is a complete example, also found in the example directory:

.pull example/example.go,code

# Who

Shuffle is written by Gyepi Sam <self-github@gyepi.com>
