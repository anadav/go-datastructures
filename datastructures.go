/*
Package datastructures exists solely to aid consumers of the go-datastructures
library when using dependency managers.  Depman, for instance, will work
correctly with any datastructure by simply importing this package instead of
each subpackage individually.

For more information about the datastructures package, see the README at

	http://github.com/anadav/go-datastructures

*/
package datastructures

import (
	_ "github.com/anadav/go-datastructures/augmentedtree"
	_ "github.com/anadav/go-datastructures/bitarray"
	_ "github.com/anadav/go-datastructures/btree/palm"
	_ "github.com/anadav/go-datastructures/btree/plus"
	_ "github.com/anadav/go-datastructures/fibheap"
	_ "github.com/anadav/go-datastructures/futures"
	_ "github.com/anadav/go-datastructures/hashmap/fastinteger"
	_ "github.com/anadav/go-datastructures/numerics/optimization"
	_ "github.com/anadav/go-datastructures/queue"
	_ "github.com/anadav/go-datastructures/rangetree"
	_ "github.com/anadav/go-datastructures/rangetree/skiplist"
	_ "github.com/anadav/go-datastructures/set"
	_ "github.com/anadav/go-datastructures/slice"
	_ "github.com/anadav/go-datastructures/slice/skip"
	_ "github.com/anadav/go-datastructures/sort"
	_ "github.com/anadav/go-datastructures/threadsafe/err"
	_ "github.com/anadav/go-datastructures/tree/avl"
	_ "github.com/anadav/go-datastructures/trie/xfast"
	_ "github.com/anadav/go-datastructures/trie/yfast"
)
