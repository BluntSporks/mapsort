// Package mapsort provides functions for sorting maps.
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

func (kl KeyList) Swap(i, j int)      { kl[i], kl[j] = kl[j], kl[i] }
func (kl KeyList) Len() int           { return len(kl) }
func (kl KeyList) Less(i, j int) bool { return kl[i].Key < kl[j].Key }

// A function to turn a map into a KeyList, then sort and return it.
func ByKey(myMap map[string]int, asc bool) KeyList {
	kl := make(KeyList, len(myMap))
	i := 0
	for key, value := range myMap {
		kl[i] = Pair{key, value}
		i++
	}
	if asc {
		sort.Sort(kl)
	} else {
		sort.Sort(sort.Reverse(kl))
	}
	return kl
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type ValueList []Pair

func (vl ValueList) Swap(i, j int)      { vl[i], vl[j] = vl[j], vl[i] }
func (vl ValueList) Len() int           { return len(vl) }
func (vl ValueList) Less(i, j int) bool { return vl[i].Value < vl[j].Value }

// A function to turn a map into a ValueList, then sort and return it.
func ByValue(myMap map[string]int, asc bool) ValueList {
	vl := make(ValueList, len(myMap))
	i := 0
	for key, value := range myMap {
		vl[i] = Pair{key, value}
		i++
	}
	if asc {
		sort.Sort(vl)
	} else {
		sort.Sort(sort.Reverse(vl))
	}
	return vl
}
