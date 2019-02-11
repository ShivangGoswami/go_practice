package main

import (
	"fmt"
	"sort"
)

type apnatype []string

func (a apnatype) Len() int {
	return len(a)
}
func (a apnatype) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a apnatype) Less(i, j int) bool {
	return a[i][len(a[i])-1] < a[j][len(a[j])-1]
}
func main() {
	items := apnatype{"Chips", "Cadbury", "Cola", "Choco"}
	sort.Sort(items)
	fmt.Println(items)
}
