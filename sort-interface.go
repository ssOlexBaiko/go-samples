package main

import (
	"sort"
	"fmt"
	"flag"
)

// hook for implementing flag.Value and passing iSortArray through flag.Var()
type iSortArray []string

func (_ *iSortArray) String() string {
	return "my custom string representation"
}

func (a *iSortArray) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func (a iSortArray) Len() int {
	return len(a)
}

func (a iSortArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a iSortArray) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

var input iSortArray

func main() {
	flag.Var(&input, "input", "sequence for sorting. Usage: -input=list1 -input=list2 -input=list3")
	flag.Parse()
	sort.Sort(input)
	fmt.Println("Sorted: ", input)
}
