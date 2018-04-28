package intervals

import (
	"fmt"
	"sort"
)

type Intervals struct {
	a []Interval
}

type Interval struct {
	L, H float64
}

func (i *Intervals) Add(l, h float64) {
	i.a = append(i.a, Interval{l, h})

	i.SortAndMerge()
}

func (i *Intervals) AddToAll(l, h float64) {
	toAdd := []Interval{}
	for k := 0; k < len(i.a); k++ {
		toAdd = append(toAdd, Interval{
			L: i.a[k].L + l,
			H: i.a[k].H + h,
		})
	}

	for _, ii := range toAdd {
		i.a = append(i.a, ii)
	}

	i.SortAndMerge()
}

func (i *Intervals) Len() int           { return len(i.a) }
func (i *Intervals) Swap(a, b int)      { i.a[a], i.a[b] = i.a[b], i.a[a] }
func (i *Intervals) Less(a, b int) bool { return i.a[a].L < i.a[b].L }

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (i *Intervals) SortAndMerge() {
	sort.Sort(i)

	fmt.Println("sorted:", i.a)

	for k := 0; k < len(i.a)-1; k++ {
		if i.a[k].H >= i.a[k+1].L {
			i.a[k].H = max(i.a[k].H, i.a[k+1].H)
			i.a = append(i.a[:k+1], i.a[k+2:]...)
			k--
		}
	}
}
