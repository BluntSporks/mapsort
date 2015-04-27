// mapsort provides functions for sorting maps.
package mapsort

import (
	"sort"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Key.
type KeyList []Pair

func (p KeyList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p KeyList) Len() int           { return len(p) }
func (p KeyList) Less(i, j int) bool { return p[i].Key < p[j].Key }

// A function to turn a map into a KeyList, then sort and return it.
func ByKey(m map[string]int, asc bool) KeyList {
	p := make(KeyList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	if asc {
		sort.Sort(p)
	} else {
		sort.Sort(sort.Reverse(p))
	}
	return p
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type ValueList []Pair

func (p ValueList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ValueList) Len() int           { return len(p) }
func (p ValueList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a ValueList, then sort and return it.
func ByValue(m map[string]int, asc bool) ValueList {
	p := make(ValueList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	if asc {
		sort.Sort(p)
	} else {
		sort.Sort(sort.Reverse(p))
	}
	return p
}
